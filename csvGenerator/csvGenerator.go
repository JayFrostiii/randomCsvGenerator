package csvGenerator

import (
	"encoding/csv"
	"os"
)

type Column struct {
	Name    string `json:"name"`
	ColType string `json:"colType"`
	Max int `json:"max"`
}

type CvsFormat struct {
	Columns []Column `json:"column"`
}

type csvGenerator struct {
	count     int
	separator rune
	format    CvsFormat
	csvWriter *csv.Writer
}

func NewCsvGenerator(csvFileName *string, count *int, separator *string) *csvGenerator {
	if csvFileName == nil || count == nil || separator == nil {
		return nil
	}

	file, err := os.Create(*csvFileName)
	if err != nil {
		return nil
	}

	writer := csv.NewWriter(file)
	if writer == nil {
		return nil
	}

	r := []rune(*separator)

	return &csvGenerator{count: *count, separator: r[0], csvWriter: writer}
}

func (gen *csvGenerator) GetFormat() *CvsFormat {
	return &gen.format
}

func (gen *csvGenerator) GenerateData() []error {
	var errorList []error
	gen.csvWriter.Comma = gen.separator

	for i := 1; i <= gen.count; i++ {
		if err := gen.csvWriter.Write(gen.generateDataset()); err != nil {
			errorList = append(errorList,err)
		}
	}

	return errorList
}

func (gen csvGenerator) generateDataset() []string {
	var dataset []string

	for _, col := range gen.format.Columns {
		switch col.ColType {
		case "string":
			dataset = append(dataset, randomWord())
		case "int":
			dataset = append(dataset, randomInt(col.Max))
		}
	}

	return dataset
}
