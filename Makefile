## build: build golang
.PHONY: build
build:
	go build -o bin/server cmd/server/server.go

## start: run server in development mode
.PHONY: start
start:
	go run cmd/server/server.go

.PHONY: watch
watch:
	CompileDaemon -include=Makefile --build="make build" --command=./bin/server --color=true --log-prefix=false

## pbgen: genrate protobug file
.PHONY: pbgen
pbgen:
	protoc --proto_path=internals/api/v1 --proto_path=thirdParty --go_out=plugins=grpc:pkg/api/v1 --grpc-gateway_out=logtostderr=true:pkg/api/v1 --openapiv2_out=logtostderr=true:swagger beef.proto
	protoc --proto_path=internals/api/v1 \
	--proto_path=thirdParty \
	--go_out=plugins=grpc:pkg/api/v1 \
	--grpc-gateway_out=logtostderr=true:pkg/api/v1 \
	--openapiv2_out=logtostderr=true,allow_merge=true,json_names_for_fields=false:www \
	beef.proto

	mv www/apidocs.swagger.json www/swagger.json

.PHONY: test
test:
	go test ./...

.PHONY: stringer
stringer:
	stringer -type ErrorCode internals/constants/error_code.go
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo