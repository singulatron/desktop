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
	"errors"
	"time"

	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
)

func (s *UserService) UpsertRole(id, name, description string, permissionIds []string) (*usertypes.Role, error) {
	permissions := s.permissionsMem.FindByIds(permissionIds)
	if len(permissions) < len(permissionIds) {
		return nil, errors.New("nonexistent permissions")
	}

	defer s.rolesFile.MarkChanged()

	role, found := s.rolesMem.FindById(id)
	if !found {
		role = &usertypes.Role{
			Id:            id,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			Name:          name,
			Description:   description,
			PermissionIds: permissionIds,
		}
		s.rolesMem.Add(role)
		return role, nil
	}

	existingPermissionIdIndex := map[string]struct{}{}
	for _, permissionId := range role.PermissionIds {
		existingPermissionIdIndex[permissionId] = struct{}{}
	}

	for _, permissionId := range permissionIds {
		_, ok := existingPermissionIdIndex[permissionId]
		if !ok {
			role.PermissionIds = append(role.PermissionIds, permissionId)
		}
	}

	return role, nil
}