package jsonReader

import (
	"encoding/json"
	"io/ioutil"
)

type jsonReader struct {
	jsonFile string
}

func NewJsonReader(jsonFile *string) *jsonReader {
	return &jsonReader{jsonFile: *jsonFile}
}

func (reader *jsonReader) Read(configPath string, csvFormatObj interface{}) error {
	file, err := ioutil.ReadFile(configPath + reader.jsonFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, csvFormatObj)
}
