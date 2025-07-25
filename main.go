package main

import (
	"kizuna/files"
	"kizuna/server"
	"kizuna/downloads"
	"fmt"
	"os"
	"kizuna/types"
	"time"

)

func main() {
	path := "C:\\Users\\tanish\\Desktop\\tanish 11th\\Acids.pdf"
	count, err := files.ChunkFile(path)
	if err != nil {
		fmt.Println("Error chunking file:", err)
		return
	}
	fmt.Println("Number of chunks created:", count)

	hashes, err := files.HashChunks("chunks")
	if err != nil {
		fmt.Println("Error hashing chunks:", err)
		return
	}
	
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	meta := types.MetaFile{
		FileName:    info.Name(),
		FileSize:    info.Size(),
		ChunkSize:   1024 * 1024,
		NumChunks:   count,
		ChunkHashes: hashes,
		Peers:       []string{"http://localhost:8080"},
	}

	err = files.NewMetaFile(meta, fmt.Sprintf("%s.meta", info.Name()))
	if err != nil {
		fmt.Println("Error creating meta file:", err)
		return
	}

	go func() {
		server.ChunkServer("8080", "chunks")

	}()
	
	time.Sleep(1 * time.Second)
	err = downloads.DownloadFile("Acids.pdf.meta", "output.pdf")
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	select {}
	

}
