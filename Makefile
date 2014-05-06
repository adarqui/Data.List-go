all:
	go build

test:
	go test -bench=".*" -test.v

clean:
	echo "clean"
