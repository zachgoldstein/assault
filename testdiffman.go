package main

import (
	"net/http"
	"io/ioutil"
	"github.com/xeipuuv/gojsonschema"
	"fmt"
	"strings"
)


func main() {

	/*
	Pseudocode for flow
	Digest command line params & config json file to generate http.Request
	Issue request to specified server
	Validate response payload against JSON schema
	Display any difference via command line and HTML output.
	 */

	res := validateResponse(testJSONDoc)
	fmt.Printf("Result of validation %v", res.Valid())
	fmt.Printf("Errors %v", res.Errors())
}

func constructRequest() (req *http.Request, err error) {
	method := "GET"
	reqURL := "localhost:8080/test"
	payload := testJSONDoc
	req, err := http.NewRequest(method, reqURL, strings.NewReader(payload))
}

func issueRequest(req *http.Request)(resPayload []byte, err error) {
	client := http.DefaultClient
	resp, err := client.Do(req)

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func validateResponse(resp string) (res *gojsonschema.Result) {
	schema := gojsonschema.NewStringLoader(testJSONSchema)
	jsonDoc := gojsonschema.NewStringLoader(resp)

	res, _ = gojsonschema.Validate(schema, jsonDoc)

	return res
}

func formatOutput(result *gojsonschema.Result) (output string, err error) {
	//stub to generate CLI output
}

func constructHTML(result *gojsonschema.Result) (output string, err error) {
	//stub to generate HTML output
}


var testJSONSchema = `
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "title": "Product",
    "description": "A product from Acme's catalog",
    "type": "object",
    "properties": {
        "id": {
            "description": "The unique identifier for a product",
            "type": "integer"
        },
        "name": {
            "description": "Name of the product",
            "type": "string"
        },
        "price": {
            "type": "number",
            "minimum": 0,
            "exclusiveMinimum": true
        },
        "tags": {
            "type": "array",
            "items": {
                "type": "string"
            },
            "minItems": 1,
            "uniqueItems": true
        }
    },
    "required": ["id", "name", "price"]
}
`

var testJSONDoc = `
{
    "id": 1,
    "name": "A green door",
    "price": 12.50,
    "tags": ["home", "green"]
}
`
