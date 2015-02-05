FILES=./lib/*.go

fmt:
	go fmt ${FILES}

deps:
	go get github.com/smartystreets/goconvey
	go get github.com/xeipuuv/gojsonschema

test:
	go test ${FILES} -v

live-test:
	goconvey

doc:
	pkill godoc; godoc -http=":7080" &