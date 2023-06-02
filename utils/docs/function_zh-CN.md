# Function
function函数包控制函数执行流程，包含部分函数式编程。

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/function/function.go](https://github.com/duke-git/lancet/blob/main/function/function.go)
- [https://github.com/duke-git/lancet/blob/main/function/watcher.go](https://github.com/duke-git/lancet/blob/main/function/watcher.go)

<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/x-module/utils/function"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [After](#After)
- [Before](#Before)
- [CurryFn](#CurryFn)
- [Compose](#Compose)
- [Debounced](#Debounced)
- [Delay](#Delay)
- [Pipeline](#Pipeline)
- [Watcher](#Watcher)

<div STYLE="page-break-after: always;"></div>

## 文档



### <span id="After">After</span>
<p>创建一个函数，当他被调用n或更多次之后将马上触发fn</p>

<b>函数签名:</b>

```go
func After(n int, fn any) func(args ...any) []reflect.Value
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/function"
)

func main() {
	arr := []string{"a", "b"}
	f := function.After(len(arr), func(i int) int {
		fmt.Println("last print")
		return i
	})

	type cb func(args ...any) []reflect.Value
	print := func(i int, s string, fn cb) {
		fmt.Printf("arr[%d] is %s \n", i, s)
		fn(i)
	}

	fmt.Println("arr is", arr)
	for i := 0; i < len(arr); i++ {
		print(i, arr[i], f)
	}

    //output:
    // arr is [a b]
    // arr[0] is a 
    // arr[1] is b 
    // last print
}
```



### <span id="Before">Before</span>

<p>创建一个函数，调用次数不超过n次，之后再调用这个函数，将返回一次最后调用fn的结果</p>

<b>函数签名:</b>

```go
func Before(n int, fn any) func(args ...any) []reflect.Value
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/function"
    "github.com/x-module/utils/internal"
)

func main() {
	arr := []string{"a", "b", "c", "d", "e"}
	f := function.Before(3, func(i int) int {
		return i
	})

	var res []int64
	type cb func(args ...any) []reflect.Value
	appendStr := func(i int, s string, fn cb) {
		v := fn(i)
		res = append(res, v[0].Int())
	}

	for i := 0; i < len(arr); i++ {
		appendStr(i, arr[i], f)
	}

	fmt.Println(res) // 0, 1, 2, 2, 2
}
```



### <span id="CurryFn">CurryFn</span>

<p>创建柯里化函数</p>

<b>函数签名:</b>

```go
type CurryFn[T any] func(...T) T
func (cf CurryFn[T]) New(val T) func(...T) T
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/function"
)

func main() {
	add := func(a, b int) int {
		return a + b
	}

	var addCurry CurryFn[int] = func(values ...int) int {
		return add(values[0], values[1])
	}
	add1 := addCurry.New(1)

	result := add1(2)

	fmt.Println(result) //3
}
```



### <span id="Compose">Compose</span>

<p>从右至左组合函数列表fnList，返回组合后的函数</p>

<b>函数签名:</b>

```go
func Compose[T any](fnList ...func(...T) T) func(...T) T
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/function"
)

func main() {
	toUpper := func(strs ...string) string {
		return strings.ToUpper(strs[0])
	}
	toLower := func(strs ...string) string {
		return strings.ToLower(strs[0])
	}
	transform := Compose(toUpper, toLower)

	result := transform("aBCde")

	fmt.Println(result) //ABCDE
}
```



### <span id="Debounced">Debounced</span>

<p>创建一个debounced函数，该函数延迟调用fn直到自上次调用debounced函数后等待持续时间过去。</p>

<b>函数签名:</b>

```go
func Debounced(fn func(), duration time.Duration) func()
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/function"
)

func main() {
	count := 0
	add := func() {
		count++
	}

	debouncedAdd := function.Debounced(add, 50*time.Microsecond)
	function.debouncedAdd()
	function.debouncedAdd()
	function.debouncedAdd()
	function.debouncedAdd()

	time.Sleep(100 * time.Millisecond)
	fmt.Println(count) //1

	function.debouncedAdd()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(count) //2
}
```



### <span id="Delay">Delay</span>

<p>延迟delay时间后调用函数</p>

<b>函数签名:</b>

```go
func Delay(delay time.Duration, fn any, args ...any)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/function"
)

func main() {
	var print = func(s string) {
		fmt.Println(count) //test delay
	}
	function.Delay(2*time.Second, print, "test delay")
}
```



### <span id="Schedule">Schedule</span>

<p>每次持续时间调用函数，直到关闭返回的 bool chan</p>

<b>函数签名:</b>

```go
func Schedule(d time.Duration, fn any, args ...any) chan bool
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/function"
)

func main() {
    var res []string
	appendStr := func(s string) {
		res = append(res, s)
	}

	stop := function.Schedule(1*time.Second, appendStr, "*")
	time.Sleep(5 * time.Second)
	close(stop)

	fmt.Println(res) //[* * * * *]
}
```



### <span id="Pipeline">Pipeline</span>

<p>执行函数pipeline.</p>

<b>函数签名:</b>

```go
func Pipeline[T any](funcs ...func(T) T) func(T) T
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/function"
)

func main() {
    addOne := func(x int) int {
		return x + 1
	}
	double := func(x int) int {
		return 2 * x
	}
	square := func(x int) int {
		return x * x
	}

	f := Pipeline(addOne, double, square)

	fmt.Println(fn(2)) //36
}
```



### <span id="Watcher">Watcher</span>

<p>Watcher用于记录代码执行时间。可以启动/停止/重置手表定时器。获取函数执行的时间。</p>

<b>函数签名:</b>

```go
type Watcher struct {
	startTime int64
	stopTime  int64
	excuting  bool
}
func NewWatcher() *Watcher
func (w *Watcher) Start() //start the watcher
func (w *Watcher) Stop() //stop the watcher
func (w *Watcher) Reset() //reset the watcher
func (w *Watcher) GetElapsedTime() time.Duration //get the elapsed time of function execution
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/function"
)

func main() {
    w := function.NewWatcher()

	w.Start()

	longRunningTask()

	fmt.Println(w.excuting) //true

	w.Stop()

	eapsedTime := w.GetElapsedTime().Milliseconds()
	
	fmt.Println(eapsedTime)

	w.Reset()

}

func longRunningTask() {
	var slice []int64
	for i := 0; i < 10000000; i++ {
		slice = append(slice, int64(i))
	}
}

```



