PROJECT_NAME = hownow
VERSION=$(shell cat VERSION)

BUILD_IMAGE_NAME = $(PROJECT_NAME)-build

APP_DIR ?= $(PWD)
BIN_DIR=$(APP_DIR)/bin
RELEASE_DIR=$(APP_DIR)/release
VENDOR_DIR=$(APP_DIR)/vendor
TMP_DIR=$(APP_DIR)/tmp

BUILD_ARGS=-i -ldflags=""

DARWIN_RELEASE=${PROJECT_NAME}_darwin_$(VERSION)
LINUX_RELEASE=${PROJECT_NAME}_linux_$(VERSION)

GITHUB_RELEASE_ARGS=--user ${GITHUB_USERNAME} --repo ${GITHUB_REPO} --tag ${VERSION}

.PHONY: create_build_image deps update_deps make_dirs test run clean

default: make_dirs test
	go build ${BUILD_ARGS} -o $(BIN_DIR)/$(PROJECT_NAME) ./cmd

run: test
	go run main.go version.go $(filter-out $@, $(MAKECMDGOALS))

release: make_dirs test
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo ${BUILD_ARGS} -o ${RELEASE_DIR}/${DARWIN_RELEASE} ./cmd
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ${BUILD_ARGS} -o ${RELEASE_DIR}/${LINUX_RELEASE} ./cmd
	upx ${RELEASE_DIR}/${DARWIN_RELEASE}
	upx ${RELEASE_DIR}/${LINUX_RELEASE}

release_upload: release
	git tag $(VERSION)
	git push
	git push --tags

	github-release release ${GITHUB_RELEASE_ARGS}
	github-release upload ${GITHUB_RELEASE_ARGS} --name ${DARWIN_RELEASE} --file ${RELEASE_DIR}/${DARWIN_RELEASE}
	github-release upload ${GITHUB_RELEASE_ARGS} --name ${LINUX_RELEASE} --file ${RELEASE_DIR}/${LINUX_RELEASE}

travis_test:
	go test -race -coverprofile=coverage.txt -covermode=atomic

test: deps
	go tool vet -tests *.go
	errcheck -ignoretests
	go test -coverprofile ./tmp/coverage.out

quick_test:
	go test -coverprofile ./tmp/coverage.out
	go tool cover -html ./tmp/coverage.out

report: test
	mkdir -p reports
	go get -u github.com/robotvert/goreportcard
	goreporter -p ../hownow -r ../hownow/reports/ -e vendor
	go tool cover -html ./tmp/coverage.out

deps: make_dirs
	go get github.com/kardianos/govendor
	go get github.com/kisielk/errcheck
	go get github.com/aktau/github-release
	go get honnef.co/go/tools/cmd/megacheck
	govendor fetch +missing
	go generate -v -x ./ ./cmd ./commands

update_deps:
	go get -u github.com/kardianos/govendor
	go get -u github.com/kisielk/errcheck
	go get -u github.com/aktau/github-release
	go get -u honnef.co/go/tools/cmd/megacheck
	govendor update +e

make_dirs:
	@mkdir -p ${BIN_DIR} ${RELEASE_DIR} ${VENDOR_DIR} ${TMP_DIR}

clean:
	rm -rf $(BIN_DIR) $(RELEASE_DIR)
