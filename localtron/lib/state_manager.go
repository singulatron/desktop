/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package lib

import (
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"
	"time"
)

type StateManager[T any] struct {
	key        string
	memStore   *MemoryStore[T]
	lock       sync.Mutex
	filePath   string
	hasChanged bool
}

// key: the key under which the slice will be saved in a JSON object in the file
func NewStateManager[T any](key string, memStore *MemoryStore[T], filePath string) *StateManager[T] {
	sm := &StateManager[T]{
		key:      key,
		memStore: memStore,
		filePath: filePath,
	}
	sm.setupSignalHandler()
	return sm
}

func (sm *StateManager[T]) LoadState() error {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	_, err := os.Stat(sm.filePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(sm.filePath), 0755)
		if err != nil {
			return err
		}
		err = os.WriteFile(sm.filePath, []byte("{}"), 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	data, err := os.ReadFile(sm.filePath)
	if err != nil {
		return err
	}

	var items []T
	err = json.Unmarshal(data, &items)
	if err != nil {
		return err
	}

	sm.memStore.Reset(items)

	return nil
}

func (sm *StateManager[T]) SaveState() error {
	shallowCopy := sm.memStore.SliceCopy()

	sm.lock.Lock()
	data, err := json.MarshalIndent(shallowCopy, "", "  ")
	if err != nil {
		sm.lock.Unlock()
		return err
	}
	sm.hasChanged = false
	sm.lock.Unlock()

	err = os.WriteFile(sm.filePath, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (sm *StateManager[T]) MarkChanged() {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	sm.hasChanged = true
}

func (sm *StateManager[T]) PeriodicSaveState(interval time.Duration) {
	for {
		time.Sleep(interval)
		if !sm.hasChanged {
			continue
		}
		if err := sm.SaveState(); err != nil {
			Logger.Error("Error saving file state",
				slog.String("filePath", sm.filePath),
			)
		}
	}
}

func (sm *StateManager[T]) setupSignalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		sm.SaveState()
	}()
}
