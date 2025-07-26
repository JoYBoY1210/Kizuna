package downloads

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"kizuna/types"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
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

	var wg sync.WaitGroup
	errs := make(chan error, meta.NumChunks)

	for i := 0; i < meta.NumChunks; i++ {
		wg.Add(1)
		go func(chunkIndex int) {
			defer wg.Done()
			chunkFile := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", chunkIndex))

			expectedHash := meta.ChunkHashes[chunkIndex]
			err := downloadChunk(chunkIndex, chunkFile, expectedHash, meta.Peers)
			if err != nil {
				errs <- fmt.Errorf("error downloading chunk %d: %v", chunkIndex, err)
				return
			}
			fmt.Println("Downloaded chunk:", chunkIndex)
		}(i)
	}

	wg.Wait()
	close(errs)

	if len(errs) > 0 {
		return <-errs
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

	peerShuffling := make([]string, len(peers))
	copy(peerShuffling, peers)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(peerShuffling), func(i, j int) {
		peerShuffling[i], peerShuffling[j] = peerShuffling[j], peerShuffling[i]
	})

	for _, peer := range peerShuffling {
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

		fmt.Printf("Downloading chunk  from peer %s\n", peer)

		return nil
	}

	return fmt.Errorf("failed to download chunk %d from all peers", index)
}
