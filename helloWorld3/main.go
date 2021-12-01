package main

import (
	"strings"
	"syscall/js"
)

func say(this js.Value, args []js.Value) interface{} {
	ss := []string{}
	for _, jss := range args {
		if s := jsString(jss); s != "" {
			ss = append(ss, s)
		}
	}
	return js.ValueOf("Hello " + strings.Join(ss, ", "))
}

func jsString(j js.Value) string {
	if j.IsUndefined() || j.IsNull() {
		return ""
	}
	return j.String()
}

func main() {
	ch := make(chan struct{})
	js.Global().Set("say", js.FuncOf(say))
	<-ch // Code must not finish
}

// // documentオブジェクト取得用
// var document = js.Global().Get("document")

// // windowオブジェクトを取得
// var window = js.Global()

// // bodyのDOM取得
// var body = document.Get("body")

// func main() {
// 	// goからbuttonのDOMを作成する
// 	cLogBtn := createElement("button")
// 	// cLogBtnボタンのテキストを設定
// 	cLogBtn.Set("textContent", "console log!!")
// 	// buttonをbodyへ追加
// 	body.Call("appendChild", cLogBtn)
// 	// cLogBtnにclickのEventLisnerを設定
// 	cLogBtn.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
// 		fmt.Println("Hello Webassembly!")
// 		return nil
// 	}))

// 	// buttonDOMを作成する
// 	textChangeBtn := createElement("button")
// 	// textChangeBtnのテキストを設定
// 	textChangeBtn.Set("textContent", "text change!!")
// 	// buttonをbodyへ追加
// 	body.Call("appendChild", textChangeBtn)
// 	// textChangeにclick時のEventLisnerを設定
// 	textChangeBtn.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
// 		message := getElementByID("message")
// 		message.Set("innerHTML", "Hello, WebAssembry!!")
// 		return nil
// 	}))

// 	// buttonのDOMを生成
// 	alertBtn := createElement("button")
// 	// alertBtnのテキストを設定
// 	alertBtn.Set("textContent", "alert!!")
// 	// buttonをbodyへ追加
// 	body.Call("appendChild", alertBtn)
// 	// alertBtnにclick時のEventLisnerを設定
// 	alertBtn.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
// 		window.Call("alert", "Hello!!")
// 		return nil
// 	}))

// 	// 既にhtmlに書かれているボタンのidを取得して、goからclick時のEventListenerを設定する
// 	getElementByID("init-text").Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
// 		message := getElementByID("message")
// 		message.Set("innerHTML", "Hello, World")
// 		return nil
// 	}))

// 	// textエリアの入力値を取得
// 	getElementByID("in").Call("addEventListener", "keyup", js.FuncOf(func(js.Value, []js.Value) interface{} {
// 		getElementByID("out").Set("value", getElementByID("in").Get("value"))
// 		return nil
// 	}))

// 	// プログラムが終了しないように待機
// 	select {}
// }

// // 使用頻度が高そうなので、対象のDOMのIDを取得する関数を用意
// func getElementByID(targetID string) js.Value {
// 	return document.Call("getElementById", targetID)
// }

// // 使用頻度が高そうなので、対象DOMを作成する関数を用意
// func createElement(elementName string) js.Value {
// 	return document.Call("createElement", elementName)
// }

// func main() {
// 	ch := make(chan struct{})
// 	js.Global().Get("document").Call("getElementById", "hello").Set("innerHTML", "Hello, World!!!")
// 	<-ch // Code must not finish
// }

// func main() {
// 	ch := make(chan struct{})
// 	js.Global().Set("say", js.FuncOf(say))
// 	<-ch // Code must not finish
// }

// func say(this js.Value, args []js.Value) interface{} {
// 	ss := []string{}
// 	for _, jss := range args {
// 		if s := jsString(jss); s != "" {
// 			ss = append(ss, s)
// 		}
// 	}
// 	return js.ValueOf("Hello, " + strings.Join(ss, ", "))
// }

// func jsString(j js.Value) string {
// 	if j.IsUndefined() || j.IsNull() {
// 		return ""
// 	}
// 	return j.String()
// }

// func main() {
// 	c := make(chan struct{}, 0)
// 	registerCallbacks()
// 	<-c
// }

// func registerCallbacks() {
// 	// js.Global().Set("print", js.NewCallback(print))
// 	js.Global().Set("print", js.FuncOf(print))
// 	// js.Global().Set("manipulateDom", js.FuncOf(manipulateDom))
// }

// func print(this js.Value, i []js.Value) interface{} {
// 	// fmt.Println(i)
// 	fmt.Println(this)
// 	// fmt.Println(i[0].String())
// 	return this
// }

// // func manipulateDom(this js.Value, i []js.Value) interface{} {
// // 	js.Global().Get("document").Call("getElementById", "result").Set("textContent", "WebAssembly!")
// // 	return i
// // }
