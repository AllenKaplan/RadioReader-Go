build:
	set GOOS = linux
	set GOARCH = amd64
	go build
	docker build -t client .

run:
	docker run radioreader-client -p 14586:14586
