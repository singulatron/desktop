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
package dapper

import (
	"encoding/json"
	"io/ioutil"

	dt "github.com/singulatron/singulatron/localtron/dapper/types"
)

// LoadAppConfiguration loads an application configuration from a JSON file.
func (cm *ConfigurationManager) LoadAppConfiguration(filePath string) (*dt.App, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var app dt.App
	err = json.Unmarshal(file, &app)
	if err != nil {
		return nil, err
	}

	if app.Name != "" {
		cm.Printf("Processing app '%s' (%s).\n", app.Name, filePath)
	} else {
		cm.Printf("Processing unnamed app (%s).\n", filePath)
	}

	return &app, nil
}
