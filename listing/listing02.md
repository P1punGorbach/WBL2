Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
...
После того, как встретился return, но до фактического выхода из функции, выполняется все defer-функция.
В defer внутри test() идёт x++. Поскольку x — это именно именованный результат, его изменение в defer влияет на итоговое значение, которое вернётся из функции.
Внутри anotherTest() x - локальная переменная. Копируется в return, и уже после этого выполняется defer функция
```
