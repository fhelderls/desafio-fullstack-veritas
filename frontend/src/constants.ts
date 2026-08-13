import type {TaskStatus} from "./types";

export const STATUS_ORDER: TaskStatus[] = ["todo", "in_progress", "done"];
export const STATUS_LABELS: Record<TaskStatus, string> = {
  todo: "A Fazer",
  in_progress: "Em Progresso",
  done: "Concluídas",
};

// ASSIGNEES e uma lista fixa de responsaveis de exemplo (nao ha cadastro de
// usuarios no escopo deste desafio) - precisa bater com validAssignees em
// backend/handlers/tasks.go.
export const ASSIGNEES = ["Felipe", "Hellen"];
