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
	"path"
	"sync"
	"time"
)

type StateManager[T any] struct {
	state      T
	lock       sync.Mutex
	filePath   string
	hasChanged bool
}

func NewStateManager[T any](initialState T, filePath string) *StateManager[T] {
	return &StateManager[T]{
		state:    initialState,
		filePath: filePath,
	}
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

	return json.Unmarshal(data, &sm.state)
}

func (sm *StateManager[T]) SaveState() error {
	sm.lock.Lock()
	data, err := json.MarshalIndent(sm.state, "", "  ")
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
		sm.lock.Lock()
		if sm.hasChanged {
			sm.lock.Unlock()
			if err := sm.SaveState(); err != nil {
				Logger.Error("Error saving file state",
					slog.String("filePath", sm.filePath),
				)
			}
		} else {
			sm.lock.Unlock()
		}
	}
}
