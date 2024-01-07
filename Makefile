run:
	go run main.go
build:
	GOOS=windows GOARCH=amd64 go build -o WS-progrma-ofppt.exe
	GOOS=linux GOARCH=amd64 go build -o linux-program-ofppt
	GOOS=darwin GOARCH=amd64 go build -o macos-program-ofppt
push:
	git add . && git status && git commit -m "version-beta" && git push origin master && gh repo list
