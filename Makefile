flutter_key:
	go run main.go -platform flutter -asset key -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export -package firebasex/analytics

flutter_code:
	go run main.go -platform flutter -asset code -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export -package firebasex/analytics

flutter_gen:
	go run main.go -platform flutter -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export -package firebasex/analytics

ios_key:
	go run main.go -platform ios -asset key -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export

ios_code:
	go run main.go -platform ios -asset code -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export

build_linux:
	env GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o ./binary/linux/analyticsgen github.com/prongbang/analyticsgen

build_macos:
	env GOOS=darwin go build -ldflags "-w" -o ./binary/macos/analyticsgen github.com/prongbang/analyticsgen && chmod +x ./binary/macos/analyticsgen

build_window:
	env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./binary/windows/analyticsgen.exe github.com/prongbang/analyticsgen