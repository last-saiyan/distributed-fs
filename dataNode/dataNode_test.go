package datanode

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// var config = utils.GetConfig()

func TestRead(t *testing.T) {
	tempFile := "temp.txt"
	tempContent := "here is some temp content"
	createTempFile(tempFile, tempContent)
	b := Block{}
	b.InitBlock(tempFile, "r")
	temp := make([]byte, 0)
	for b.HasNextChunk() {
		chunk, size, _ := b.GetNextChunk()
		temp = append(temp, (*chunk)[:size]...)
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

func TestWrite(t *testing.T) {
	tempFile := "temrp.txt"
	tempContent := "here is some temp content"
	// createTempFile(tempFile, tempContent)
	b := Block{}
	b.InitBlock(tempFile, "w")

	fileBytes := []byte(tempContent)
	iter := 0
	for iter < len(fileBytes) {
		end := iter + config.IoSize
		if end > len(fileBytes) {
			end = len(fileBytes)
		}
		chunk := fileBytes[iter:end]

		err := b.WriteChunk(chunk)
		if err != nil {
			fmt.Println(err)
		}
		iter = iter + config.IoSize
	}

	b = Block{}
	b.InitBlock(tempFile, "r")
	temp := make([]byte, 0)
	for b.HasNextChunk() {
		chunk, size, _ := b.GetNextChunk()
		temp = append(temp, (*chunk)[:size]...)
	}
	if tempContent != string(temp) {
		t.Fail()
	}

	deleteTempFile(tempFile)
}
