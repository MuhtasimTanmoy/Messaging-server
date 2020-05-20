    # Messaging Server

Simple messeging server and channel test.



# TODO
- Create channel
- Publish Suscribe
- Client create



# Later
- Docker
- Travis




## Notes
- Module init 
    -  go mod init github.com/MuhtasimTanmoy/messaging_server
- Init go

```go
cat <<EOF > hello.go
package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
EOF
```

- 