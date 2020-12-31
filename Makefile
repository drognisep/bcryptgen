test:
	go test ./...

local: test
	go run main.go

check-fyne-cross:
	fyne-cross version || go get github.com/fyne-io/fyne-cross

package-linux: check-fyne-cross
	fyne-cross linux --pull

package-windows: check-fyne-cross
	fyne-cross windows --pull

release: test package-linux package-windows
