package cli

import (
	"To-Do_List/handlers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RunCLI запускает командный интерфейс пользователя
func RunCLI() {
	fmt.Println("Добро пожаловать в To-Do List CLI!")
	fmt.Println("Введите команду (help для списка команд):")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "help":
			printHelp()
		case "list":
			listTasks()
		case "add":
			addTask(reader)
		case "update":
			updateTask(reader)
		case "delete":
			deleteTask()
		case "exit":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неизвестная команда. Введите help для списка команд.")
		}
	}
}

// printHelp выводит список доступных команд
func printHelp() {
	fmt.Println("Доступные команды:")
	fmt.Println("  help   - показать список команд")
	fmt.Println("  list   - показать все задачи")
	fmt.Println("  add    - добавить новую задачу")
	fmt.Println("  update - обновить существующую задачу")
	fmt.Println("  delete - удалить задачу")
	fmt.Println("  exit   - выйти из программы")
}

// listTasks выводит все задачи
func listTasks() {
	tasks := handlers.GetTasksSlice()
	if len(tasks) == 0 {
		fmt.Println("Задач нет.")
		return
	}
	fmt.Println("Список задач:")
	for _, task := range tasks {
		status := "Не выполнено"
		if task.Done {
			status = "Выполнено"
		}
		fmt.Printf("ID: %d, Title: %s, Status: %s\n", task.ID, task.Title, status)
	}
}

// addTask добавляет новую задачу
func addTask(reader *bufio.Reader) {
	fmt.Print("Введите название задачи: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Задача выполнена? (true/false): ")
	doneStr, _ := reader.ReadString('\n')
	doneStr = strings.TrimSpace(doneStr)
	done := strings.ToLower(doneStr) == "true"

	handlers.AddTask(title, done)
	fmt.Println("Задача добавлена.")
}

// updateTask обновляет существующую задачу
func updateTask(reader *bufio.Reader) {
	fmt.Print("Введите ID задачи для обновления: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Неверный формат ID. Введите целое число.")
		return
	}

	fmt.Print("Введите новое название задачи: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Задача выполнена? (true/false): ")
	doneStr, _ := reader.ReadString('\n')
	doneStr = strings.TrimSpace(doneStr)
	done := strings.ToLower(doneStr) == "true"

	if handlers.UpdateTaskByID(id, title, done) {
		fmt.Println("Задача обновлена.")
	} else {
		fmt.Println("Задача с таким ID не найдена.")
	}
}

// deleteTask удаляет задачу
func deleteTask() {
	fmt.Print("Введите ID задачи для удаления: ")
	var id int
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Неверный формат ID. Введите целое число.")
		return
	}

	if handlers.DeleteTaskByID(id) {
		fmt.Println("Задача удалена.")
	} else {
		fmt.Println("Задача с таким ID не найдена.")
	}
}
