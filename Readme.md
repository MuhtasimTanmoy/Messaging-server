    # Messaging Server

Simple messeging server and channel test.



# TODO
- Create channel
- Publish Suscribe
- Client create



# Later
- Docker
- Travis
- Makefile




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

- Healthcheck
    - curl localhost:8080/_healthcheck
- ENV Setup
    - export PATH="/Users/tanmoy/go/bin/:$PATH" 
- Improper import solve. VSCODE ctrl+shift+p. Configure language specific settings.
```
"[go]": {

        "editor.formatOnSave": false,
        "editor.codeActionsOnSave": {
            "source.organizeImports": false
        },
    },
"go.formatTool": "gofmt",
```