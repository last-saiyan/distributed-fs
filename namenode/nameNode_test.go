package namenode

import (
	"testing"
)

func TestInit(t *testing.T) {
	if filenameTochunk != nil {
		t.Error("filenameTochunk is initilized before init")
	}
	if chunkNameToServer != nil {
		t.Error("chunkNameToServer is initilized before init")
	}

	Init()

	if filenameTochunk == nil {
		t.Error("filenameTochunk is not initilized after init")
	}
	if chunkNameToServer == nil {
		t.Error("chunkNameToServer is not initilized after init")
	}
}

func TestAppendToFile(t *testing.T) {

}
