build:
	set GOOS = linux
	set GOARCH = amd64
	go build
	docker build -t reader .

run:
	docker run radioreader-reader -p 14586:14586
