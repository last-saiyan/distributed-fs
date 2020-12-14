package datanode

import (
	"bufio"
	"io"
	"log"
	"os"
)

// Block struct provides api to read and write block
type Block struct {
	blockName string
	offset    int
	chunkSize int
	reader    *bufio.Reader
	buffer    *[]byte
	dataRead  int
	file      *os.File
	blockSize int
}

func (b *Block) initBlock(new bool, blockName string) {
	file, err := os.Open(blockName)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	reader := bufio.NewReader(file)
	b.file = file
	b.blockName = blockName
	b.reader = reader
	buffer := make([]byte, b.chunkSize)
	b.buffer = &buffer
	b.dataRead = -1
}

func (b *Block) hasNextChunk() bool {
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

func (b *Block) getNextChunk() ([]byte, int, error) {
	if b.dataRead == -1 {
		return nil, 0, nil
	}
	n := b.dataRead
	b.dataRead = -1
	return *b.buffer, n, nil
}

func (b *Block) writeChunk(chunk []byte) (int, error) {
	// n, err := b.file.Write(chunk)
	finfo, err := b.file.Stat()
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	currBlockSize := finfo.Size()

	if b.blockSize < (len(chunk) + int(currBlockSize)) {
		n, err := b.file.Write(chunk)
		if err != nil {
			log.Fatal("cannot open image file: ", err)
		}
		return n, nil
	}
	if b.blockSize < int(currBlockSize) {
		return -1, nil
	}
	if b.blockSize == int(currBlockSize) {
		return 0, nil
	}
	capacity := (b.blockSize - int(currBlockSize))
	n, err := b.file.Write(chunk[:capacity])
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	return n, nil
}

// storageid, blockreport, commit, handshake
