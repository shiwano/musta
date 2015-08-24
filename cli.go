package main

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/hoisie/mustache"
	"os"
	"regexp"
	"strings"
)

// Run executes the command.
func Run(args []string, jsonString string, templateFileName string, templateData string) {
	jsonContainer := ParseArgs(args, jsonString)
	template := createTemplate(templateFileName, templateData)
	output := template.Render(jsonContainer.Data())

	if templateFileName == "" {
		fmt.Println(output)
	} else {
		fmt.Print(output)
	}
}

// ParseArgs parses arguments to JSON data
func ParseArgs(args []string, jsonString string) *gabs.Container {
	jsonContainer := createJSONContainer(jsonString)
	argPattern := regexp.MustCompile("([^=]+)=([^=]+)")

	for _, arg := range args {
		group := argPattern.FindStringSubmatch(arg)
		if len(group) != 3 {
			continue
		}

		values := strings.Split(strings.Trim(group[2], ","), ",")
		if len(values) == 1 {
			jsonContainer.SetP(group[2], group[1])
			continue
		}

		arrayItems := make([]map[string]interface{}, len(values))
		for index, value := range values {
			arrayItems[index] = map[string]interface{}{
				"index": index,
				"first": index == 0,
				"last":  index == len(values)-1,
				"value": value,
			}
		}
		jsonContainer.SetP(arrayItems, group[1])
	}

	return jsonContainer
}

// Fatalf writes to stderr and exits.
func Fatalf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "\n  %s\n\n", fmt.Sprintf(msg, args...))
	os.Exit(1)
}

func createTemplate(templateFileName string, templateData string) *mustache.Template {
	if templateFileName != "" {
		template, err := mustache.ParseFile(templateFileName)
		if err != nil {
			Fatalf("Failed to load a template file: ", err)
		}
		return template
	} else if templateData != "" {
		template, err := mustache.ParseString(templateData)
		if err != nil {
			Fatalf("Failed to parse a template data: ", err)
		}
		return template
	} else {
		Fatalf("No template file or template data")
		return nil
	}
}

func createJSONContainer(jsonString string) *gabs.Container {
	if jsonString == "" {
		return gabs.New()
	}

	container, err := gabs.ParseJSON([]byte(jsonString))
	if err != nil {
		Fatalf("Failed to parse piped JSON string: %s", err)
	}

	return container
}
