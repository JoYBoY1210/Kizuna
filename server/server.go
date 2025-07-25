package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func ChunkServer(port string, chunkDir string) {
	http.HandleFunc("/chunk/", func(w http.ResponseWriter, r *http.Request) {
		chunkID := strings.TrimPrefix(r.URL.Path, "/chunk/")
		chunkFile := fmt.Sprintf("%s/chunk_%s", chunkDir, chunkID)
		
		file, err := os.Open(chunkFile)
		if err != nil {
			http.Error(w, "Chunk not found", http.StatusNotFound)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = io.Copy(w, file)

	})
	fmt.Println("server running on port: ", port)
	http.ListenAndServe(":"+port, nil)

}
