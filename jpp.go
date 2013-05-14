package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("error:", err)
	}
	var i interface{}
	err = json.Unmarshal(bytes, &i)
	if err != nil {
		fmt.Println("error:", err)
	}
	b, err := json.MarshalIndent(i, "", "   ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
