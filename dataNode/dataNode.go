package datanode

import (
	"bufio"
	"dfs/proto"
	"io"
	"log"
	"os"
)

var blockSize = 1000

// WriteDataToBlock dadf
func WriteDataToBlock(blockName string, data []byte) {

}

// ReadBlockData asdfa
func ReadBlockData(blockName string, stream proto.) []byte {

	file, err := os.Open(blockName)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		res := &proto.File{
			Content: buffer[:n],
		}

		// res := &pb.UploadImageRequest{
		// 	Data: &pb.UploadImageRequest_ChunkData{
		// 		ChunkData: buffer[:n],
		// 	},
		// }

		err = stream.Send(res)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}

	// response := make([]byte, blockSize)
	// return response
}

// storageid, blockreport, commit, handshake
