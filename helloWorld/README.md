main.go

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

```

```
[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $go mod init
go: creating new go.mod: module github.com/ludwig125/golang_wasm/helloWorld
go: to add module requirements and sums:
        go mod tidy
[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $go mod tidy
[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $GOOS=js GOARCH=wasm go build -o main.wasm
```

main.wasmが作られる
```
[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $ls
README.md  go.mod  main.go  main.wasm*
```


WasmはJavaScriptから呼び出す必要があるので、そのためにindex.htmlを用意する

index.html
```html
<html>
    <head>
       <meta charset="utf-8">
       <script src="wasm_exec.js"></script>
       <script>
           const go = new Go();
           WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
               go.run(result.instance);
           });
       </script>
   </head>
    <body></body>
</html>
```

上の `wasm_exec.js` はGoに最初から含まれているものなので、それをカレントディレクトリにコピーする

```
[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $go env GOROOT
/usr/local/go
[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $ls
README.md  go.mod  index.html  main.go  main.wasm*  wasm_exec.js
```

サーバを立てるのに `goexec` を使う

```
$ go get -u github.com/shurcooL/goexec
```

```
$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```

http://localhost:8080/

ChromeでF12でConsoleを開いて

![image](https://user-images.githubusercontent.com/18366858/140625693-1439a489-129c-4ad8-95ec-07c9c6b6d8ef.png)

`Hello World!` が表示された
