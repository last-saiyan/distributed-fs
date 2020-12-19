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

func (b *Block) initBlock(blockName string, mode string, blockSize int, chunkSize int) {
	var file *os.File
	var err error
	var reader *bufio.Reader
	if mode == "r" {
		file, err = os.Open(blockName)
		reader = bufio.NewReader(file)
	} else if mode == "w" {
		file, err = os.OpenFile(blockName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	}
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	b.file = file
	b.blockName = blockName
	b.reader = reader
	b.chunkSize = chunkSize
	b.blockSize = blockSize
	buffer := make([]byte, b.chunkSize)
	b.buffer = &buffer
	b.dataRead = -1
	b.offset = 0
}

func (b *Block) hasNextChunk() bool {
	if b.dataRead != -1 {
		return true
	}
	n, err := b.reader.Read(*b.buffer)
	if err == io.EOF {
		b.dataRead = -1
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
	finfo, err := b.file.Stat()
	if err != nil {
		log.Fatal("cannot open the file: ", err)
	}
	currBlockSize := int(finfo.Size())

	if b.blockSize < (len(chunk) + currBlockSize) {
		n, err := b.file.Write(chunk)
		if err != nil {
			log.Fatal("cannot write to file: ", err)
		}
		return n, nil
	}
	if b.blockSize < currBlockSize {
		return -1, nil
	}
	if b.blockSize == currBlockSize {
		return 0, nil
	}
	capacity := (b.blockSize - currBlockSize)
	n, err := b.file.Write(chunk[:capacity])
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	return n, nil
}

// storageid, blockreport, commit, handshake
