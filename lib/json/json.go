package json

import (
	"encoding/json"
	"log"
	"os"
)

func ReadFile(name string) map[string]string {
	// TODO check if file is JSON
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatalln(err)
	}

	value := map[string]string{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		log.Fatalln(err)
	}

	return value
}
