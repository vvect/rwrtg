package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"rwrtg/models"
	"strings"
	"text/template"
)

func escapeJS(js string) string{
	replaceMap := map[string]string{
		"$": "_inner_",
		".": "_",
		"[": "_arr_",
		";": "_semicolon_",
		"-": "_dash_",
	}
	for key, val := range replaceMap{
		js = strings.ReplaceAll(js, key, val)
	}
	return js
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}


	var templateFile string
	var profileFile string

	flag.StringVar(&templateFile, "t", "advanced.logging.js", "Specify the template to use, eg advanced.logging.js")
	flag.StringVar(&profileFile, "p", "static.rwrt.json", "Specify the input profile to use (JSON)")
	flag.Parse()

	dat, err := os.ReadFile(profileFile)
	if err != nil {
		println("Could not read profile")
		panic(err)
	}
	var profile models.StaticProfile

	if !filepath.IsAbs(templateFile){
		templateFile = cwd + "/templates/" + templateFile
	}

	templateText, err := os.ReadFile(templateFile)
	if err != nil {
		println("Could not read template")
		panic(err)
	}

	err = json.Unmarshal(dat, &profile)
	if err != nil {
		println("Could not deserialize JSON object, is it a profile?")
		panic(err)
	}

	// Create a new template, passing in our functions as a map of string -> func
	tmpl, _ := template.New("test").Funcs(template.FuncMap{
		"escapeJS": func(js string) string {
			return escapeJS(js)
		},
		"getOverloadString": func(a []models.TypeMap) string {
			var args []string
			for _, arg := range a {
				args = append(args, fmt.Sprintf("'%v'", arg.ClassName))
			}
			return strings.Join(args,", ")
		},
		"getTypedArguments": func(a []models.TypeMap) string {
			var args []string
			i := 0
			for _, arg := range a {
				args = append(args, fmt.Sprintf("%v_%d", escapeJS(arg.ClassName), i))
				i += 1
			}
			return strings.Join(args,", ")
		},
	}).Parse(string(templateText))
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, profile)
	if err != nil{
		println(err)
	}

	fmt.Println(tpl.String())
}