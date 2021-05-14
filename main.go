package main

import (
	"encoding/json"
	"errors"
	"os"
	"text/template"
)

func usage() string {
	return "template: JSON_FILE SOURCE_FILE"
}

func exit(s string) {
	os.Stderr.WriteString(s + "\n")
	os.Exit(0)
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		exit(usage())
	}
	data, err := unmarshalJSON(args[0])
	if err != nil {
		exit(err.Error())
	}
	sourcePath := args[1]
	t, err := template.New(sourcePath).ParseFiles(sourcePath)
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
