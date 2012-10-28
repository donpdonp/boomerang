export GOPATH=$(CURDIR)
export CGO_CFLAGS=-I$(CURDIR)/leveldb-1.7.0/include
export CGO_LDFLAGS=-L$(CURDIR)/leveldb-1.7.0

all: leveldb-1.7.0/libleveldb.a src/code.google.com/p/goweb src/github.com/jmhodges/levigo boomerang60
	@echo building in $$GOPATH
	go build

src/code.google.com/p/goweb:
	go get code.google.com/p/goweb

src/github.com/jmhodges/levigo:
	go get github.com/jmhodges/levigo

leveldb-1.7.0/libleveldb.a: leveldb-1.7.0.tar.gz
	$(MAKE) -C leveldb-1.7.0

leveldb-1.7.0:
	@echo downloading leveldb 1.7.0
	wget http://leveldb.googlecode.com/files/leveldb-1.7.0.tar.gz	
	tar zxf leveldb-1.7.0.tar.gz 
