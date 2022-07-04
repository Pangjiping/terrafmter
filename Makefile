default: install

install: clean mac copy

clean:
	rm -rf bin/*
	rm -f usr/local/bin/terrafmtter

mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/terrafmtter
	tar czvf bin/terrafmtter_darwin-amd64.tgz bin/terrafmtter
	rm -rf bin/terrafmtter

copy:
	tar -xvf bin/terrafmtter_darwin-amd64.tgz
	mv bin/terrafmtter /usr/local/bin
	