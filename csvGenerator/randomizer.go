package csvGenerator

import (
	"io/ioutil"
	"math/rand"
	"strings"
)

var wordList []string

func InitRandomizer() error {
	data, err := ioutil.ReadFile("wordlist.txt")
	if err != nil {
		return err
	}

	wordList = strings.Split(string(data), "\n")

	return nil
}

func randomWord() string {
	index := rand.Intn(len(wordList)-1)
	return wordList[index]
}

func randomInt(max int) string {
	return string(rand.Intn(max))
}