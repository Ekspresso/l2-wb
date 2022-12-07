package pattern

// package main

import "fmt"

// Паттерн "команда" - это шаблон проектирования поведения. Он предлагает инкапсулировать запрос как отдельный объект.
// Созданный объект содержит всю информацию о запросе и, следовательно, может выполнять его независимо.
// Эта информация включает в себя имя метода, объект, которому принадлежит метод, и значения параметров метода.

// Четыре термина присущие паттерну команды: command (команда), receiver (получатель), invoker (вызывающий объект) и client.
// Объект command знает о receiver и вызывает его метод. Значения параметров метода приемника хранятся в команде.
// Объект receiver для выполнения этих методов также сохраняется в объекте command путем агрегирования.
// Затем получатель выполняет работу при execute()вызове метода в command.
// Invoker знает, как вызвать выполнение команды, и, при необходимости, ведет учет выполнения команд.
// Вызывающий объект ничего не знает о конкретной команде, он знает только об интерфейсе взаимодействия.
// Объект(ы) invoker, объекты command и объекты receiver нахоятся в распоряжении объекта client,
// клиент решает, какие объекты receiver он назначает объектам command, а какие команды он назначает invoker.
// Клиент решает, какие команды выполнять и в какой момент. Чтобы выполнить команду, он передает объект command объекту invoker.

// В качестве реализации паттерна "команда" приведён пример с телевизором, который может включать / выключать, а также изменять громкость.

// Интерфейс устройства. Содержит в себе методы включения/выключения устройства и увеличения/уменьшения громкости
type device interface {
	on()
	off()
	increaseVolume()
	decreaseVolume()
}

// Интерфейс команд
type command interface {
	execute()
}

// // Структура команды включения. Содержит поле с конкретным устройством.
type onCommand struct {
	device device
}

// Метод команды включения. Выполняет команду включения переданного устройтва.
func (c *onCommand) execute() {
	c.device.on()
}

// // Структура команды выключения. Содержит поле с конкретным устройством.
type offCommand struct {
	device device
}

// Метод команды выключения. Выполняет команду выключения переданного устройтва.
func (c *offCommand) execute() {
	c.device.off()
}

// // Структура команды увеличения громкости. Содержит поле с конкретным устройством.
type increaseVolumeCommand struct {
	device device
}

// Метод команды увеличения громкости. Выполняет команду увеличения громкости переданного устройтва.
func (c *increaseVolumeCommand) execute() {
	c.device.increaseVolume()
}

// // Структура команды уменьшения громкости. Содержит поле с конкретным устройством.
type decreaseVolumeCommand struct {
	device device
}

// Метод команды уменьшения громкости. Выполняет команду уменьшения громкости переданного устройтва.
func (c *decreaseVolumeCommand) execute() {
	c.device.decreaseVolume()
}

// // Структура tv, которая реализует интерфейс устройства device
type tv struct {
	isOn   bool
	volume int
}

// Метод включения tv
func (t *tv) on() {
	t.isOn = true
	fmt.Println("Turning tv on")
}

// Метод выключения tv
func (t *tv) off() {
	t.isOn = false
	fmt.Println("Turning tv off")
}

// Метод увеличения громкости tv
func (t *tv) increaseVolume() {
	// Если tv включен
	if t.isOn {
		// Громкость может быть в диапазоне от 0 до 100
		if t.volume >= 0 && t.volume < 100 {
			t.volume++
			fmt.Printf("Increased volume to %d\n", t.volume)
		} else {
			fmt.Println("Max volume.")
		}
	} else {
		fmt.Println("Cannot change volume, tv is off.")
	}
}

// Метод уменьшения громкости tv
func (t *tv) decreaseVolume() {
	// Если tv включен
	if t.isOn {
		// Громкость может быть в диапазоне от 0 до 100
		if t.volume > 0 && t.volume <= 100 {
			t.volume--
			fmt.Printf("Decreased volume to %d\n", t.volume)
		} else {
			fmt.Println("Min volume.")
		}
	} else {
		fmt.Println("Cannot change volume, tv is off.")
	}
}

// // Телевизор управляется кнопками. Структура button, которая соержит в себе определённую команду.
type button struct {
	command command
}

// Метод структуры button вызывающий выполнение команды.
func (b *button) press() {
	b.command.execute()
}

// // Функция с примером использования
// func main() {
// 	// Инициализация конкретного выключенного телевизора с значением громкости 10.
// 	tv := &tv{
// 		isOn:   false,
// 		volume: 10,
// 	}

// 	// Инициализация команд для телевизора
// 	onCommand := &onCommand{
// 		device: tv,
// 	}
// 	offCommand := &offCommand{
// 		device: tv,
// 	}
// 	increaseVolumeCommand := &increaseVolumeCommand{
// 		device: tv,
// 	}
// 	decreaseVolumeCommand := &decreaseVolumeCommand{
// 		device: tv,
// 	}

// 	// Инициализация кнопок с назначением каждой из них определённой команды
// 	onButton := &button{
// 		command: onCommand,
// 	}
// 	offButton := &button{
// 		command: offCommand,
// 	}
// 	increaseVolumeButton := &button{
// 		command: increaseVolumeCommand,
// 	}
// 	decreaseVolumeButton := &button{
// 		command: decreaseVolumeCommand,
// 	}

// 	// Управление телевизором с помощью кнопок
// 	increaseVolumeButton.press()
// 	onButton.press()
// 	increaseVolumeButton.press()
// 	decreaseVolumeButton.press()
// 	offButton.press()
// }
