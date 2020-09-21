# randomCsvGenerator

## Description
randomCsvGenerator generates a randomized CSV-file with configurable colums and data types with for testing purpose.

## Development
Code is wirtten in Go and needs just common go confogurations like GOROOT.

## Build
<code>go build</code>

## Config
The config folder contains the wordlist.txt and the format.json. These files will be looked up by the program in this folder. So please be sure to place new wordlist.txt and \*.json in here. 

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
The JSNS's root element is colomns. Each column has a name, colType and max element.
The name is a string but not in use in the code. Its only for readability. 
The colType is a string that defines whether to generate a random word or a random number.
Currently recognized colTypes:
* "string"  -> random word
* "int"     -> random number
The max element ist a int that determines wich is the maximal value ob a int-colType. If max=100 it will be generated a number between 0 and 100.
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
