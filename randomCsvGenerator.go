package main

import (
	"flag"
	"fmt"
	"github.com/JayFrostiii/randomCsvGenerator/csvGenerator"
	"github.com/JayFrostiii/randomCsvGenerator/jsonReader"
)

var CONFIG_PATH = "config/"

func main() {
	fileJSONFormatPtr := flag.String("jsonFile", "format.json", "fileJSON")
	csvfilePtr := flag.String("csvFile", "test.csv", "fileCSV")
	separatorPtr := flag.String("separator", ";", "Separator for each cvs dataset")
	datasetCountPtr := flag.Int("count", 1, "Amount of Dataset that will be created")
	flag.Parse()

	fmt.Println("START")
	defer fmt.Println("End")

	generateCSV(fileJSONFormatPtr, csvfilePtr, separatorPtr, datasetCountPtr)
}

func generateCSV(fileJSONFormat, csvfile, separator *string, datasetCount *int) {
	fmt.Println("Initializing...")
	if err := csvGenerator.InitRandomizer(CONFIG_PATH); err != nil {
		fmt.Println(err)
		return
	}

	generator := csvGenerator.NewCsvGenerator(csvfile, datasetCount, separator)
	reader := jsonReader.NewJsonReader(fileJSONFormat)

	fmt.Println("Reading JSON...")
	if err := reader.Read(CONFIG_PATH, generator.GetFormat()); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Generating CSV...")
	multErr := generator.GenerateData()
	for err := range multErr {
		fmt.Println(err)
	}
}
