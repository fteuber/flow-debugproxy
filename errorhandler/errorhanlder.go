// Copyright 2015 Dominique Feyer <dfeyer@ttree.ch>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errorhandler

import (
	"../logger"
	"os"
)

// PanicHandling handle error and output log message
func PanicHandling(err error, logger *logger.Logger) {
	if err != nil {
		logger.Warn(err.Error())
		os.Exit(1)
	}
}
