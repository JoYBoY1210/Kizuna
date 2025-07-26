package main

import (
	"fmt"
	"kizuna/downloads"
)

func handleDownload(metaPath, outputFile string) {
	err := downloads.DownloadFile(metaPath, outputFile)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	fmt.Println("Download completed!")
}
