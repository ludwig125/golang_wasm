参考：https://egghead.io/lessons/go-access-js-functions-and-variables-from-a-go-webassembly-program

lib.js

```js
function add(a, b) {
	return a + b;
}

function hello() {
	return "Hello World2!";
}

var env = "PROD";

var config = {};
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

#### syscall/js が認識されない

VSCode で `syscall/js`が読み込めていなかった。

![image](https://user-images.githubusercontent.com/18366858/141660201-8666bd42-324f-4ec7-b59e-864c968074e6.png)

> could not import syscall/js

と出たので、

https://github.com/Microsoft/vscode-go/issues/1874#issuecomment-416014631
を参考に、以下を設定したら認識された。

```
    "go.toolsEnvVars": {
        "GOOS": "js",
    	"GOARCH": "wasm",
    },
```

#### build

```
[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $go mod tidy
[~/go/src/github.com/ludwig125/golang_wasm/helloWorld] $GOOS=js GOARCH=wasm go build -o main.wasm
```

main.wasm が作られる

index.html

```html
<html>
	<head>
		<meta charset="utf-8" />
		<script src="wasm_exec.js"></script>
		<script src="lib.js"></script>
		<script>
			const go = new Go();
			WebAssembly.instantiateStreaming(
				fetch("main.wasm"),
				go.importObject
			).then((result) => {
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

Chrome で F12 で Console を開いて

![image](https://user-images.githubusercontent.com/18366858/140662250-81728739-7454-4225-8d34-b019b9c18622.png)

もし表示されたページの内容が前と変わっていなかったら、前回のキャッシュが残っている可能性があるので、Ctrl+Shift+R でページをリロードされれば更新されるかも。

以下のように、JS の変数を上書きすることもできる

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

index.html に

```
<link rel="shortcut icon" href="名前" type="＜画像のパス＞">
```

のように設定すれば自由な画像をタブに出すことができる

これがないとエラーを出すのが嫌なので、

以下のように `<link rel="shortcut icon" href="#" />` を追加すると、

```html
<meta charset="utf-8" />
<link rel="shortcut icon" href="#" />
<script src="wasm_exec.js"></script>
<script src="lib.js"></script>
```

参考：

- https://stackoverflow.com/questions/31075893/im-getting-favicon-ico-error
- https://wa3.i-3-i.info/word14168.html

このようにエラーがなくなる

![image](https://user-images.githubusercontent.com/18366858/140662967-50977aca-b0af-4c00-b4fc-54dfa93996bd.png)
