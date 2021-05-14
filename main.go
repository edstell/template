package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"text/template"
)

func usage() string {
	return "template: JSON_FILE [SOURCE_FILE]"
}

func exit(s string) {
	os.Stderr.WriteString(s + "\n")
	os.Exit(0)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		exit(usage())
	}
	data, err := unmarshalJSON(args[0])
	if err != nil {
		exit(err.Error())
	}
	f := os.Stdin
	if len(args) >= 2 {
		f, err = os.Open(args[1])
		if err != nil {
			exit(err.Error())
		}
		defer f.Close()
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		exit(err.Error())
	}
	t, err := template.New("").Parse(string(b))
	if err != nil {
		exit(err.Error())
	}
	if err := t.Execute(os.Stdout, data); err != nil {
		exit(err.Error())
	}
}

func unmarshalJSON(path string) (interface{}, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var data interface{}
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		return nil, errors.New("invalid JSON: " + err.Error())
	}
	return data, nil
}
