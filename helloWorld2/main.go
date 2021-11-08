package main

import "syscall/js"

// ref: https://egghead.io/lessons/go-access-js-functions-and-variables-from-a-go-webassembly-program

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
