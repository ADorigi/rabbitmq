build:
	rm -f rabbitmq 
	go build -v -o rabbitmq ./...

run: build  
	./rabbitmq

test:
	go clean -testcache && go test -v ./...

clean:
	rm -f rabbitmq