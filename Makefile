BIN := cloudwatch-logger

$(BIN): vendor main.go
	go build -o $(BIN) main.go

.PHONY: vendor
vendor: govendor vendor/vendor.json
	govendor sync

.PHONY: govendor
govendor:
	go get github.com/kardianos/govendor

.PHONY: clean
clean:
	rm -f $(BIN)
