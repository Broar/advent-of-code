YEAR ?= 2020
DAY ?= 1

.PHONY: run-challenge
run-challenge:
	go run $(YEAR)/$(DAY)/main.go --input-filepath $(YEAR)/$(DAY)/input.txt