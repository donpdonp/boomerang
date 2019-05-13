export GOPATH=$(CURDIR)

all: src/github.com/stretchr/goweb src/github.com/jmhodges/levigo boomerang

boomerang: web.go
	@echo building in $$GOPATH
	go build

src/github.com/stretchr/goweb:
	go get github.com/stretchr/goweb

src/github.com/jmhodges/levigo:
	go get github.com/jmhodges/levigo
