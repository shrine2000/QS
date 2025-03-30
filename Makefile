.PHONY: build install clean

build:
	go build -o qs cmd/qs/main.go

install: build
	mkdir -p ~/.local/bin
	cp qs ~/.local/bin/
	chmod +x ~/.local/bin/qs

clean:
	rm -f qs

uninstall:
	rm -f ~/.local/bin/qs 