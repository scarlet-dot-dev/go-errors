// Copyright (c) 2020 Eddy <eddy@scarlet.dev>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package errors

// Handle will call fn on *e if it is not nil, or if the function is recovering
// from a panic.
func Handle(e *error, fn func(err error) error) {
	// first check for a panic
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			panic(err)
		}
		*e = fn(err)
		return
	}

	// not a panic / recover
	if err := *e; err != nil {
		*e = fn(err)
	}
}

// Check error will panic with err if the error is not nil.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
