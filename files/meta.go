package files

import (
	"encoding/json"
	"os"
	"kizuna/types"
)



func NewMetaFile(meta types.MetaFile,filepath string) error {
	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}
