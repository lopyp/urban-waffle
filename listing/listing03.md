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
nill
false
Интерфейс равен nill когда тип и данные равны nill, в данном случае тип *os.PathError
у пустого интерфейса таблица методов равна nill, следовательно ему удовлетворяет любой тип данных

```
