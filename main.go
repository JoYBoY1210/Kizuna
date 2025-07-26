package main

import (
	// "kizuna/files"
	// "kizuna/server"
	// "kizuna/downloads"
	"fmt"
	"os"
	// "kizuna/types"
	// "time"
	"flag"

)

func main() {
	seedCmd:=flag.NewFlagSet("seed", flag.ExitOnError)
	downloadCmd:=flag.NewFlagSet("download", flag.ExitOnError)

	filePath:=seedCmd.String("file", "", "File to seed")
	port:=seedCmd.String("port", "6253", "Port to run the server on")
	peerList:=seedCmd.String("peers", "", "list of peer addresses")
	metaPath:=seedCmd.String("meta", "", "Path to metadata file(optional)")

	downloadMeta:=downloadCmd.String("meta", "", "Path to metadata file")
	outputFile:=downloadCmd.String("output", "output_file", "Output file path")

	if len(os.Args)<2{
		fmt.Println("expected 'seed' or 'download' subcommands")
		os.Exit(1)
	}

	switch os.Args[1]{
	case "seed":
		seedCmd.Parse(os.Args[2:])
		handleSeed(*filePath, *port, *peerList, *metaPath)

	case "download":
		downloadCmd.Parse(os.Args[2:])
		handleDownload(*downloadMeta, *outputFile)

	default:
		fmt.Println("expected 'seed' or 'download' subcommands")
		os.Exit(1)
	}


}
