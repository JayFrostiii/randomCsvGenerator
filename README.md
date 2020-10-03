# randomCsvGenerator

## Description
randomCsvGenerator generates a randomized CSV-file with configurable colums and data types for testing purpose.

## Development
The code base is wirtten in Go and needs just common go configurations like GOROOT.

## Build
<code>go build</code>

## Config
The config folder contains the wordlist.txt and the format.json. These files will be looked up by the program in this folder. So please make sure to place the new wordlist.txt and \*.json in here. 

### wordlist.txt
The list contains a number of words. One per line. Seperated by newline "\n".
Clear:
<code>
  bla1\n
  bla2\n
</code>
Interpreted:
<code>
  bla1
  bla2
</code>

### format.json
The JSON's root element is columns. Each column has a name, colType and max element.
The name is a string but not in use in the code. Its only for readability. 
The colType is a string that defines whether to generate a random word or a random number.
Currently recognized colTypes:
* "string"  -> random word
* "int"     -> random number
The max element is a int that determines wich is the maximal value of an int-colType. If max=100, it will generate a number between 0 and 100.
Example JSON:
<code>
{
  "column": [
    {
      "name": "name",
      "colType": "string",
      "max": 0
    },
    {
      "name": "age",
      "colType": "int",
      "max": 100
    }
  ]
}
</code>

## Usage
The executable can be called with following flags:
* -jsonFile \<name of JSON containig format\>           DEFAULT: format.json
* -csvFile \<name of CSV file that will be created\>    DEFAULT: test.csv
* -separator \<symbol for separation in CSV\>           DEFAULT: ;
* -count \<number of datasets to create randomly\>      DEFAULT: 1

e.g.
./randomCsvGenerator.exe -jsonFile myFormat.json -csvFile myTest.csv -separator , -count 100
