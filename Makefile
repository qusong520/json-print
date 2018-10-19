binary_name = json-print

build:
	go build -o ${binary_name} -ldflags "-X main.version=`cat ./version`" .

install:
	go install -ldflags "-X main.version=`cat ./version`" .
