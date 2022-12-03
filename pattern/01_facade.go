// Фасад – это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.

// Фасад – это простой интерфейс для работы со сложной подсистемой, содержащей множество классов.
// Фасад может иметь урезанный интерфейс, не имеющий 100% функциональности, которой можно достичь,
// используя сложную подсистему напрямую. Но он предоставляет именно те фичи, которые нужны клиенту, и скрывает все остальные.
// Фасад полезен, если вы используете какую-то сложную библиотеку со множеством подвижных частей, но вам нужна только часть её возможностей.

// В данной реализации приведён в качестве примера фасад, который позволяет пользователю
// включать и выключать компьютер, не заморачиваясь над тем, как правильно это сделать.

// На практике в фасад может содержать несколько других структур со своими методами.
// Он может вызывать методы этих структур в своих методах, упрощая работу пользователя.

// Закоментировать для теста
package pattern

// Раскомментировать для теста
// package main

import "fmt"

// Структура copmuter, которая имеет свой набор методов для выполнения каких=либо действий
type computer struct {
}

// Метод показа экрана загрузки
func (c computer) ShowLoadingScreen() {
	fmt.Println("Loading...")
}

// Метод показа экрана выхода
func (c computer) ShowExitScreen() {
	fmt.Println("Exit...")
}

// Метод включения системного блока
func (c computer) OnSystemBlock() {
	fmt.Println("Turning on the computer system unit")
}

// Метод выключения системного блока
func (c computer) OffSystemBlock() {
	fmt.Println("Turning off the system unit")
}

// Метод включения монитора
func (c computer) OnScreen() {
	fmt.Println("Turning on the monitor")
}

// Метод выключения монитора
func (c computer) OffScreen() {
	fmt.Println("Turning off the monitor")
}

// Метод входа в систему
func (c computer) LoginSystem() {
	fmt.Println("Log in to the system")
}

// Метод выхода из системы
func (c computer) LogoutSystem() {
	fmt.Println("Log out of the system")
}

// Метод выполнения каких-либо других действий.
func (c computer) DoSomething() {
	fmt.Println("Do something...")
}

// Паттерн фасад используется для включения и выключения компьютера с использованием методов структуры computer.
// Его использование в данном случае упрощает эти процедуры, т.к. для их выполнения нужно вызвать некоторое множество методов структуры computer.
// Пользователю не нужно знать какие методы вызывать для корректного включения или выключения компьютера.
// Он может просто использовать методы структуры FacadeComputer для этого.
type FacadeComputer struct {
	computer *computer
}

// Метод включения компьютера. Последовательно вызывает методы из структуры computer в нужном порядке.
func (f *FacadeComputer) OnComputer() {
	f.computer.OnSystemBlock()
	f.computer.OnScreen()
	f.computer.ShowLoadingScreen()
	f.computer.LoginSystem()
}

// Метод выключения компьютера. Последовательно вызывает методы из структуры computer в нужном порядке.
func (f *FacadeComputer) OffComputer() {
	f.computer.LogoutSystem()
	f.computer.ShowExitScreen()
	f.computer.OffScreen()
	f.computer.OffSystemBlock()
}

// Раскомментировать для теста
// Здесь используется фасад для упрощённого включения и выключения компьютера.
// func main() {
// 	comp := &computer{}
// 	facade := &FacadeComputer{comp}
// 	facade.OnComputer()
// 	facade.OffComputer()
// }
