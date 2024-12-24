Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false
...
Функция Foo() возвращает (*os.PathError)(nil) в виде интерфейса error.
В Go интерфейс (даже если внутри указатель nil) считается «не nil», если у него есть тип. Здесь тип — *os.PathError.
При выводе fmt.Println(err) видим <nil>, но сравнение err == nil даёт false, так как у интерфейса не обнулён тип.
```
