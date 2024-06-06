package promptservice

import (
	"strings"

	"github.com/singulatron/singulatron/localtron/llm"
)

func llmResponseToText(responses []*llm.CompletionResponse) string {
	var result strings.Builder

	first := true
	for _, v := range responses {
		if len(v.Choices) == 0 {
			continue
		}
		choice := v.Choices[0]

		var textToAdd string
		if strings.Contains(result.String(), "```") {
			// Handling for inline code formatting if the resulting string is already within a code block
			count := strings.Count(result.String(), "```")
			if count%2 == 1 { // If the count of ``` is odd, we are inside a code block
				textToAdd = choice.Text // No escaping needed inside code block
			} else {
				textToAdd = escapeHtml(choice.Text) // Apply HTML escaping when outside code blocks
			}
		} else {
			textToAdd = escapeHtml(choice.Text) // Apply HTML escaping if there is no code block
		}

		if first {
			textToAdd = strings.TrimLeft(textToAdd, " ")
			first = false
		}

		result.WriteString(textToAdd)

		if choice.FinishReason == "stop" {
			break
		}
	}

	return result.String()
}

func escapeHtml(input string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(input)
}

func errToString(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}
