package files

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func HashChunks(chunkDir string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(chunkDir, "chunk_*"))
	if err != nil {
		fmt.Println("Error finding chunk files:", err)
		return nil, err
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	var hashes []string
	for _, file := range files {
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
		hashes = append(hashes, hash)
	}
	return hashes, nil
}
