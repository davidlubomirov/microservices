SHELL := /bin/bash

CURL_SAMPLES_FILE=curl-samples

build:
	sudo docker build --no-cache -t proxy-server -f _infrastructure/Dockerfile .

run-dev:
	source .env && go run ./cmd/proxy-server/main.go

run-http-samples:
	sudo chmod +x ${CURL_SAMPLES_FILE} && ./${CURL_SAMPLES_FILE}
