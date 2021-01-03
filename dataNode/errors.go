package datanode

import "errors"

// ErrFileExceedsBlockSize file is not found
var ErrFileExceedsBlockSize = errors.New("file is greater than block size")
