export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)
APP_NAME:="smsmanager"
VERSION:="1.0.1"
ARCH:="amd64"
BUILD_DIR:=$(APP_NAME)_$(VERSION)_$(ARCH)

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: ### run app
	go mod tidy && go mod download && \
	go run ./cmd/app -config_file=./config/smsmanager.yml
.PHONY: run

linter-golangci: ### check by golangci linter
	golangci-lint -v run
.PHONY: linter-golangci

build: ### build app
	mkdir -p $(BUILD_DIR)/usr/local/bin
	mkdir -p $(BUILD_DIR)/etc/smsmanager
	mkdir -p $(BUILD_DIR)/lib/systemd/system
	CGO_ENABLED=1 go build -ldflags "-s -w" -o $(BUILD_DIR)/usr/local/bin/smsmanager -v ./cmd/app
	cp ./config/smsmanager.yml $(BUILD_DIR)/etc/smsmanager
	cp ./config/smsmanager.service $(BUILD_DIR)/lib/systemd/system

deb-package: build ### build debian package
	mkdir -p $(BUILD_DIR)/DEBIAN
	cp -r ./DEBIAN/* $(BUILD_DIR)/DEBIAN
	sed -i "s/<VERSION>/$(VERSION)/g" $(BUILD_DIR)/DEBIAN/control
	sed -i "s/<ARCH>/$(ARCH)/g" $(BUILD_DIR)/DEBIAN/control
	dpkg-deb --build --root-owner-group $(BUILD_DIR)

install: build ### build and install app
	cp -r $(BUILD_DIR)/etc/smsmanager /etc/
	cp $(BUILD_DIR)/lib/systemd/system /lib/systemd/system

uninstall: ### uninstall app
	rm -rf /etc/smsmanager
	rm -f /lib/systemd/system/smsmanager.service

clean: ### delete temp files
	rm -rf $(BUILD_DIR)
	rm -f $(BUILD_DIR).deb
	rm -f ./db.sql
