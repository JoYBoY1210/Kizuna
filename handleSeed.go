package main

import (
	"encoding/json"
	"fmt"
	"kizuna/files"
	"kizuna/server"
	"kizuna/types"
	"os"
	"strings"
)

func ReadMetaFile(path string) (types.MetaFile, error) {
	var meta types.MetaFile
	data, err := os.ReadFile(path)
	if err != nil {
		return meta, err
	}

	err = json.Unmarshal(data, &meta)
	if err != nil {
		return meta, err
	}

	return meta, nil
}

func handleSeed(filePath, port, peerList, metaPath string) {
	var meta types.MetaFile
	var err error
	if metaPath != "" {
		
		meta, err = ReadMetaFile(metaPath)
		if err != nil {
			fmt.Println("Error reading metadata file:", err)
			return
		}
	} else {
		count, err := files.ChunkFile(filePath)
		if err != nil {
			fmt.Println("Error chunking file:", err)
			return
		}

		hashes, err := files.HashChunks("chunks")
		if err != nil {
			fmt.Println("Error hashing chunks:", err)
			return
		}

		info, err := os.Stat(filePath)
		if err != nil {
			fmt.Println("Error getting file info:", err)
			return
		}

		meta = types.MetaFile{
			FileName:    info.Name(),
			FileSize:    info.Size(),
			ChunkSize:   1024 * 1024,
			NumChunks:   count,
			ChunkHashes: hashes,
			Peers:       strings.Split(peerList, ","),
		}

		err = files.NewMetaFile(meta, fmt.Sprintf("%s.meta", info.Name()))
		if err != nil {
			fmt.Println("Error writing meta file:", err)
			return
		}

	}

	server.ChunkServer(port, "chunks")
}
