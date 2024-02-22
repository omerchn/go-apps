package main

import (
	"fmt"

	"slices"
)

func main() {
	tasks := []string{}
	for {
		if len(tasks) == 5 {
			fmt.Println("thank you! the tasks list is ready and i love omer cohen", tasks)
			break
		}
		fmt.Println("please enter a task")
		var task string
		fmt.Scan(&task)

		if slices.Contains(tasks, task) {
			fmt.Println("task already exists")
			continue
		}
		tasks = append(tasks, task)
		println("you added", task)

	}
}
