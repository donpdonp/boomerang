GOPATH=`pwd` CGO_CFLAGS="-I$GOPATH/leveldb-1.7.0/include" CGO_LDFLAGS="-L$GOPATH/leveldb-1.7.0" go build
