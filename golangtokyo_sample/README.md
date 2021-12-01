https://golangtokyo.github.io/codelab/go-webassembly/?index=codelab#4

js.Value と js.Type
次に、JavaScript のオブジェクトを扱う方法について学びましょう。Go から JavaScript のオブジェクトを扱うためには、Go1.11 から標準パッケージに入った syscall/js パッケージを用います。

Go の上では JavaScript のオブジェクトは、js.Value 型という値で表現されます。Go の値を js.Value 型に変換するためには、js.ValueOf という関数を用います。js.ValueOf 関数は引数に Go の任意の値を interface{}型（どんな型の値でも入る）として受け取り、js.Value 型の値を戻り値として返す関数です。

Go の値と JavaScript の値は表 1 のような対応になっています。

表 1：Go の値と JavaScript の値の対応
|Go|JavaScript|
|:----|:----|
|js.Value|JavaScript の任意の値|
|js.TypedArray|typed array|
|js.Callback|関数|
|nil|null|
|bool|Boolean|
|整数と浮動小数点数|Number|
|string|String|
|[]interface{}|新しい配列|
|map[string]interface{}|新しいオブジェクト|

js.Value 型はすべての JavaScript 上の値を 1 つの型として表現しているため、各メソッドが予期しない値に対して呼ばれた場合に、panic を起こしてしまいます。例えば、Int メソッドは js.Value 型の値を数値として扱い、その値を Go の int 型として取得することができます。しかし、js.Value 型は関数や文字列の値も扱うことができるため、数値ではない値に呼ばれた場合には panic が発生します。

そのため、js.Type 型の値を用いることで js.Value 型の値が具体的にはどのような型なのかをハンドリングすることができ、panic を避けることができます。例えば、Int メソッドの場合は、次のように Type メソッドが js.TypeNumber を返したときのみ呼ぶほうが良いでしょう。

```
func printNumber(v js.Value) {
        if v.Type() == js.TypeNumber {
                fmt.Printf("%d\n", v.Int())
        }
}
```
