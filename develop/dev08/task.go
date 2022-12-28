package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Необходимо реализовать свою собственную UNIX-шелл-утилиту с поддержкой ряда простейших команд:

// - cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
// - pwd - показать путь до текущего каталога
// - echo <args> - вывод аргумента в STDOUT
// - kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
// - ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

// Так же требуется поддерживать функционал fork/exec-команд

// Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

// *Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
// в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
// и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

// Command - интерфейс команд программы. Все команды утилиты удовлетворяют этому интерфейсу.
type Command interface {
	Exec(args ...string) ([]byte, error)
}

// cmdEcho - структура команды echo
type cmdEcho struct {
}

// Функция выполнения команды echo
func (e *cmdEcho) Exec(args ...string) ([]byte, error) {
	return exec.Command("echo", args...).Output()
}

// cmdPwd - структура команды pwd
type cmdPwd struct {
}

// Функция выполнения команды pwd
func (p *cmdPwd) Exec(args ...string) ([]byte, error) {
	dir, err := os.Getwd()
	return []byte(dir), err
}

// cmdPs - структура команды ps
type cmdPs struct {
}

// Функция выполнения команды ps
func (p *cmdPs) Exec(args ...string) ([]byte, error) {
	return exec.Command("ps").Output()
}

// cmdCd - структура команды cd
type cmdCd struct {
}

// Функция выполнения команды cd
func (c *cmdCd) Exec(args ...string) ([]byte, error) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	dir, err = os.Getwd()
	if err != nil {
		return nil, err
	}
	return []byte(dir), nil
}

// cmdKill - структура команды kill
type cmdKill struct {
}

// Функция выполнения команды kill. Убивает процесс, переданный в 0 аргументе. Не поддерживает -сигнал.
// Можно указать через пробел несколько процессов
func (c *cmdKill) Exec(args ...string) ([]byte, error) {
	for _, arg := range args {
		pid, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			return nil, err
		}

		err = proc.Kill()
		if err != nil {
			return nil, err
		}
	}
	return []byte("killed"), nil
}

// Shell - основная структура программы. Хранит в себе переменную типа Command,
// через которую в последствии вызываются обработчики определённых команд.
type Shell struct {
	command Command
}

// StartComands - функция обработки комманд.
// Разбивает каждую строку команды на саму команду и аргументы.
// Хатем в зависимости от команды выполняет определённые действия.
func (s *Shell) StartComands(cmds []string) int {
	for _, comArgs := range cmds {
		args := strings.Split(comArgs, " ")
		command := args[0]
		if len(args) > 1 {
			args = args[1:]
		}
		switch command {
		case "cd":
			s.command = &cmdCd{}
		case "pwd":
			s.command = &cmdPwd{}
		case "echo":
			s.command = &cmdEcho{}
		case "kill":
			s.command = &cmdKill{}
		case "ps":
			s.command = &cmdPs{}
		case "q":
			fmt.Println("Exiting the program...")
			return 1
		default:
			fmt.Printf("error: the command '%s' does not exist\n", args[0])
			continue
		}
		res, err := s.command.Exec(args...)
		fmt.Fprintln(os.Stdout, string(res))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return 0
}

// Run - функция структуры Shell, запускающая бесконечный цикл, который обрабатывает стандартный ввод.
// Разбивает строку на отдельные команды через "|" и передаёт обработчику команд.
func (s *Shell) Run() {
	scan := bufio.NewScanner(os.Stdin)
	q := 0
	for q == 0 {
		fmt.Print("Enter the command: ")
		if scan.Scan() {
			str := scan.Text()
			cmds := strings.Split(str, " | ")
			q = s.StartComands(cmds)
		}
	}
}

func main() {
	s := &Shell{}
	s.Run()
}
