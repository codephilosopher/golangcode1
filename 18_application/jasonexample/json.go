package jasonexample

import (
	"encoding/json"
	"fmt"
)

//JSONExample returns json examples
func JSONExample() {
	type Message struct {
		Head string
		Body string
		Time int64
	}

	a := Message{
		Head: "Hi there",
		Body: "Here is your body",
		Time: 22222,
	}
	b := []Message{a}
	// marshal data to json
	v, err := json.Marshal(b)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(v))

	//unmarshal
	jsonBlob := []byte(`[{"Head":"Hi there","Body":"Here is your body","Time":22222}]`)
	var message []Message
	errs := json.Unmarshal(jsonBlob, &message)
	if errs != nil {
		fmt.Println(errs)

	}
	for _, v := range message {
		fmt.Println(v.Head)
	}
}
