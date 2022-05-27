package basicSearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Load(fileName string, m interface{}) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("read file %s: %w", fileName, err)
	}
	err = json.Unmarshal(file, m)
	if err != nil {
		return fmt.Errorf("unmarshal data from %s: %w", fileName, err)
	}

	return nil
}
