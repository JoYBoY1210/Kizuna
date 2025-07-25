package main

import (
	"fmt"
	"kizuna/files"
)

func main() {
	count, err := files.ChunkFile("C:\\Users\\tanish\\Desktop\\tanish 11th\\Acids.pdf")
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
	fmt.Println("Hashes of chunks:")
	for _, hash := range hashes {
		fmt.Println(hash)
	}
}
