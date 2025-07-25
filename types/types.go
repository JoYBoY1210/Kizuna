package types

type MetaFile struct {
	FileName    string         `json:"file_name"`
	FileSize    int64          `json:"file_size"`
	ChunkSize   int            `json:"chunks_size"`
	NumChunks   int            `json:"num_chunks"`
	ChunkHashes map[int]string `json:"chunk_hashes"`
	Peers       []string       `json:"peers"`
}
