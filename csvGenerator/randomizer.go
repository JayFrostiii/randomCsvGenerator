package csvGenerator

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
)

var wordList []string

func InitRandomizer(configPath string) error {
	data, err := ioutil.ReadFile(configPath + "wordlist.txt")
	if err != nil {
		return err
	}

	wordList = strings.Fields(string(data))

	return nil
}

func randomWord() string {
	index := rand.Intn(len(wordList) - 1)
	return wordList[index]
}

func randomInt(max int) string {
	return strconv.Itoa(rand.Intn(max))
}
