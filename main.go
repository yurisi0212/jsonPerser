package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const uri = "https://pokeapi.co/api/v2/pokemon/magikarp"

func main() {
	json := takeJsonData()
	if !checkJsonFormat(json) {
		panic(errors.New("this file is not json"))
	}
	parseJSON(json)
}

func takeJsonData() string {
	req, err := http.Get(uri)
	if err != nil {
		panic(errors.New("get http error"))
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(errors.New("io read error"))
	}
	req.Body.Close()
	return string(body)
}

func checkJsonFormat(str string) bool {
	if str == "" {
		return false
	}
	if str[0] != '{' {
		return false
	}
	if str[len(str)-1] != '}' {
		return false
	}
	return true
}

func parseJSON(str string) {
	whitespace := 0
	for _, c := range str {
		switch c {
		case '{':
			fmt.Printf("{\n%s", strings.Repeat(" ", whitespace+1))
			whitespace++
			continue
		case '}':
			whitespace--
			fmt.Printf("\n%s}", strings.Repeat(" ", whitespace))
			continue
		case ':':
			fmt.Printf(": ")
			continue
		case ',':
			fmt.Printf(",\n%s", strings.Repeat(" ", whitespace))
			continue
		case '[':
			whitespace++
			fmt.Printf("[\n%s", strings.Repeat(" ", whitespace))

			continue
		case ']':
			whitespace--
			fmt.Printf("\n%s]", strings.Repeat(" ", whitespace))
			continue
		}
		fmt.Printf("%c", c)
	}
}
