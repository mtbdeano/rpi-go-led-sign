COMMIT=$(shell git rev-parse --short=8 HEAD)
VERSION=$(shell git describe --tags)
BRANCH=$(shell git branch | grep \* | cut -d ' ' -f2)
PACKAGE=github.com/mtbdeano/rpi-go-led-sign

BINARY=rpi-go-led-sign

all: ${BINARY}


$(BINARY): 
	go build -ldflags "-w -s -X ${PACKAGE}/cmd.Version=${VERSION} -X ${PACKAGE}/cmd.Branch=${BRANCH} -X ${PACKAGE}/cmd.Commit=${COMMIT}"

clean:
	rm -f ${BINARY}
