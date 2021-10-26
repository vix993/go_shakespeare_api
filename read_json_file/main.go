package read_json_file

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJsonFile() {
	jsonFile, err := os.Open("shakespeare.json")

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var quotes []string

	json.Unmarshal(byteValue, &quotes)

	for i := 0; i < len(quotes); i++ {
		println(quotes[i])
	}
}
