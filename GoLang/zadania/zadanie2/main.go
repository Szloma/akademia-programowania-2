package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Task struct {
	Name      string
	Deadline  time.Time
	Completed bool
}

func saveTask(task Task) {
	f, err := os.OpenFile("zadania.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Błąd podczas zapisywania zadania:")
		log.Fatal(err)
		return
	}
	defer f.Close()
	taskS := fmt.Sprintf("%s|%s|%t\n", task.Name, task.Deadline.Format("2006-01-02 15:04"), task.Completed)
	_, err = f.WriteString(taskS)
	if err != nil {
		fmt.Println("Błąd podczas zapisywania zadania:")
		log.Fatal(err)
		return
	}
	fmt.Println("Zadanie zostało zapisane poprawnie")
}

func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nazwa zadania: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)
	fmt.Print("Deadline (format YYYY-MM-DD HH:MM): ")
	deadInput, _ := reader.ReadString('\n')
	deadInput = strings.TrimSpace(deadInput)
	d, err := time.Parse("2006-01-02 15:04", deadInput)
	if err != nil {
		fmt.Println("Błędna data.")
		return
	}
	task := Task{
		Name:      taskName,
		Deadline:  d,
		Completed: false,
	}
	fmt.Println("Zadanie dodane")
	saveTask(task)
}
func displayTasks() {
	f, err := os.Open("zadania.txt")
	if err != nil {
		fmt.Println("Błąd podczas odczytywania pliku:")
		log.Fatal(err)
		return
	}
	defer f.Close()
	fmt.Println("\nTwoje zadania:")
	s := bufio.NewScanner(f)
	for s.Scan() {
		taskI := strings.Split(s.Text(), "|")
		taskN := taskI[0]
		d, _ := time.Parse("2006-01-02 15:04", taskI[1])
		completed := taskI[2] == "true"
		fmt.Printf("- Nazwa: %s, Deadline: %s, Ukończone: %t\n", taskN, d.Format("2006-01-02 15:04"), completed)
		if err := s.Err(); err != nil {
			fmt.Println("Błąd podczas odczytywania pliku:")
			log.Fatal(err)
		}
	}
}

func main() {

	fmt.Printf("Menedzer codziennych zadań v1")
	for {
		fmt.Println("Opcje do wyboru:")
		fmt.Println("DODAJ zadanie")
		fmt.Println("USUŃ zadanie")
		fmt.Println("ZOBACZ zadania")
		fmt.Println("ZAKOŃCZ")
		var choice string
		fmt.Scanln(&choice)
		choice = strings.ToLower(choice)
		switch choice {
		case "dodaj":
			addTask()
		case "zobacz":
			displayTasks()
		case "zakoncz":
			return
		default:
			fmt.Println("Niepoprawny wybór")
		}
	}
}
