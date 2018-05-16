/*-
 * Copyright (c) 2018, 1&1 Internet SE
 * All rights reserved
 *
 * Use of this source code is governed by a 2-clause BSD license
 * that can be found in the LICENSE file.
 */

package msg // import "github.com/mjolnir42/eye/internal/eye.msg"

import "time"

var (
	// AssertionsAreFatal causes triggered assertions to panic within
	// the library
	AssertionsAreFatal bool

	// go time.Time.Parse can not handle time strings outside of
	// years [0,9999]. Therefor, obviously, negative infinity is 1...

	// NegTimeInf will be used as mapping for the PostgreSQL time value
	// -infinity. Dates earlier than this will be truncated to
	// NegTimeInf. RFC3339: 0001-01-01T00:00:00Z
	NegTimeInf = time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	// PosTimeInf will be used as mapping for the PostgreSQL time value
	// +infinity. Dates later than this will be truncated to PosTimeInf.
	// RFC3339: 8192-01-01T00:00:00Z
	PosTimeInf = time.Date(8192, time.January, 1, 0, 0, 0, 0, time.UTC)
)

// assertIsNil verifies that err is nil
func assertIsNil(err error) {
	if AssertionsAreFatal && err != nil {
		panic(err)
	}
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix
