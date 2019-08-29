package tfgen

import (
	"log"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extractResourceName(arn string) string {
	components, err := awsarn.Parse(arn)
	if err != nil {
		panic(err)
	}
	return components.Resource
}

func resourceName(resourceName string) string {

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(resourceName, "_")

	return processedString
}
