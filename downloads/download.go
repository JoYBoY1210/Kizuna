package downloads

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"kizuna/types"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(metaPath string, outputPath string) error {

	metaData, err := os.ReadFile(metaPath)
	if err != nil {
		return fmt.Errorf("error reading meta file: %v", err)
	}

	var meta types.MetaFile
	err = json.Unmarshal(metaData, &meta)
	if err != nil {
		return fmt.Errorf("error unmarshalling meta file: %v", err)
	}

	tempDir := "temp_chunks"
	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating temp directory: %v", err)
	}

	for i := 0; i < meta.NumChunks; i++ {
		chunkFile := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", i))
		
		expectedHash := meta.ChunkHashes[i]
		if err := downloadChunk(i, chunkFile, expectedHash, meta.Peers); err != nil {
			return fmt.Errorf("error downloading chunk %d: %v", i, err)
		}
		fmt.Println("Downloaded chunk:", i)
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer out.Close()

	for i := 0; i < meta.NumChunks; i++ {
		chunkFile := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", i))
		chunk, err := os.Open(chunkFile)
		if err != nil {
			return fmt.Errorf("error opening chunk file %d: %v", i, err)
		}
		_, err = io.Copy(out, chunk)
		chunk.Close()
		if err != nil {
			return fmt.Errorf("error merging chunk %d: %v", i, err)
		}
	}

	fmt.Println("File downloaded successfully to:", outputPath)

	os.RemoveAll(tempDir)

	return nil
}

func downloadChunk(index int, path string, expectedHash string, peers []string) error {

	for _, peer := range peers {
		url := fmt.Sprintf("%s/chunk/%d", peer, index)

		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != http.StatusOK {
			fmt.Printf("Error from peer: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		chunkData, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			continue
		}

		hasher := sha256.New()
		hasher.Write(chunkData)
		hash := hex.EncodeToString(hasher.Sum(nil))

		if hash != expectedHash {
			fmt.Printf("Hash mismatch\n")
			continue
		}

		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("error creating chunk file %d: %v", index, err)
		}
		defer file.Close()

		_, err = file.Write(chunkData)
		if err != nil {
			return fmt.Errorf("error writing chunk %d to file: %v", index, err)
		}

		return nil
	}

	return fmt.Errorf("failed to download chunk %d from all peers", index)
}
