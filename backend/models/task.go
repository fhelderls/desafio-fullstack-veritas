package models

// Task e a entidade principal do Kanban. DueDate e Assignee ficam vazios
// ("") quando a tarefa nao tem data de vencimento ou responsavel definido -
// sao campos opcionais.
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
	Assignee    string `json:"assignee"`
}
