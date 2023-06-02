# Maputil
maputil包包括一些操作map的函数。

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/maputil/map.go](https://github.com/duke-git/lancet/blob/main/maputil/map.go)


<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/x-module/utils/maputil"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录:
- [ForEach](#ForEach)
- [Filter](#Filter)
- [Intersect](#Intersect)
- [Keys](#Keys)
- [Merge](#Merge)
- [Minus](#Minus)
- [Values](#Values)
- [IsDisjoint](#IsDisjoint)

<div STYLE="page-break-after: always;"></div>

## API文档:



### <span id="ForEach">ForEach</span>
<p>对map中的每对key和value执行iteratee函数</p>

<b>函数签名:</b>

```go
func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V))
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/maputil"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	var sum int

	maputil.ForEach(m, func(_ string, value int) {
		sum += value
	})
	fmt.Println(sum) // 10
}
```




### <span id="Filter">Filter</span>
<p>迭代map中的每对key和value, 返回符合predicate函数的key, value</p>

<b>函数签名:</b>

```go
func Filter[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/maputil"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	isEven := func(_ string, value int) bool {
		return value%2 == 0
	}

	maputil.Filter(m, func(_ string, value int) {
		sum += value
	})
	res := maputil.Filter(m, isEven)
	fmt.Println(res) // map[string]int{"b": 2, "d": 4,}
}
```




### <span id="Intersect">Intersect</span>
<p>多个map的交集操作</p>

<b>函数签名:</b>

```go
func Intersect[K comparable, V any](maps ...map[K]V) map[K]V
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/maputil"
)

func main() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 6,
		"d": 7,
	}

	m3 := map[string]int{
		"a": 1,
		"b": 9,
		"e": 9,
	}

	fmt.Println(maputil.Intersect(m1)) // map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println(maputil.Intersect(m1, m2)) // map[string]int{"a": 1, "b": 2}

	fmt.Println(maputil.Intersect(m1, m2, m3)) // map[string]int{"a": 1}
}
```




### <span id="Keys">Keys</span>
<p>返回map中所有key的切片</p>

<b>函数签名:</b>

```go
func Keys[K comparable, V any](m map[K]V) []K
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/maputil"
)

func main() {
	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	keys := maputil.Keys(m)
	sort.Ints(keys)
	fmt.Println(keys) // []int{1, 2, 3, 4, 5}
}
```




### <span id="Merge">Merge</span>
<p>合并多个maps, 相同的key会被后来的key覆盖</p>

<b>函数签名:</b>

```go
func Merge[K comparable, V any](maps ...map[K]V) map[K]V
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/maputil"
)

func main() {
	m1 := map[int]string{
		1: "a",
		2: "b",
	}
	m2 := map[int]string{
		1: "1",
		3: "2",
	}
	fmt.Println(maputil.Merge(m1, m2)) // map[int]string{1:"1", 2:"b", 3:"2",}
}
```



### <span id="Minus">Minus</span>
<p>返回一个map，其中的key存在于mapA，不存在于mapB.</p>

<b>函数签名:</b>

```go
func Minus[K comparable, V any](mapA, mapB map[K]V) map[K]V
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/maputil"
)

func main() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 11,
		"b": 22,
		"d": 33,
	}

	fmt.Println(maputil.Minus(m1, m2)) //map[string]int{"c": 3}
}
```



### <span id="Values">Values</span>
<p>返回map中所有value的切片</p>

<b>函数签名:</b>

```go
func Values[K comparable, V any](m map[K]V) []V
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/maputil"
)

func main() {
	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	values := maputil.Values(m)
	sort.Strings(values)

	fmt.Println(values) // []string{"a", "a", "b", "c", "d"}
}
```


### <span id="IsDisjoint">IsDisjoint</span>
<p>验证两个map是否具有不同的key</p>

<b>函数签名:</b>

```go
func IsDisjoint[K comparable, V any](mapA, mapB map[K]V) bool
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/maputil"
)

func main() {
	m1 := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	m2 := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	m3 := map[int]string{
		6: "a",
	}

	ok := maputil.IsDisjoint(m2, m1)
	fmt.Println(ok) // false

	ok = maputil.IsDisjoint(m2, m3)
	fmt.Println(ok) // true
}
```
