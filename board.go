package main

import "fmt"

// truncateString обрезает длинное название задачи и добавляет ...
func truncateString(s string, maxLen int) string {
runes := []rune(s)
if len(runes) > maxLen {
return string(runes[:maxLen-3]) + "..."
}
return s
}

// printBoard выводит красивую ASCII-таблицу с границами из +, -, |
func printBoard(tasks []Task) {
if len(tasks) == 0 {
fmt.Println("Список задач пуст.")
return
}

fmt.Println("+----+----------------------+--------+")
fmt.Printf("| %-2s | %-20s | %-6s |\n", "ID", "Title", "Status")
fmt.Println("+----+----------------------+--------+")

for _, task := range tasks {
shortTitle := truncateString(task.Title, 20)
fmt.Printf("| %-2d | %-20s | %-6s |\n", task.ID, shortTitle, task.Status)
}

fmt.Println("+----+----------------------+--------+")
}15:25
