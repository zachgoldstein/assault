package lib

import (
	"testing"
	c "github.com/smartystreets/goconvey/convey"
	"os"
	"io/ioutil"
)

func TestFormatOutput(t *testing.T) {
	c.Convey("With a valid result", t, func() {
		res, err := validateResponse([]byte(testJSONDoc))
		c.So(err, c.ShouldBeNil)
		c.So(res.Valid(), c.ShouldBeTrue)

		c.Convey("I can generate simple output that can be printed to stdOut and indicates it's valid", func() {
			output, err := formatOutput(res)
			c.So(err, c.ShouldBeNil)
			c.So(output, c.ShouldContainSubstring, "valid")

			_, err = os.Stdout.Write([]byte(output))
			c.So(err, c.ShouldBeNil)
		})
	})

	c.Convey("With an invalid result", t, func() {
		res, err := validateResponse([]byte(testFailJSONDoc))
		c.So(err, c.ShouldBeNil)
		c.So(res.Valid(), c.ShouldBeFalse)

		c.Convey("I can generate simple output that can be printed to stdOut and includes errors", func() {
			output, err := formatOutput(res)
			c.So(err, c.ShouldBeNil)
			c.So(output, c.ShouldContainSubstring, "Error")

			_, err = os.Stdout.Write([]byte(output))
			c.So(err, c.ShouldBeNil)
		})
	})
}

func TestConstructHTML(t *testing.T) {

	c.Convey("With a valid result", t, func() {
		res, err := validateResponse([]byte(testJSONDoc))
		c.So(err, c.ShouldBeNil)
		c.So(res.Valid(), c.ShouldBeTrue)

		c.Convey("I can construct an html view of it", func() {
			html, err := constructHTML(res, testJSONDoc, testJSONSchema)
			c.So(err, c.ShouldBeNil)
			c.So(html, c.ShouldContainSubstring, "A Empty Title")

			err = ioutil.WriteFile("./testValid.html", []byte(html), 0644)
			c.So(err, c.ShouldBeNil)
		})
	})

	c.Convey("With a invalid result", t, func() {
		res, err := validateResponse([]byte(testFailJSONDoc))
		c.So(err, c.ShouldBeNil)
		c.So(res.Valid(), c.ShouldBeFalse)

		c.Convey("I can construct an html view of it", func() {
			html, err := constructHTML(res, testFailJSONDoc, testJSONSchema)
			c.So(err, c.ShouldBeNil)
			c.So(html, c.ShouldContainSubstring, "A Empty Title")

			err = ioutil.WriteFile("./testError.html", []byte(html), 0644)
			c.So(err, c.ShouldBeNil)
		})
	})
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
        "stringNumber": {
            "description": "A number from 0-9",
            "type": "string",
            "pattern":"[0-9]"
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
    "stringNumber": "3",
    "price": 12.50,
    "tags": ["home", "green"]
}
`

var testFailJSONDoc = `
{
    "id": "woah wtf mate",
    "name": "A green door",
    "stringNumber": "test",
    "price": "not a number",
    "tags": ["home", "green"]
}
`
