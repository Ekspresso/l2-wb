// Фабричный метод (Factory method) также известный как Виртуальный конструктор (Virtual Constructor) -
// пораждающий шаблон проектирования, определяющий общий интерфейс создания объектов в родительском классе (пользовательский тип struct)
// и позволяющий изменять создаваемые объекты в дочерних классах (пользовательских типах).

// Шаблон позволяет структуре делегировать создание объектов подструктурам. Используется, когда:
// -Структуре заранее неизвестно, объекты каких подструктур ему нужно создать.
// -Обязанности делегируются подструктуре, а знания о том, какая подструктура принимает эти обязанности, локализованы.
// -Создаваемые объекты родительской структуры специализируются подструктурами.

// package main

package pattern

import "fmt"

// Интерфейс для транспортных объектов
type iTransport interface {
	// Установить имя для транспорта
	setName(n string)
	// Получить имя транспорта
	getName() string
	// Установить скорость транспорта
	setSpeed(s uint)
	// Получить скорость транспорта
	getSpeed() uint
}

// Общая структура транспортного средства
type transport struct {
	name  string
	speed uint
}

// Методы, удовлетворяющие интерфейсу транспортных объектов
// Установить имя для транспорта
func (t *transport) setName(n string) {
	t.name = n
}

// Получить имя транспорта
func (t *transport) getName() string {
	return t.name
}

// Установить скорость транспорта
func (t *transport) setSpeed(s uint) {
	t.speed = s
}

// Получить скорость транспорта
func (t *transport) getSpeed() uint {
	return t.speed
}

// Структура транспорта Квадрокоптер, встраивающая в себя структуру Транспорт
type quadcopter struct {
	transport
}

// Создание объекта Квадрокоптер с соответствующим названием и скоростью 14
func newQuadcopter() iTransport {
	return &quadcopter{
		transport: transport{
			name:  "Quadcopter",
			speed: 14,
		},
	}
}

// Структура транспорта Скутер, встраивающая в себя структуру Транспорт
type electricScooter struct {
	transport
}

// Создание объекта Скутер с соответствующим названием и скоростью 4
func newElectricScooter() iTransport {
	return &electricScooter{
		transport: transport{
			name:  "Scooter",
			speed: 4,
		},
	}
}

// Фабричный метод. В зависимоти от переданного строкового значения фабричный метод возвращает созданный запрашиваемый объект.
func getTransport(tt string) (iTransport, error) {
	// Создание скутера
	if tt == "scooter" {
		return newElectricScooter(), nil
	}
	// Создание квадрокоптера
	if tt == "quadcopter" {
		return newQuadcopter(), nil
	}
	// Если нет подходящего значения, то вернётся ошибка
	return nil, fmt.Errorf("Wrong type")
}

// func main() {
// 	// Инициализация новых транспортных средств с созданием объектов через фабричный метод.
// 	scooter, _ := getTransport("scooter")
// 	quad, _ := getTransport("quadcopter")

// 	fmt.Println(scooter)
// 	fmt.Println(quad)
// }
