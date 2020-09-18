package jsonReader

import (
	"encoding/json"
	"io/ioutil"
)



type jsonReader struct {
	jsonFile string
}

func NewJsonReader( jsonFile *string ) *jsonReader {
	return &jsonReader{jsonFile:*jsonFile}
}

func (reader *jsonReader) Read(csvFormatObj interface{}) error {

	file, _ := ioutil.ReadFile(reader.jsonFile)
	err := json.Unmarshal(file, csvFormatObj)
	return err
}