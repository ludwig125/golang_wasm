https://github.com/golang/go/wiki/WebAssembly


```
package main

import "fmt"

func main() {
	fmt.Println("Hello, WebAssembly!")
}
```

build
```
GOOS=js GOARCH=wasm go build -o main.wasm
```

```
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

```

http://localhost:8080/

![image](https://user-images.githubusercontent.com/18366858/140580741-5bbc5f59-9162-475b-8ebb-650cc4cc36f7.png)

