package promptservice

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/singulatron/singulatron/localtron/lib"
	"github.com/singulatron/singulatron/localtron/llm"
	apptypes "github.com/singulatron/singulatron/localtron/services/app/types"
	prompttypes "github.com/singulatron/singulatron/localtron/services/prompt/types"
)

const (
	maxRetries = 5
	baseDelay  = time.Second
)

// a blocking method, call it in a goroutine
func (p *PromptService) processPrompts() {
	ticker := time.NewTicker(2000 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		case <-p.trigger:
		}

		err := p.processNextPrompt()
		if err != nil {
			lib.Logger.Error("Error processing prompt",
				slog.String("error", err.Error()),
			)
		}
	}
}

func (p *PromptService) processNextPrompt() error {
	p.runMutex.Lock()
	defer p.runMutex.Unlock()

	if p.promptsMem.Count(func(v *prompttypes.Prompt) bool {
		return v.Status == prompttypes.PromptStatusRunning
	}) > 0 {
		return nil
	}

	currentPrompt := selectPrompt(p.promptsMem)
	if currentPrompt == nil {
		return nil
	}

	return p.processPrompt(currentPrompt)
}

func (p *PromptService) processPrompt(currentPrompt *prompttypes.Prompt) (err error) {
	lib.Logger.Info("Picking up prompt from queue",
		slog.String("promptId", currentPrompt.Id),
	)

	p.firehoseService.Publish(prompttypes.EventPromptProcessingStarted{
		PromptId: currentPrompt.Id,
	})

	p.promptsFile.MarkChanged()

	defer p.firehoseService.Publish(prompttypes.EventPromptProcessingFinished{
		PromptId: currentPrompt.Id,
		Error:    errToString(err),
	})
	defer func() {
		if err != nil {
			currentPrompt.Error = err.Error()
			currentPrompt.Status = prompttypes.PromptStatusErrored
		} else {
			currentPrompt.Status = prompttypes.PromptStatusCompleted
		}
		p.promptsFile.MarkChanged()
	}()

	currentPrompt.Error = ""
	currentPrompt.Status = prompttypes.PromptStatusRunning
	currentPrompt.RunCount++

	err = p.appService.AddChatMessage(&apptypes.ChatMessage{
		// not a fan of this but at least it makes this idempotent
		// in case prompts get retried over and over again
		Id:             currentPrompt.Id,
		ThreadId:       currentPrompt.ThreadId,
		IsUserMessage:  true,
		MessageContent: currentPrompt.Message,
		Time:           time.Now().Format(time.RFC3339),
	})
	if err != nil {
		return err
	}

	stat, err := p.modelService.Status(currentPrompt.ModelId)
	if err != nil {
		return errors.Wrap(err, "error getting model status")
	}
	if !stat.Running {
		return fmt.Errorf("model '%v' is not running", currentPrompt.ModelId)
	}
	if stat.ModelAddress == "" {
		return errors.Wrap(err, "missing model address")
	}
	if !strings.HasPrefix(stat.ModelAddress, "http") {
		stat.ModelAddress = "http://" + stat.ModelAddress
	}
	llmClient := llm.Client{
		LLMAddress: stat.ModelAddress,
	}

	err = llmClient.PostCompletionsStreamed(llm.PostCompletionsRequest{
		Prompt:    currentPrompt.Prompt,
		Stream:    true,
		MaxTokens: 4096,
	}, func(resp *llm.CompletionResponse) {
		p.StreamManager.Broadcast(currentPrompt.ThreadId, resp)
		if len(resp.Choices) > 0 && resp.Choices[0].FinishReason == "stop" {
			err := p.appService.AddChatMessage(&apptypes.ChatMessage{
				Id:             uuid.New().String(),
				ThreadId:       currentPrompt.ThreadId,
				MessageContent: llmResponseToText(p.StreamManager.history[currentPrompt.ThreadId]),
			})
			if err != nil {
				lib.Logger.Error("Error when saving chat message after broadcast",
					slog.String("error", err.Error()))
				return
			}

			delete(p.StreamManager.history, currentPrompt.ThreadId)
		}
	})
	if err != nil {
		return errors.Wrap(err, "error streaming llm")
	}

	return nil
}
