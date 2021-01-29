package word

import (
	"encoding/json"
	"io/ioutil"
)

// GetWords get the words from the words file
func GetWords(wordsFile string) []string {
	bytes, err := ioutil.ReadFile(wordsFile)
	checkErr(err)
	var res []string
	err = json.Unmarshal(bytes, &res)
	checkErr(err)
	return res
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
