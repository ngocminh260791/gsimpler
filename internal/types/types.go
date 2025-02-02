// Code generated by goctl. DO NOT EDIT.
package types

type CreateTodoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TodoResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type GetTodoRequest struct {
	ID string `json:"id"`
}
