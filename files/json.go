package files

import (
	"encoding/json"
	"os"
)

func JSON(path string, out any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(out)
}
