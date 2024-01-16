build:
	rm -f rabbitmq 
	go build -v -o rabbitmq ./...

run: build  
	./rabbitmq

clean:
	rm -f rabbitmq