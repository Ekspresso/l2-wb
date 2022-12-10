// Паттерн State относится к поведенческим паттернам уровня объекта.

// Паттерн State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния
// и является объектно-ориентированной реализацией конечного автомата.
// Поведение объекта изменяется настолько, что создается впечатление, будто изменился класс объекта.

// Паттерн должен применяться:
// -когда поведение объекта зависит от его состояния
// -поведение объекта должно изменяться во время выполнения программы
// -состояний достаточно много и использовать для этого условные операторы, разбросанные по коду, достаточно затруднительно

// Очень важным нюансом, отличающим этот паттерн от Стратегии, является то, что и контекст,
// и сами конкретные состояния могут знать друг о друге и инициировать переходы от одного состояния к другому.

// В данном примере приведена реализация структуры door с 2 состояниями: открыто, закрыто.
// В отурытую дверь можно войти, в закрытую - нельзя.
// Открытую дверь можно закрыть, но нельзя открыть.
// Закрытую дверь можно открыть, но нельзя закрыть.
// Начальное состояние двери является закрытым.

// package main

package pattern

import (
	"fmt"
	// "log"
)

// Интерфейс состояния. Все состояния дожны уовлетворять этому интерфейсу.
type state interface {
	enterDoor() error
	openDoor() error
	closeDoor() error
}

// Структура основного объекта door. Методы входа, открытия и закрытия структуры вызывают соответствующие методы поля текущего состояния
type door struct {
	open  state
	close state

	currentState state
}

// Конструктор создания новой двери
func newDoor() *door {
	d := &door{}
	openState := &openState{
		door: d,
	}
	closeState := &closeState{
		door: d,
	}
	// Инициализация полей структуры
	d.open = openState
	d.close = closeState
	d.setState(closeState)
	return d
}

// Метод структуры door устанавливающий состояние двери в требуемое положение
func (d *door) setState(s state) {
	d.currentState = s
}

// Метод структуры door входа в дверь
func (d *door) enterDoor() error {
	return d.currentState.enterDoor()
}

// Метод структуры door открытия двери
func (d *door) openDoor() error {
	return d.currentState.openDoor()
}

// Метод структуры door закрытия двери
func (d *door) closeDoor() error {
	return d.currentState.closeDoor()
}

// Структура состояния "Открыто"
type openState struct {
	door *door
}

// Метод структуры состояния "Открыто" для входа в дверь. Не меняет текущего состояния
func (o *openState) enterDoor() error {
	fmt.Println("Enter the door")
	return nil
}

// Метод структуры состояния "Открыто" для открытия двери. Вызывает ошибку, т.к. дверь уже открыта.
func (o *openState) openDoor() error {
	return fmt.Errorf("The door is already open")
}

// Метод структуры состояния "Открыто" для для закрытия двери. Меняет состояние двери на "Закрыто"
func (o *openState) closeDoor() error {
	fmt.Println("Closing the door")
	o.door.setState(o.door.close)
	return nil
}

// Структура состояния "Закрыто"
type closeState struct {
	door *door
}

// Метод структуры состояния "Закрыто" для входа в дверь. Вызывает ошибку, т.к. невозможно войти в закрытую дверь
func (c *closeState) enterDoor() error {
	return fmt.Errorf("It is impossible to enter the door. the door is closed")
}

// Метод структуры состояния "Закрыто" для открытия двери. Меняет состояние двери на "Открыто"
func (c *closeState) openDoor() error {
	fmt.Println("Opening the door")
	c.door.setState(c.door.open)
	return nil
}

// Метод структуры состояния "Закрыто" для закрытия двери. Вызывает ошибку, т.к. закрытую дверь нельзя закрыть
func (c *closeState) closeDoor() error {
	return fmt.Errorf("The door is already closed")
}

// func main() {
// // Создание новой двери
// 	door := newDoor()
// // Открытие двери
// 	err := door.openDoor()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}
// // Вход в дверь
// 	err = door.enterDoor()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}
// // Закрытие двери
// 	err = door.closeDoor()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}
// }
