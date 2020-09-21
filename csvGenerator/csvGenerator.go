package csvGenerator

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Column struct {
	Name    string `json:"name"`
	ColType string `json:"colType"`
	Max     int    `json:"max"`
}

type CvsFormat struct {
	Columns []Column `json:"column"`
}

type csvGenerator struct {
	count     int
	separator rune
	Format    CvsFormat
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
	return &gen.Format
}

func (gen *csvGenerator) GenerateData() []error {
	var errorList []error
	gen.csvWriter.Comma = gen.separator

	for i := 1; i <= gen.count; i++ {
		dataset := gen.generateDataset()
		fmt.Println(dataset)
		if err := gen.csvWriter.Write(dataset); err != nil {
			errorList = append(errorList, err)
		}
	}

	gen.csvWriter.Flush()
	return errorList
}

func (gen csvGenerator) generateDataset() []string {
	var dataset []string

	for _, col := range gen.Format.Columns {
		switch col.ColType {
		case "string":
			dataset = append(dataset, randomWord())
		case "int":
			dataset = append(dataset, randomInt(col.Max))
		}
	}

	return dataset
}
