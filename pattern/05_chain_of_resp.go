// Цепочка обязанностей (ответственности/вызовов) – это поведенческий паттерн проектирования, который позволяет
// передавать запросы последовательно по цепочке обработчиков. Каждый последующий обработчик решает,
// может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

// Каждый обработчик хранит ссылку на следующий.

// Может использоваться в случае, когда необходимо проверить какие-либо данные через цепочку обработчиков
// и прервать выполнение проверки, если с данными что-то не так при проверке одним из обработчиков цепи.

// Также обработчик может прерывать выполнение, если смог обработать запрос. Пример использования: обработка событий,
// создаваемых классами графического интерфейса при работе с пользователем.

// Здесь показан пример паттерна цепочки обязанностей на примере больницы.
// Пациент сначала отправляется на прием, а затем, основываясь на текущем статусе пациента, отправляется следующему обработчику в цепочке.

package main

// package pattern

import "fmt"

// Интерфейс handler. Ему должны удовлетворять все обработчики.
type department interface {
	execute(*patient)
	setNext(department)
}

// Структура регистратуры. Хранит ссылку на следующий обработчик
type reception struct {
	next department
}

// Метод регистратуры, обрабатывает структуру Пациент и передаёт её следующему обработчику.
func (r *reception) execute(p *patient) {
	// Если пациент зарегистрирован, то выводит сообщение об этом и идёт дальше, иначе регистрирует его.
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

// Метод регистратуры, использующийся для задания последовательности
func (r *reception) setNext(next department) {
	r.next = next
}

// Структура доктор. Хранит ссылку на следующий обработчик
type doctor struct {
	next department
}

// Метод доктора, обрабатывает структуру Пациент и передаёт её следующему обработчику.
func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

// Метод доктора, использующийся для задания последовательности
func (d *doctor) setNext(next department) {
	d.next = next
}

// Структура медицинского кабинета. Хранит ссылку на следующий обработчик
type medical struct {
	next department
}

// Метод медицинского кабинета, обрабатывает структуру Пациент и передаёт её следующему обработчику.
func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

// Метод медицинского кабинета, использующийся для задания последовательности
func (m *medical) setNext(next department) {
	m.next = next
}

// Структура кассы. Хранит ссылку на следующий обработчик
type cashier struct {
	next department
}

// Метод кассы, обрабатывает структуру Пациент. Завершающий обработчик в последовательности.
func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

// Метод кассы, использующийся для задания последовательности
func (c *cashier) setNext(next department) {
	c.next = next
}

// Структура пациента. Хранит данные о прохождении каждого обработчика.
type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// func main() {
// 	// Инициализация кассы
// 	cashier := &cashier{}

// 	// Инициализация медицинского кабинета и установка значения следующего обработчика на кассу.
// 	medical := &medical{}
// 	medical.setNext(cashier)

// 	// Инициализация доктора и установка значения следующего обработчика на медицинский кабинет.
// 	doctor := &doctor{}
// 	doctor.setNext(medical)

// 	// Инициализация регистратуры и установка значения следующего обработчика на доктора.
// 	reception := &reception{}
// 	reception.setNext(doctor)

// 	// Инициализация пациента и задание его имени.
// 	patient := &patient{name: "abc"}

// 	// Передача структуры пациента в работу обработчику Регистратура
// 	reception.execute(patient)
// }
