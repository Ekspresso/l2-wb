// Паттерн Builder (Строитель) отделяет конструирование сложного объекта от его представления,
// так что один и тот же процесс строительства может создать разные представления.
// Используется для инкапсуляции логики построения объекта.

// Паттерн стоитель стоит использовать:
// когда создаваемый объект большой (сложный) и его создание требует нескольких шагов.
// когда необходимо создавать несколько разных версий объекта с разными параметрами.

// В даннной реализации представлен паттерн проектирования Строитель на примере создания бургеров.

// Раскомментировать для теста
// package main
// import "fmt"

// Закомментировать для теста
package pattern

// Объявление новых типов данных для удобства работы с конкретными примерами. Определение продуктов как типов данных.
type bread string
type cheese string
type cutlet string
type sauce string

// Задание набора базовых значений каждого типа продуктов. В дальнейшем из этого набора можно будет создавать конкретную реализацию (бургер).
const (
	blackBread bread = "Black bread"
	whiteBread bread = "White bread"

	parmesan   cheese = "Parmesan"
	blueCheese cheese = "Blue cheese"
	maasdam    cheese = "Maasdam"

	chickenCutlet cutlet = "Chicken cutlet"
	beefCutlet    cutlet = "Beef cutlet"

	cheeseSauce sauce = "Cheese sauce"
	mayonnaise  sauce = "Mayonnaise"
	ketchup     sauce = "Ketchup"
	mustard     sauce = "Mustard"
)

// Структура объекта (бургера)
type burger struct {
	bread  bread
	cheese cheese
	cutlet cutlet
	sauces []sauce
}

// Интерфейс конструктора объектов (бургеров). Для создания конкретного представления объекта необходимо реализовать конструктор этого представления,
// удовлетворяющий данному интерфейсу.
type burgerBuilder interface {
	getBread()
	getCheese()
	getCutlet()
	getSauces()
	getBurger() burger
}

// Конструктор чизбургера (конкретного объекта с определнным набором параметров). Он содержит все методы, объявленные в интерфейсе конструктора.
// Каждый метод позволяет определить значение конкретного параметра. Последний метод возвращает созданный и заполненный объект.
type cheeseburger struct {
	burger
}

func (c *cheeseburger) getBread() {
	c.burger.bread = whiteBread
}

func (c *cheeseburger) getCheese() {
	c.burger.cheese = parmesan
}

func (c *cheeseburger) getCutlet() {
	c.burger.cutlet = beefCutlet
}

func (c *cheeseburger) getSauces() {
	c.burger.sauces = []sauce{cheeseSauce, ketchup, mayonnaise}
}

func (c *cheeseburger) getBurger() burger {
	return c.burger
}

// Конструктор чикенбургера. Второй пример создания объекта со своим набором параметров.
type chickenburger struct {
	burger
}

func (c *chickenburger) getBread() {
	c.burger.bread = blackBread
}

func (c *chickenburger) getCheese() {
	c.burger.cheese = maasdam
}

func (c *chickenburger) getCutlet() {
	c.burger.cutlet = chickenCutlet
}

func (c *chickenburger) getSauces() {
	c.burger.sauces = []sauce{mustard, ketchup, mayonnaise}
}

func (c *chickenburger) getBurger() burger {
	return c.burger
}

// Директор - элемент, который создаёт объект через интерфейс builder.
type director struct {
	builder burgerBuilder
}

func (d *director) setBuilder(builder burgerBuilder) {
	d.builder = builder
}

func (d *director) buildBurger() burger {
	d.builder.getBread()
	d.builder.getCheese()
	d.builder.getCutlet()
	d.builder.getSauces()

	return d.builder.getBurger()
}

// Раскомментировать для теста
// func main() {
// 	cheeseburgerBuilder := &cheeseburger{}
// 	director := &director{
// 		builder: cheeseburgerBuilder,
// 	}
// 	cheeseburger := director.buildBurger()
// 	fmt.Println(cheeseburger)

// 	director.setBuilder(&chickenburger{})
// 	chickenburger := director.buildBurger()
// 	fmt.Println(chickenburger)
// }
