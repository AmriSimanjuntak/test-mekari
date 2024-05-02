NAME := /test-mekari
TAG := $$(git log -1 --pretty=%h)
IMG := ${NAME}:${TAG}
LATEST := ${NAME}:nc-latest
DST_PATH := test-mekari

buildimage:
	@docker build --platform=linux/amd64 -t ${IMG} .
	@docker tag ${IMG} ${LATEST}

pushimage: buildimage
	@docker push ${LATEST}

runapp: buildimage
	@docker-compose up -d

buildapp:
	rm -rf ${DST_PATH}
	swag init -g main.go --parseDependency
	go build .
.PHONY = run

run: buildapp
	go run test-mekari