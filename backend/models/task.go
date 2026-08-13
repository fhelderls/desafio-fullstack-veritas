package models

// Task e a entidade principal do Kanban. DueDate fica vazio ("") quando a
// tarefa nao tem data de vencimento definida - e um campo opcional.
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}
