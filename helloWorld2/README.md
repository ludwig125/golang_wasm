lib.js
```js
function add(a, b){
    return a + b;
}

function hello() {
    return "Hello World2!";
}

var env = "PROD";

var config = {}

```


参考：

- https://pkg.go.dev/syscall/js

main.go
```go
package main

import "syscall/js"

func main() {
	num := js.Global().Call("add", 3, 4)
	println(num.Float())
	println(num.Int())

	s := js.Global().Call("hello").String()
	println(s)

	env := js.Global().Get("env").String()
	println(env)
}

```


```
[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $go mod tidy
[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $GOOS=js GOARCH=wasm go build -o main.wasm
```

main.wasmが作られる


index.html
```html
<html>
    <head>
       <meta charset="utf-8">
       <script src="wasm_exec.js"></script>
       <script src="lib.js"></script>
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

 `goexec`

```
$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```

http://localhost:8080/

ChromeでF12でConsoleを開いて

![image](https://user-images.githubusercontent.com/18366858/140662250-81728739-7454-4225-8d34-b019b9c18622.png)


以下のように、JSの変数を上書きすることもできる

- この例では、`PROD` を `DEV` に、`config` に `key` を設定している

```go
package main

import "syscall/js"

func main() {
	num := js.Global().Call("add", 3, 4)
	println(num.Float())
	println(num.Int())

	s := js.Global().Call("hello").String()
	println(s)

	env := js.Global().Get("env").String()
	println(env)

	js.Global().Set("env", "DEV")
	env2 := js.Global().Get("env").String()
	println(env2)

	js.Global().Get("config").Set("key", "12345")
	conf := js.Global().Get("config").Get("key").String()
	println(conf)

}
```

再度ビルドしてサーバを立ち上げる

```
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```



![image](https://user-images.githubusercontent.com/18366858/140662751-ab4b2fcc-1b25-43fd-b2cf-8d70f26afd1a.png)

`GET http://localhost:8080/favicon.ico 404 (Not Found)` と出ているのが気になる。。

`favicon.ico` というのは、ブラウザのタブのところのアイコンのこと。

![image](https://user-images.githubusercontent.com/18366858/140662872-b408b2e3-e87f-4fbc-a065-ea750915595b.png)

index.htmlに

```
<link rel="shortcut icon" href="名前" type="＜画像のパス＞">
```

のように設定すれば自由な画像をタブに出すことができる

これがないとエラーを出すのが嫌なので、

以下のように `<link rel="shortcut icon" href="#" />` を追加すると、
```html
       <meta charset="utf-8">
       <link rel="shortcut icon" href="#" />
       <script src="wasm_exec.js"></script>
       <script src="lib.js"></script>
```

参考：

- https://stackoverflow.com/questions/31075893/im-getting-favicon-ico-error
- https://wa3.i-3-i.info/word14168.html

このようにエラーがなくなる

![image](https://user-images.githubusercontent.com/18366858/140662967-50977aca-b0af-4c00-b4fc-54dfa93996bd.png)
