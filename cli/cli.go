package cli

import (
	"To-Do_List/handlers"
	"fmt"
)

// RunCLI запускает командный интерфейс пользователя
func RunCLI() {
	fmt.Println("Добро пожаловать в Ежедневник To-D0 List CLI!")
	fmt.Println("Введите команду (help для списка команд):")

	for {
		fmt.Print(">")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "help":
			printHelp()
		case "list":
			listTask()
		case "add":
			addTask()
		case "update":
			updateTask()
		case "delete":
			deleteTask()
		case "exit":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неизвестная комманда. Введите help для списка комманд.")

		}
	}
}

// printHelp выводит список доступных команд
func printHelp() {
	fmt.Println("Доступные команд")
	fmt.Println("	help	- показать список команд")
	fmt.Println("	list	- показать все задачи")
	fmt.Println("	add		- добавить новую задачи")
	fmt.Println("	update	- обновить существующее задачу")
	fmt.Println("	delete	- удалить задачу")
	fmt.Println("	exit	- выйти из программы")
}

// listTasks выводит все задачи
