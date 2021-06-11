
clean:
	rm out/gogen

build:
	go build -o out/gogen

install: build
	go install