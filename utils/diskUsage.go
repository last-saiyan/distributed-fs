package utils

import (
	"fmt"

	"github.com/minio/minio/pkg/disk"
)

// DiskUsage returns bytes of free space
func DiskUsage(path string) uint64 {
	di, err := disk.GetInfo(path)
	if err != nil {
		fmt.Println(err, "err")
	}
	return di.Ffree
}
