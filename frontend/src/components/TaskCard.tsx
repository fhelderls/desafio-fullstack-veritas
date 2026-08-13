import type { Task } from "../types";

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
    </div>
  );
}
