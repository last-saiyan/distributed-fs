package datanode

import (
	"log"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	tempFile := "temp.txt"
	tempContent := "here is some temp content"
	createTempFile(tempFile, tempContent)
	b := Block{}
	b.initBlock(tempFile, "r", 500, 20)
	temp := make([]byte, 0)
	for b.hasNextChunk() {
		chunk, size, _ := b.getNextChunk()
		temp = append(temp, chunk[:size]...)
	}
	if tempContent != string(temp) {
		t.Fail()
	}
	deleteTempFile(tempFile)
}

func createTempFile(name string, content string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal("cannot create file: ", err)
	}
	_, err1 := file.Write([]byte(content))
	if err1 != nil {
		log.Fatal("cannot write file: ", err1)
	}
}

func deleteTempFile(name string) {
	err := os.Remove(name)
	if err != nil {
		log.Fatal("cannot delete file: "+name, err)
	}
}
