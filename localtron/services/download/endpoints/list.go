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
package downloadendpoints

import (
	"encoding/json"
	"net/http"

	downloadservice "github.com/singulatron/singulatron/localtron/services/download"
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	types "github.com/singulatron/singulatron/localtron/services/download/types"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
)

func List(
	w http.ResponseWriter,
	r *http.Request,
	userService *userservice.UserService,
	ds *downloadservice.DownloadService,
) {
	err := userService.IsAuthorized(downloadtypes.PermissionDownloadView.Id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	details, err := ds.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(types.DownloadsResponse{
		Downloads: details,
	})
	w.Write(jsonData)
}
