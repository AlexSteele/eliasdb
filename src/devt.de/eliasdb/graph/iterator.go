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
This module contains iterator related code.
*/
package graph

import (
	"devt.de/eliasdb/graph/util"
	"devt.de/eliasdb/hash"
)

/*
NodeKeyIterator can be used to iterate node keys of a certain node kind.
*/
type NodeKeyIterator struct {
	gm        *GraphManager       // GraphManager which created the iterator
	it        *hash.HTreeIterator // Internal HTree iterator
	LastError error               // Last encountered error
}

/*
Next returns the next node key. Sets the LastError attribute if an error occurs.
*/
func (it *NodeKeyIterator) Next() string {

	// Take reader lock

	it.gm.mutex.RLock()
	defer it.gm.mutex.RUnlock()

	k, _ := it.it.Next()

	if it.it.LastError != nil {
		it.LastError = &util.GraphError{util.ErrReading, it.it.LastError.Error()}
		return ""
	} else if len(k) == 0 {
		return ""
	}

	return string(k[len(PREFIX_NS_ATTRS):])
}

/*
HasNext returns if there is a next node key.
*/
func (it *NodeKeyIterator) HasNext() bool {
	return it.it.HasNext()
}
