#! /usr/bin/make
ifeq ($(OS),Windows_NT)
	BUILD_TARGET_FILES = goacontext.exe main.go
else
	BUILD_TARGET_FILES ?= goacontext main.go
endif

.DEFAULT_GOAL := prepare

all: cleandep depend clean mkdir precompile gen precompile build build-console

prepare: cleandep depend clean mkdir precompile gen precompile

depend:
	@glide install

cleandep:
	@rm -rf vendor

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf swagger
	@rm -rf assets

mkdir:
	@mkdir swagger

confinit:
	@cp config/local.toml.example config/local.toml 
	@cp config/test.toml.example config/test.toml 

gen:
	@goagen app -d goa-context-sample/design
	@goagen swagger -d goa-context-sample/design
	@goagen controller -d goa-context-sample/design --pkg controllers -o controllers
	@touch app/.gitkeep swagger/.gitkeep

precompile:
	@go-bindata -pkg=swaggerassets -o=swaggerassets/swagger.go swagger/... swaggerui/...
	@go-bindata -pkg=assets -o=assets/bindata.go config/...

build:
	@go build -o $(BUILD_TARGET_FILES)

build-console:
	@go build -o $(BUILD_TARGET_FILES_CONSOLE)

go-run:
	@go run main.go

run: precompile go-run

test:
	@go test -tags=test goa-context-sample/controllers/... -cover