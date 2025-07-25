package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const ChunkSize = 1024 * 1024

func ChunkFile(filePath string) (int, error) {
	file, error := os.Open(filePath)
	if error != nil {
		return 0, error
	}
	defer file.Close()

	os.RemoveAll("chunks")
	error = os.MkdirAll("chunks", os.ModePerm)
	if error != nil {
		return 0, error
	}

	buffer := make([]byte, ChunkSize)
	chunkCount := 0

	for {
		bytesRead, readErr := file.Read(buffer)
		if bytesRead == 0 {
			break
		}
		chunkFileName := filepath.Join("chunks", fmt.Sprintf("chunk_%d", chunkCount))
		chunkFile, err := os.Create(chunkFileName)
		if err != nil {
			return 0, err
		}
		defer chunkFile.Close() 
		_, err = chunkFile.Write(buffer[:bytesRead])
		if err != nil {
			return 0, err
		}
		chunkCount++

		if readErr == io.EOF {
			break
		} else if readErr != nil {
			return chunkCount, readErr 
		}
	}
	return chunkCount, nil
}
