# Retry
Package retry is for executing a function repeatedly until it was successful or canceled by the context.

<div STYLE="page-break-after: always;"></div>

## Source:

- [https://github.com/duke-git/lancet/blob/main/retry/retry.go](https://github.com/duke-git/lancet/blob/main/retry/retry.go)


<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/x-module/utils/retry"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
- [Context](#Context)
- [Retry](#Retry)
- [RetryFunc](#RetryFunc)
- [RetryDuration](#RetryDuration)
- [RetryTimes](#RetryTimes)

<div STYLE="page-break-after: always;"></div>

## Documentation


### <span id="Context">Context</span>
<p>Set retry context config, can cancel the retry with context.</p>

<b>Signature:</b>

```go
func Context(ctx context.Context)
```
<b>Example:</b>

```go
import (
	"context"
	"errors"
	"fmt"
	"github.com/x-module/utils/retry"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	var number int
	increaseNumber := func() error {
		number++
		if number > 3 {
			cancel()
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber,
		retry.RetryDuration(time.Microsecond*50),
		retry.Context(ctx),
	)

	if err != nil {
		fmt.Println(err) //retry is cancelled
	}
}
```




### <span id="RetryFunc">RetryFunc</span>
<p>Function that retry executes.</p>

<b>Signature:</b>

```go
type RetryFunc func() error
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/x-module/utils/retry"
)

func main() {
    var number int

	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber, retry.RetryDuration(time.Microsecond*50))
    if err != nil {
		log.Fatal(err)
	}

    fmt.Println(number) //3
}
```



### <span id="RetryTimes">RetryTimes</span>
<p>Set times of retry. Default times is 5.</p>

<b>Signature:</b>

```go
func RetryTimes(n uint)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/x-module/utils/retry"
)

func main() {
    var number int

	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber, retry.RetryTimes(2))
    if err != nil {
		log.Fatal(err) //2022/02/01 18:42:25 function main.main.func1 run failed after 2 times retry exit status 1
	}
}
```



### <span id="RetryDuration">RetryDuration</span>
<p>Set duration of retries. Default duration is 3 second.</p>

<b>Signature:</b>

```go
func RetryDuration(d time.Duration)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/x-module/utils/retry"
)

func main() {
    var number int
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber, retry.RetryDuration(time.Microsecond*50))
    if err != nil {
		log.Fatal(err)
	}

    fmt.Println(number) //3
}
```


### <span id="Retry">Retry</span>
<p>Executes the retryFunc repeatedly until it was successful or canceled by the context.</p>

<b>Signature:</b>

```go
func Retry(retryFunc RetryFunc, opts ...Option) error
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/x-module/utils/retry"
)

func main() {
    var number int
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber, retry.RetryDuration(time.Microsecond*50))
    if err != nil {
		log.Fatal(err)
	}

    fmt.Println(number) //3
}
```
