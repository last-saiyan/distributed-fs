package datanode

import (
	"bufio"
	"dfs/utils"
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
	blockSize int64
}

var config = utils.GetConfig()

// GetNewBlock gets new initilized struct
func GetNewBlock(blockName string, mode string) *Block {
	b := Block{}
	b.initBlock(blockName, mode)
	return &b
}

// InitBlock initilizes the struct with necessary details
func (b *Block) initBlock(blockName string, mode string) {
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
	b.chunkSize = config.IoSize
	b.blockSize = config.BlockSize
	buffer := make([]byte, b.chunkSize)
	b.buffer = &buffer
	b.dataRead = -1
	b.offset = 0
}

// HasNextChunk checks if next chunk of data is present in the block
func (b *Block) HasNextChunk() bool {
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

// GetNextChunk gets the next chunk of data
func (b *Block) GetNextChunk() (*[]byte, int, error) {
	if b.dataRead == -1 {
		return nil, 0, nil
	}
	n := b.dataRead
	b.dataRead = -1
	return b.buffer, n, nil
}

// WriteChunk writes given chunk
func (b *Block) WriteChunk(chunk []byte) error {
	finfo, err := b.file.Stat()
	if err != nil {
		log.Fatal("cannot open the file: ", err)
	}
	currBlockSize := finfo.Size()
	if b.blockSize >= (int64(len(chunk)) + currBlockSize) {
		_, err := b.file.Write(chunk)
		if err != nil {
			log.Fatal("cannot write to file: ", err)
		}
		return nil
	}
	return ErrFileExceedsBlockSize
}

// Close is used to close the file
func (b *Block) Close() error {
	return b.file.Close()
}

// storageid, blockreport, commit, handshake
