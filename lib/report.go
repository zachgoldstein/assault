package lib

import (
	"github.com/xeipuuv/gojsonschema"
	"fmt"
	"io/ioutil"
	"html/template"
	"bytes"
)

func validateResponse(resp []byte) (res *gojsonschema.Result, err error) {
	schema := gojsonschema.NewStringLoader(testJSONSchema)
	jsonDoc := gojsonschema.NewStringLoader(string(resp))

	return gojsonschema.Validate(schema, jsonDoc)
}

func formatOutput(res *gojsonschema.Result) (output string, err error) {
	if (res.Valid()) {
		output = "The response is valid! You rock bud."
	} else {
		for index, err := range res.Errors() {
			output = output + fmt.Sprintf("Error #%v: %v \n", index+1, err)
		}
	}
	return output, nil
}

type HTMLData struct {
	Title string
	Errors []gojsonschema.ResultError
	JSONSchema string
	JSON	string
	JSONIsValid bool
}

func constructHTML(res *gojsonschema.Result, JSON string, JSONSchema string) (output string, err error) {
	templateFile, err := ioutil.ReadFile("./template.html")
	if (err != nil) {
		return output, err
	}

	templ, err := template.New("html").Parse(string(templateFile))
	if (err != nil) {
		return output, err
	}

	writeBuffer := bytes.NewBuffer([]byte{})

	data := HTMLData {
		Title : "A Empty Title",
		Errors : res.Errors(),
		JSON : JSON,
		JSONSchema: JSONSchema,
	}
	data.JSONIsValid = res.Valid()

	err = templ.ExecuteTemplate(writeBuffer, "html", data)
	if (err != nil) {
		return output, err
	}

	htmlBytes, err := ioutil.ReadAll(writeBuffer)
	if (err != nil) {
		return output, err
	}

	//insert in

	return string(htmlBytes), err
}
