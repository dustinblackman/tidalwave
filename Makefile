VERSION := 0.0.1
GOVENDOR_TAG := v1.0.8
LINTER_TAG := V1.0.0

# Creates binary
build:
	go build -ldflags="-X github.com/dustinblackman/tidalwave/cmd.version=$(VERSION)" -o tidalwave *.go

# Creates bash autocomplete file
bashautocomplete:
	go run ./tools/bash-autocomplete/bash.go

# Gets govendor if not found and installs all dependencies
deps:
	@if [ "$$(which govendor)" = "" ]; then \
		go get -v github.com/kardianos/govendor; \
		cd $$GOPATH/src/github.com/kardianos/govendor;\
		git checkout tags/$(GOVENDOR_TAG);\
		go install;\
	fi
	govendor sync

dev:
	which reflex && echo "" || go get github.com/cespare/reflex
	reflex -R '^vendor/' -r '\.go$\' -s -- sh -c 'go run tidalwave.go -server -client'

dev-build:
	which reflex && echo "" || go get github.com/cespare/reflex
	reflex -R '^vendor/' -r '\.go$\' -s -- sh -c 'go build -o tidalwave tidalwave.go && ./tidalwave -server'

# Creates binarys for all available systems in gox and then zips/tars for distribution.
dist: bashautocomplete
	which gox && echo "" || go get github.com/mitchellh/gox
	rm -rf tmp dist
	gox -os="linux windows freebsd" -osarch="darwin/amd64" -output='tmp/{{.OS}}-{{.Arch}}-$(VERSION)/{{.Dir}}' -ldflags="-X github.com/dustinblackman/tidalwave/cmd.version=$(VERSION)"
	mkdir dist

	# Build for Windows
	@for i in $$(find ./tmp -type f -name "tidalwave.exe" | awk -F'/' '{print $$3}'); \
	do \
	  zip -j "dist/tidalwave-$$i.zip" "./tmp/$$i/tidalwave.exe"; \
	done

	# Build for everything else
	@for i in $$(find ./tmp -type f -not -name "tidalwave.exe" | awk -F'/' '{print $$3}'); \
	do \
	  chmod +x "./tmp/$$i/tidalwave"; \
	  tar -zcvf "dist/tidalwave-$$i.tar.gz" --directory="./tmp/$$i" "./tidalwave"; \
	done

	rm -rf tmp

# Creates docker image
# TODO This should change to building with version argument to support docker hub.
# Or have two docker files (which sounds silly)
docker:
	gox -os="linux" -arch="amd64" -output="tidalwave" -ldflags="-X github.com/dustinblackman/tidalwave/cmd.version=$(VERSION)"
	chmod +x ./tidalwave
	docker build -t dustinblackman/tidalwave:latest .
	rm ./tidalwave
	docker push dustinblackman/tidalwave:latest

# Creates easyjson file for parser/parser.go
easyjson:
	easyjson parser/parser.go

# Builds and installs binary. Mainly used from people wanting to install from source.
install: deps
	go install -ldflags="-X github.com/dustinblackman/tidalwave/cmd.version $(VERSION)" *.go

# Setups linter configuration for tests
setup-linter:
	@if [ "$$(which gometalinter)" = "" ]; then \
		go get -v github.com/alecthomas/gometalinter; \
		cd $$GOPATH/src/github.com/alecthomas/gometalinter;\
		git checkout tags/$(LINTER_TAG);\
		go install;\
		gometalinter --install;\
	fi

# Runs tests
test:
	make setup-linter
	gometalinter --vendor --fast --dupl-threshold=100 --cyclo-over=25 --min-occurrences=5 --disable=gas ./...
