package files

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	// "sort"
)

func HashChunks(chunkDir string) (map[int]string, error) {
	files, err := filepath.Glob(filepath.Join(chunkDir, "chunk_*"))
	if err != nil {
		fmt.Println("Error finding chunk files:", err)
		return nil, err
	}
	// sort.Slice(files, func(i, j int) bool {
	// 	return files[i] < files[j]
	// })
	

	hashes := make(map[int]string)
	for _, file := range files {
		
		var index int
		_, err := fmt.Sscanf(filepath.Base(file), "chunk_%d", &index)
		if err != nil {
			return nil, fmt.Errorf("invalid chunk filename format: %s", file)
		}

		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening chunk file:", err)
			return nil, err
		}
		defer f.Close()

		hasher := sha256.New()
		if _, err := io.Copy(hasher, f); err != nil {
			f.Close()
			return nil, err
		}
		hash := hex.EncodeToString(hasher.Sum(nil))
		hashes[index] = hash
	}
	return hashes, nil
}
