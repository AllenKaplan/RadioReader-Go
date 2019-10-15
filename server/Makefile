build:
	set GOOS = linux
	set GOARCH = amd64
	go build
	docker build -t server .

run:
	docker run radioreader-server -p 14586:14586
