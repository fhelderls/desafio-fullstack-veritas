import type {TaskStatus} from "./types";

export const STATUS_ORDER: TaskStatus[] = ["todo", "in_progress", "done"];
export const STATUS_LABELS: Record<TaskStatus, string> = {
  todo: "A Fazer",
  in_progress: "Em Progresso",
  done: "Concluídas",
};


