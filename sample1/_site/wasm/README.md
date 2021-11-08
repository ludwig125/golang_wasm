```
go: cannot find main module, but found .git/config in /home/ludwig125/go/src/github.com/ludwig125/asset-simulator
        to create a module there, run:
        go mod init
[~/go/src/github.com/ludwig125/asset-simulator] $go mod init
go: creating new go.mod: module github.com/ludwig125/asset-simulator
go: to add module requirements and sums:
        go mod tidy
[~/go/src/github.com/ludwig125/asset-simulator] $go mod tidy
[~/go/src/github.com/ludwig125/asset-simulator] $
```

```
[~/go/src/github.com/ludwig125/asset-simulator] $cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
GOOS=js GOARCH=wasm go build -o main.wasm
```

```
[~/go/src/github.com/ludwig125/asset-simulator] $ls /usr/local/go/misc/wasm/wasm_exec.js
/usr/local/go/misc/wasm/wasm_exec.js
[~/go/src/github.com/ludwig125/asset-simulator] $

[~/go/src/github.com/ludwig125/asset-simulator] $ls main.wasm
main.wasm*
```

サーバ立ち上げる
```
[~/go/src/github.com/ludwig125/asset-simulator] $bundle exec jekyll serve
Configuration file: /home/ludwig125/go/src/github.com/ludwig125/asset-simulator/_config.yml
            Source: /home/ludwig125/go/src/github.com/ludwig125/asset-simulator
       Destination: /home/ludwig125/go/src/github.com/ludwig125/asset-simulator/_site
 Incremental build: disabled. Enable with --incremental
      Generating...
       Jekyll Feed: Generating feed for posts
                    done in 0.271 seconds.
/home/ludwig125/gems/gems/pathutil-0.16.2/lib/pathutil.rb:502: warning: Using the last argument as keyword parameters is deprecated
                    Auto-regeneration may not work on some Windows versions.
                    Please see: https://github.com/Microsoft/BashOnWindows/issues/216
                    If it does not work, please upgrade Bash on Windows or run Jekyll with --no-watch.
 Auto-regeneration: enabled for '/home/ludwig125/go/src/github.com/ludwig125/asset-simulator'
    Server address: http://127.0.0.1:4000/
  Server running... press ctrl-c to stop.
```

```
http://localhost:4000/ をブラウザで見ると

```

![image](https://user-images.githubusercontent.com/18366858/138359677-36234fbb-8939-467d-8fbe-d3f05f93a3e4.png)
![image](https://user-images.githubusercontent.com/18366858/138361172-7fa4dfd4-139c-4f06-8afe-c37842cbc80b.png)

![image](https://user-images.githubusercontent.com/18366858/139741327-aa09c991-6663-4ab9-b984-e60f33faea64.png)
↑ F12押すと出てくる

ouch TypeError: Failed to execute 'compile' on 'WebAssembly': Incorrect response MIME type. Expected 'application/wasm'.
:4000/favicon.ico:1 Failed to load resource: the server responded with a status of 404 (Not Found)


この `GET http://localhost:4000/favicon.ico 404 (Not Found)` については、

https://github.com/golang/go/wiki/WebAssembly
を参考に、
https://github.com/golang/go/blob/b2fcfc1a50fbd46556f7075f7f1fbf600b5c9e5d/misc/wasm/wasm_exec.html
のファイルをwasm_exec.jsの隣に置けばなくなった

nodejs install
```
sudo apt install nodejs
```

Node.jsで使用するモジュールとパッケージをインストールするため
```
sudo apt install npm
```


```
npm init fastify
```


package.jsonを以下のようにする
```
{
    "dependencies": {
        "fastify": "^3.6.0",
        "fastify-static": "^3.2.1"
    }
}

```


```
npm install
```

サーバ起動
```
$node index.js
```

http://127.0.0.1:8080/

![image](https://user-images.githubusercontent.com/18366858/139743928-2dc0d348-a046-46da-9042-1768abfbc37f.png)
