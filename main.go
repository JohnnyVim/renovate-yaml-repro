package main

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func main() {
	msg := cases.Title(language.English).String("Renovate is cool!")
	yaml.Marshal(msg)
}
