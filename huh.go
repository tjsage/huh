package huh

import (
	"encoding/json"
	"fmt"
)

func Simple(description, value interface{}) {
	fmt.Printf("%s: %+v\n", description, value)
}

func Message(message string) {
	fmt.Println(message)
}

func Json(description, item interface{}) {
	data, err := json.MarshalIndent(&item, "", "\t")

	if err != nil {
		fmt.Printf("%s: %s\n", description, err.Error())
	}

	fmt.Printf("%s: \n %s\n", description, string(data))
}
