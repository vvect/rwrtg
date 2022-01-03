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

func escapeJS(js string) string {
	replaceMap := map[string]string{
		"$": "_inner_",
		".": "_",
		"[": "_arr_",
		";": "_semicolon_",
		"-": "_dash_",
	}
	for key, val := range replaceMap {
		js = strings.ReplaceAll(js, key, val)
	}
	return js
}

func main() {
	pathSlice := strings.Split(os.Args[0], "/")
	cwd := strings.Join(pathSlice[:len(pathSlice)-1], "/")

	var templateFile string
	var profileFile string

	flag.StringVar(&templateFile, "t", "android.advanced.logging.js", "Specify the template to use, eg android.advanced.logging.js")
	flag.StringVar(&profileFile, "p", "static.rwrt.json", "Specify the input profile to use (JSON)")
	flag.Parse()
	dat, err := os.ReadFile(profileFile)
	if err != nil {
		println("Could not read profile")
		panic(err)
	}

	var profile models.Profile
	err = json.Unmarshal(dat, &profile)
	if err != nil {
		println("Could not deserialize JSON object, is it a profile?")
		panic(err)
	}

	if templateFile == "android.advanced.logging.js" && profile.Metadata.Runtime == "objc" {
		templateFile = "ios.logging.js"
	}
	if !filepath.IsAbs(templateFile) {
		templateFile = cwd + "/templates/" + templateFile
	}

	templateText, err := os.ReadFile(templateFile)
	if err != nil {
		println("Could not read template")
		panic(err)
	}

	// Create a new template, passing in our functions as a map of string -> func
	tmpl, _ := template.New("test").Funcs(template.FuncMap{
		"escapeJS": func(js string) string {
			return escapeJS(js)
		},
		"getOverloadString": func(a []models.AndroidParameter) string {
			var args []string
			for _, arg := range a {
				args = append(args, fmt.Sprintf("'%v'", arg.ClassName))
			}
			return strings.Join(args, ", ")
		},
		"getTypedArguments": func(a []models.AndroidParameter) string {
			var args []string
			i := 0
			for _, arg := range a {
				args = append(args, fmt.Sprintf("%v_%d", escapeJS(arg.ClassName), i))
				i += 1
			}
			return strings.Join(args, ", ")
		},
		"getReturnValueType": func(c models.AndroidClass, m models.AndroidMethod, a *[]string, i int) string {
			if m.Name == "$init" {
				return c.Name
			}
			if a != nil {
				return (*m.ReturnTypes)[i]
			} else {
				return "void"
			}
		},
		"hasReturnValue": func(name string, a *[]string) bool {
			if name == "$init" {
				return true
			}
			return a != nil
		},
		"getIOSClass": func(fqn string) string {
			var classStart, classEnd int
			classStart = strings.Index(fqn, "[")
			classEnd = strings.Index(fqn, " ")
			return fqn[classStart+1 : classEnd]
		},
		"getIOSMethod": func(fqn string) string {
			var classEnd, methodEnd int
			classEnd = strings.Index(fqn, " ")
			methodEnd = strings.Index(fqn, "]")

			return fqn[0:1] + " " + fqn[classEnd+1:methodEnd]
		},
	}).Parse(string(templateText))
	var tpl bytes.Buffer
	if profile.Metadata.Runtime == "java" {
		var tap models.TempAndroidProfile
		err = json.Unmarshal(dat, &tap)
		androidProfile := tap.GetAndroidProfile()
		if err != nil {
			println("ERROR:")
			println(err.Error())
			return
		}
		err = tmpl.Execute(&tpl, androidProfile)
		if err != nil {
			println("ERROR:")
			println(err.Error())
			return
		}
	} else if profile.Metadata.Runtime == "objc" {
		var profile models.IOSProfile
		err = json.Unmarshal(dat, &profile)
		if err != nil {
			println("ERROR:")
			println(err.Error())
			return
		}
		err = tmpl.Execute(&tpl, profile)
		if err != nil {
			println("ERROR:")
			println(err.Error())
			return
		}
	}

	fmt.Println(tpl.String())
}
