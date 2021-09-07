name = archive-provide

build:
	go build -o $(name) main.go

docker:
	docker build -t $(name) .

clean:
	rm -f $(name)
