package models

import (
	"finpos-absen-api/config"
)

func GetTodos() ([]config.Todo, error) {
	var todos []config.Todo
	result := config.DB.Find(&todos)

	if result.Error != nil {
		return []config.Todo{}, result.Error
	}

	return todos, nil
}

func CreateTodo(title string, description string) (config.Todo, error) {
	todo := config.Todo{
		Title:       title,
		Description: description,
		Completed:   false,
	}
	result := config.DB.Create(&todo)

	if result.Error != nil {
		return config.Todo{}, result.Error
	}

	return GetTodoById(int(todo.ID))
}

func GetTodoById(id int) (config.Todo, error) {
	var todo config.Todo

	result := config.DB.First(&todo, id)

	if result.Error != nil {
		return config.Todo{}, result.Error
	}

	return todo, nil
}

func UpdateTodo(id int, title string, description string, completed bool) (config.Todo, error) {
	result := config.DB.Updates(&config.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	})

	if result.Error != nil {
		return config.Todo{}, result.Error
	}

	return GetTodoById(id)
}

func DeleteTodo(id int) error {
	result := config.DB.Delete(&config.Todo{}, id)
	return result.Error
}
