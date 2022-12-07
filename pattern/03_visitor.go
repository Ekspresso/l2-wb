// Шаблон проектирования посетитель - это поведенческий шаблон проектирования,
// который позволяет добавлять поведение к структуре без фактического её изменения.

// Использование паттерна посетитель выгодно, когда:
// -требуется много несвязанных операций над структурой объекта,
// -структура объекта не изменяется,
// -необходимо часто добавлять новые операции,
// -алгоритм включает в себя несколько элементов объекта, но желательно управлять им в одном единственном месте.

// Недостатком этого шаблона, является то, что он усложняет расширение объекта,
// поскольку новые структуры эелементов обычно требуют добавления нового visit-метода для каждого нового элемента.

// В данном примере рассматривается существующий интерфейс shapes с методом получить параметр (getParam), и несколько фигур, в структурах которых указаны их параметры,
// с набором методов, удовлетворяющих интерфейсу.

// Для добавления возможности расчёта площади и периметра в существующую модель добавляеются элементы паттерна visitor:
// -создаётся интерфейс visitor с набором методов для посещения элементов. Для каждого элемента свой метод. Все способы посещения должны соответствовать интерфейсу.
// -в интерфейс shapes добавляется метод accept(visitor), где visitor это способ посещения определённого объекта, удовлетворяющий интерфейсу visitor.
// -к методам структур фигур также добавляется метод accept, в котором вызывается оперделённый метод интерфейса visitor (стандартная конструкция).
// -добавляются структуры расчёта периметра и площади, которые содержат поле с одноимённым параметром. Методы структур удовлетворяют интерфейсу visitor.
// (то есть расчитывают площадь или периметр для каждой фигуры в методе посещения фигуры)

package pattern

// package main

import "fmt"

// Интерфейс объекта с фигурами.
type shapes interface {
	getParam()
	accept(visitor)
}

// Структура фигуры "квадрат"
type square struct {
	side int
}

func (s *square) getParam() {
	fmt.Printf("Is a square with side %d\n", s.side)
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

// Структура фигуры "окружность"
type circle struct {
	radius int
}

func (c *circle) getParam() {
	fmt.Printf("Is a circle with radius %d\n", c.radius)
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

// Структура фигуры "прямоугольник"
type rectangle struct {
	a int
	b int
}

func (r *rectangle) getParam() {
	fmt.Printf("Is a rectangle with sides a = %d, b = %d\n", r.a, r.b)
}

func (r *rectangle) accept(v visitor) {
	v.visitForRectangle(r)
}

// Интерфейс паттерна посетитель с методами для посещения каждой фигуры
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForRectangle(*rectangle)
}

// Способ посещения расчёта площади
type areaCalc struct {
	area int
}

// Методы расчёта площади для каждой фигуры
func (a *areaCalc) visitForSquare(s *square) {
	fmt.Println("Calculating area for square")
	a.area = s.side * s.side
}

func (a *areaCalc) visitForCircle(c *circle) {
	fmt.Println("Calculating area for circle")
	a.area = int(float64(3.14) * float64(c.radius) * float64(c.radius))
}

func (a *areaCalc) visitForRectangle(r *rectangle) {
	fmt.Println("Calculating area for rectangle")
	a.area = r.a * r.b
}

// Способ посещения расчёта периметра
type perimeterCalc struct {
	perim int
}

// Методы расчёта периметра для каждой фигуры
func (p *perimeterCalc) visitForSquare(s *square) {
	fmt.Println("Calculating perimeter for square")
	p.perim = 4 * s.side
}

func (p *perimeterCalc) visitForCircle(c *circle) {
	fmt.Println("Calculating perimeter for circle")
	p.perim = int(float64(6.28) * float64(c.radius))
}

func (p *perimeterCalc) visitForRectangle(r *rectangle) {
	fmt.Println("Calculating perimeter for rectangle")
	p.perim = 2 * (r.a + r.b)
}

// Вызов способов посещения для расчёта периметра и площади осуществляется путём передачи методу accept каждой фигуры этого способа посещения.
// Значения хранятся в поле структуры способа посещения и перезаписываются при каждом вызове.
// func main() {
// 	sq := &square{side: 3}
// 	ci := &circle{radius: 5}
// 	re := &rectangle{a: 2, b: 4}

// 	areaCalc := &areaCalc{}
// 	perimeterCalc := &perimeterCalc{}

// 	sq.accept(areaCalc)
// 	fmt.Println(areaCalc.area)
// 	ci.accept(areaCalc)
// 	fmt.Println(areaCalc.area)
// 	re.accept(areaCalc)
// 	fmt.Println(areaCalc.area)

// 	fmt.Println()

// 	sq.accept(perimeterCalc)
// 	fmt.Println(perimeterCalc.perim)
// 	ci.accept(perimeterCalc)
// 	fmt.Println(perimeterCalc.perim)
// 	re.accept(perimeterCalc)
// 	fmt.Println(perimeterCalc.perim)
// }
