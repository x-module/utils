# Linklist
Linklist a linked list, whose node has a value and a pointer points to next node of the link.

<div STYLE="page-break-after: always;"></div>

## Source

- [https://github.com/duke-git/lancet/blob/main/datastructure/link/singlylink.go](https://github.com/duke-git/lancet/blob/main/datastructure/link/singlylink.go)
- [https://github.com/duke-git/lancet/blob/main/datastructure/link/doublylink.go](https://github.com/duke-git/lancet/blob/main/datastructure/link/doublylink.go)


<div STYLE="page-break-after: always;"></div>

## Usage
```go
import (
    link "github.com/x-module/utils/datastructure/link"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

### 1. SinglyLink

- [NewSinglyLink](#NewSinglyLink)
- [Values](#SinglyLink_Values)
- [InsertAt](#SinglyLink_InsertAt)
- [InsertAtHead](#SinglyLink_InsertAtHead)
- [InsertAtTail](#SinglyLink_InsertAtTail)
- [DeleteAt](#SinglyLink_DeleteAt)
- [DeleteAtHead](#SinglyLink_DeleteAtHead)
- [DeleteAtTail](#SinglyLink_DeleteAtTail)
- [DeleteValue](#SinglyLink_DeleteValue)
- [Reverse](#SinglyLink_Reverse)
- [GetMiddleNode](#SinglyLink_GetMiddleNode)
- [Size](#SinglyLink_Size)
- [IsEmpty](#SinglyLink_IsEmpty)
- [Clear](#SinglyLink_Clear)
- [Print](#SinglyLink_Print)

### 2. DoublyLink

- [NewDoublyLink](#NewDoublyLink)
- [Values](#DoublyLink_Values)
- [InsertAt](#DoublyLink_InsertAt)
- [InsertAtHead](#DoublyLink_InsertAtHead)
- [InsertAtTail](#DoublyLink_InsertAtTail)
- [DeleteAt](#DoublyLink_DeleteAt)
- [DeleteAtHead](#DoublyLink_DeleteAtHead)
- [DeleteAtTail](#DoublyLink_DeleteAtTail)
- [Reverse](#DoublyLink_Reverse)
- [GetMiddleNode](#DoublyLink_GetMiddleNode)
- [Size](#DoublyLink_Size)
- [IsEmpty](#DoublyLink_IsEmpty)
- [Clear](#DoublyLink_Clear)
- [Print](#DoublyLink_Print)


<div STYLE="page-break-after: always;"></div>

## Documentation

### 1. SinglyLink
SinglyLink a linked list, whose node has a value and a pointer points to next node of the link.

### <span id="NewSinglyLink">NewSinglyLink</span>
<p>Return a singly link(SinglyLink) instance</p>

<b>Signature:</b>

```go
type LinkNode[T any] struct {
	Value T
	Next  *LinkNode[T]
}
type SinglyLink[T any] struct {
	Head   *datastructure.LinkNode[T]
	length int
}
func NewSinglyLink[T any]() *SinglyLink[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()
    fmt.Println(lk)
}
```



### <span id="SinglyLink_Values">Values</span>
<p>Return a slice of all node value in singly linklist</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) Values() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```




### <span id="SinglyLink_InsertAt">InsertAt</span>
<p>Insert value into singly linklist at index, param `index` should between [0, len(SinglyLink)], if index do not meet the conditions, do nothing</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) InsertAt(index int, value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAt(1, 1) //do nothing
    
    lk.InsertAt(0, 1)
    lk.InsertAt(1, 2)
    lk.InsertAt(2, 3)
    lk.InsertAt(2, 4)

    fmt.Println(lk.Values()) //[]int{1, 2, 4, 3}
}
```




### <span id="SinglyLink_InsertAtHead">InsertAtHead</span>
<p>Insert value into singly linklist at head(first) index</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) InsertAtHead(value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtHead(1)
    lk.InsertAtHead(2)
    lk.InsertAtHead(3)

    fmt.Println(lk.Values()) //[]int{3, 2, 1}
}
```




### <span id="SinglyLink_InsertAtTail">InsertAtTail</span>
<p>Insert value into singly linklist at tail(last) index</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) InsertAtTail(value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```



### <span id="SinglyLink_DeleteAt">DeleteAt</span>
<p>Delete value at specific index, param `index` should be [0, len(SinglyLink)-1]</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) DeleteAt(index int)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    lk.InsertAtTail(4)

    lk.DeleteAt(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```



### <span id="SinglyLink_DeleteAtHead">DeleteAtHead</span>
<p>Delete value in singly linklist at first index</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) DeleteAtHead()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    lk.InsertAtTail(4)

    lk.DeleteAtHead()
    
    fmt.Println(lk.Values()) //[]int{2, 3, 4}
}
```




### <span id="SinglyLink_DeleteAtTail">DeleteAtTail</span>
<p>Delete value in singly linklist at last index</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) DeleteAtTail()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.DeleteAtTail()
    
    fmt.Println(lk.Values()) //[]int{1, 2}
}
```



### <span id="SinglyLink_DeleteValue">DeleteValue</span>
<p>Delete all `value` in singly linklist</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) DeleteValue(value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.DeleteValue(2)
    fmt.Println(lk.Values()) //[]int{1, 3}
}
```




### <span id="SinglyLink_Reverse">Reverse</span>
<p>Reverse all nodes order in linkist</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) Reverse()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.Reverse()
    fmt.Println(lk.Values()) //[]int{3, 2, 1}
}
```



### <span id="SinglyLink_GetMiddleNode">GetMiddleNode</span>
<p>Get the node at middle index of linkist</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) GetMiddleNode() *datastructure.LinkNode[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    midNode := lk.GetMiddleNode()
    fmt.Println(midNode.Value) //2
}
```



### <span id="SinglyLink_Size">Size</span>
<p>Get the number of nodes in linklist</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) Size() int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Size()) //3
}
```



### <span id="SinglyLink_IsEmpty">IsEmpty</span>
<p>Checks if linklist is empty or not</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) IsEmpty() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()
    fmt.Println(lk.IsEmpty()) //true

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.IsEmpty()) //false
}
```



### <span id="SinglyLink_Clear">Clear</span>
<p>Clear all nodes in the linklist, make it empty</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) Clear()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.Clear()

    fmt.Println(lk.Values()) //
}
```



### <span id="SinglyLink_Print">Print</span>
<p>Print all nodes info of linklist</p>

<b>Signature:</b>

```go
func (link *SinglyLink[T]) Clear()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    
    lk.Print() //[ &{Value:1 Pre:<nil> Next:0xc0000a4048}, &{Value:2 Pre:<nil> Next:0xc0000a4060}, &{Value:3 Pre:<nil> Next:<nil>} ]
}
```



### 2. DoublyLink
DoublyLink is a linked list, whose node has a value, a next pointer points to next node and pre pointer points to previous node of the link.

### <span id="NewDoublyLink">NewDoublyLink</span>
<p>Return a doubly link instance</p>

<b>Signature:</b>

```go
type LinkNode[T any] struct {
	Value T
    Pre   *LinkNode[T]
	Next  *LinkNode[T]
}
type DoublyLink[T any] struct {
	Head   *datastructure.LinkNode[T]
	length int
}
func NewDoublyLink[T any]() *DoublyLink[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()
    fmt.Println(lk)
}
```



### <span id="DoublyLink_Values">Values</span>
<p>Return a slice of all node value in doubly linklist</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) Values() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```




### <span id="DoublyLink_InsertAt">InsertAt</span>
<p>Insert value into doubly linklist at index, param `index` should between [0, len(DoublyLink)], if index do not meet the conditions, do nothing</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) InsertAt(index int, value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAt(1, 1) //do nothing

    lk.InsertAt(0, 1)
    lk.InsertAt(1, 2)
    lk.InsertAt(2, 3)
    lk.InsertAt(2, 4)

    fmt.Println(lk.Values()) //[]int{1, 2, 4, 3}
}
```




### <span id="DoublyLink_InsertAtHead">InsertAtHead</span>
<p>Insert value into doubly linklist at head(first) index</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) InsertAtHead(value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtHead(1)
    lk.InsertAtHead(2)
    lk.InsertAtHead(3)

    fmt.Println(lk.Values()) //[]int{3, 2, 1}
}
```




### <span id="DoublyLink_InsertAtTail">InsertAtTail</span>
<p>Insert value into doubly linklist at tail(last) index</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) InsertAtTail(value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```



### <span id="DoublyLink_DeleteAt">DeleteAt</span>
<p>Delete value at specific index, param `index` should be [0, len(DoublyLink)-1]</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) DeleteAt(index int)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    lk.InsertAtTail(4)

    lk.DeleteAt(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```



### <span id="DoublyLink_DeleteAtHead">DeleteAtHead</span>
<p>Delete value in doubly linklist at first index</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) DeleteAtHead()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    lk.InsertAtTail(4)

    lk.DeleteAtHead()
    
    fmt.Println(lk.Values()) //[]int{2, 3, 4}
}
```




### <span id="DoublyLink_DeleteAtTail">DeleteAtTail</span>
<p>Delete value in doubly linklist at last index</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) DeleteAtTail() error
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    err := lk.DeleteAtTail()
    
    fmt.Println(err) //nil
    fmt.Println(lk.Values()) //[]int{1, 2}
}
```




### <span id="DoublyLink_Reverse">Reverse</span>
<p>Reverse all nodes order in linkist</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) Reverse()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.Reverse()
    fmt.Println(lk.Values()) //[]int{3, 2, 1}
}
```



### <span id="DoublyLink_GetMiddleNode">GetMiddleNode</span>
<p>Get the node at middle index of linkist</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) GetMiddleNode() *datastructure.LinkNode[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    midNode := lk.GetMiddleNode()
    fmt.Println(midNode.Value) //2
}
```



### <span id="DoublyLink_Size">Size</span>
<p>Get the number of nodes in linklist</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) Size() int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Size()) //3
}
```



### <span id="DoublyLink_IsEmpty">IsEmpty</span>
<p>Checks if linklist is empty or not</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) IsEmpty() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()
    fmt.Println(lk.IsEmpty()) //true

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.IsEmpty()) //false
}
```



### <span id="DoublyLink_Clear">Clear</span>
<p>Clear all nodes in the linklist, make it empty</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) Clear()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.Clear()

    fmt.Println(lk.Values()) //
}
```



### <span id="DoublyLink_Print">Print</span>
<p>Print all nodes info of linklist</p>

<b>Signature:</b>

```go
func (link *DoublyLink[T]) Clear()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    link "github.com/x-module/utils/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    
    lk.Print() //
}
```
