/*
 * EliasDB
 *
 * Copyright 2016 Matthias Ladkau. All rights reserved.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

/*
Package api contains general REST API definitions.

Endpoint which returns an object with version information.

/about

{
    api_versions : List of available API versions e.g. [ "v1" ]
    product      : Name of the API provider (EliasDB)
    version:     : Version of the API provider
    revision:    : Revision of the API provider
}
*/
package api

import (
	"encoding/json"
	"net/http"

	"devt.de/eliasdb/version"
)

/*
EndpointAbout is the about endpoint URL (rooted). Handles about/
*/
const EndpointAbout = APIRoot + "/about/"

/*
AboutEndpointInst creates a new endpoint handler.
*/
func AboutEndpointInst() RestEndpointHandler {
	return &aboutEndpoint{}
}

/*
Handler object for about operations.
*/
type aboutEndpoint struct {
	*DefaultEndpointHandler
}

/*
HandleGET returns about data for the REST API.
*/
func (a *aboutEndpoint) HandleGET(w http.ResponseWriter, r *http.Request, resources []string) {

	data := map[string]interface{}{
		"api_versions": []string{"v1"},
		"product":      "EliasDB",
		"version":      version.VERSION,
		"revision":     version.REV,
	}

	// Write data

	w.Header().Set("content-type", "application/json; charset=utf-8")

	ret := json.NewEncoder(w)
	ret.Encode(data)
}
