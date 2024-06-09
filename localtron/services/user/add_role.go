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
package userservice

import (
	"time"

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) AddRole(userID string, role *usertypes.Role) error {
	s.usersMem.ForeachStop(func(i int, user *usertypes.User) bool {
		if user.Id == userID {
			alreadyHas := false
			for _, v := range user.RoleIds {
				if v == role.Id {
					alreadyHas = true
				}
			}
			if alreadyHas {
				return true
			}
			user.RoleIds = append(user.RoleIds, role.Id)
			user.UpdatedAt = time.Now()
			return true
		}
		return false
	})

	s.usersFile.MarkChanged()
	return nil
}
