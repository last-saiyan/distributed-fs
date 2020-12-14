package datanode

import (
	"bufio"
	"io"
	"log"
	"os"
)

var bufferSize = 1000

// WriteDataToBlock dadf
func WriteDataToBlock(blockName string, data []byte) {

}

type block struct {
	blockName string
	offset    int
	chunkSize int
	reader    *bufio.Reader
	buffer    *[]byte
	dataRead  int
	file      *os.File
}

func (b block) initBlock(new bool, blockName string) {
	file, err := os.Open(blockName)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	reader := bufio.NewReader(file)
	b.file = file
	b.blockName = blockName
	b.reader = reader
	buffer := make([]byte, bufferSize)
	b.buffer = &buffer
	b.dataRead = -1
}

func (b block) hasNextChunk() bool {
	n, err := b.reader.Read(*b.buffer)
	if err == io.EOF {
		b.file.Close()
		return false
	}
	if err != nil {
		log.Fatal("cannot read chunk to buffer: ", err)
	}
	b.dataRead = n
	return true
}

func (b block) getNextChunk() ([]byte, int, error) {
	if b.dataRead == -1 {
		return nil, 0, nil
	}
	n := b.dataRead
	b.dataRead = -1
	return *b.buffer, n, nil
}

// storageid, blockreport, commit, handshake
