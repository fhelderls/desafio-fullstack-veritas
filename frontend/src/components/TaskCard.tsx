import type { Task } from "../types";
import { formatDate } from "../utils";

interface TaskCardProps {
  task: Task;
  onOpen: (task: Task) => void;
}

export function TaskCard({ task, onOpen }: TaskCardProps) {
  return (
    <div
      className="task-card"
      draggable
      onDragStart={(e) => e.dataTransfer.setData("text/plain", task.id)}
      onClick={() => onOpen(task)}
    >
      <h4>{task.title}</h4>
      {task.description && <p>{task.description}</p>}
      {task.due_date && <p className="task-card-due-date">Data limite: {formatDate(task.due_date)}</p>}
      {task.assignee && <p className="task-card-assignee">{task.assignee}</p>}
    </div>
  );
}
