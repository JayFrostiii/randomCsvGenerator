package main

import (
	"flag"
	"fmt"
	"github.com/JayFrostiii/randomCsvGenerator/csvGenerator"
	"github.com/JayFrostiii/randomCsvGenerator/jsonReader"
)

type JsonReader interface{
	Read(csvFormatObj interface{}) error
}

type CsvGenerator interface{
	GetFormat() *csvGenerator.CvsFormat
	GenerateData()
}

func main() {
	fileJSONFormatPtr := flag.String("jsonFile", "test.json", "fileJSON")
	csvfilePtr := flag.String("csvFile", "test.csv", "fileCSV")
	datasetCountPtr := flag.Int("count", 1, "Amount of Dataset that will be created")
	separatorPtr := flag.String("separator", ";", "Separator for each cvs dataset")
	flag.Parse()

	fmt.Println("START...")
	defer fmt.Println("End...")

	if err := initApp(); err != nil {
		return
	}

	generator := csvGenerator.NewCsvGenerator(csvfilePtr, datasetCountPtr, separatorPtr)

	reader := jsonReader.NewJsonReader(fileJSONFormatPtr)

	if err := reader.Read(generator.GetFormat()); err != nil {
		fmt.Println(err)
		return
	}

	generator.GenerateData()

}

func initApp() error {
	return csvGenerator.InitRandomizer()
}