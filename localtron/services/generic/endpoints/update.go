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
package genericendpoints

import (
	"encoding/json"
	"net/http"

	genericservice "github.com/singulatron/singulatron/localtron/services/generic"
	generictypes "github.com/singulatron/singulatron/localtron/services/generic/types"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

func Update(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	genericService *genericservice.GenericService,
) {
	err := userService.IsAuthorized(generictypes.PermissionGenericEdit.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, found, err := userService.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !found {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	req := &generictypes.UpdateRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, `invalid JSON`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = genericService.Update(req.Table, user.Id, req.Conditions, req.Object)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{}`))
}
