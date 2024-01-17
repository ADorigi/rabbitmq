# rabbitmq
rabbitmq wrapper library in go

# How to Use

## Module Init and dependency addition
Assuming you have go installed(I love GVM), the commands to init module and add this package's dependency are as follows

```
go mod init <<module name>>
go get -u github.com/adorigi/rabbitmq
```

## Required Environment Variables: 
Make these two environment variables accessible to your go executable.

```
RABBIT_ENDPOINT="amqp://<<username>>:<<password>>@<<hostname>>:"
RABBIT_AMQP_PROTOCOL="<<port>>"
```

## Sample program

```
package main

import (
	"context"
	"time"

    "github.com/adorigi/rabbitmq"
)

func main() {

	rabbit := rabbitmq.NewRabbit()
	rabbit.Configure()
	rabbit.ConnectSocket()
	rabbit.ConnectChannel()
	rabbit.DeclareQueue(<<queuename>>)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	rabbit.PublishCTXByte(ctx, []byte(body))

	rabbit.Close()

}


```


<!-- 
TODO:
done- Remove dependecy on godotenv
done- add argument to declarequeue for accepting queue name
    - improve tests 

My aim for this repo:
- error code masking - the bad(but not ugly) side of go
- useful abstraction
- will be updated - acc. to how my needs change (but I will tag releases)
 -->