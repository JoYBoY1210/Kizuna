package files

import (
	"encoding/json"
	"os"
	"kizuna/types"
)



func NewMetaFile(meta types.MetaFile, outputPath string) error {
	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, data, 0644)
}
