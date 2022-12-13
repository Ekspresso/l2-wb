Что выведет программа? Объяснить вывод программы.

```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```

Ответ:
Программа выведет сообщение error и завершит работу.
Это происходит потому, что тип переменной err является интерфейсом error, которому удовлетвряет структура customError, поэтому возможно выполнить присваивание.
Но в таком случае переменная всё ещё будет типа interface, значение которого не nil. 
А вот если привести переменную err к типу customError, то значение структуры err будет nil.