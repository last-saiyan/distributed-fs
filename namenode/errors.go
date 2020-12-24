package namenode

import "errors"

// ErrFileNotFound file is not found
var ErrFileNotFound = errors.New("file not found")

// ErrFileExists file with this name is already present
var ErrFileExists = errors.New("file with this name is present")

// ErrFileOpen file is not yet closed
var ErrFileOpen = errors.New("file is not yet closed")

// ErrReplicaCount not enough datanode to replicate
var ErrReplicaCount = errors.New("replica count exceeds datanode count")
