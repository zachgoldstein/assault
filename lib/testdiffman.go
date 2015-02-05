package lib

import (
	"os"
	"io/ioutil"
)

func DoReqDiff() {

	/*
	Pseudocode for flow
	Digest command line params & config json file to generate http.Request
	Issue request to specified server
	Validate response payload against JSON schema
	Display any difference via command line and HTML output.
	 */

	reqOpts, outOpts, err := digestOptions()
	if (err != nil) {
		issueError(err)
	}

	req, err := constructRequest(reqOpts)
	if (err != nil) {
		issueError(err)
	}

	respBody, err := issueRequest(req, createHttpClient())
	if (err != nil) {
		issueError(err)
	}

	res, err := validateResponse(respBody)
	if (err != nil) {
		issueError(err)
	}

	output, err := formatOutput(res)
	if (err != nil) {
		issueError(err)
	}

	_, err = os.Stdout.Write([]byte(output))
	if (err != nil) {
		issueError(err)
	}

	if (outOpts.OutputHTML) {
		html, err := constructHTML(res, string(respBody), outOpts.JSONSchema)
		if (err != nil) {
			issueError(err)
		}

		err = ioutil.WriteFile(outOpts.HTMLOutputLocation, []byte(html), 0644)
		if (err != nil) {
			issueError(err)
		}
	}
}

//issueError will print an error to stdOut that is better formatted than a normal panic
func issueError(err error) {
	panic(err)
}
