import type { Task, TaskStatus } from "../types";
import { STATUS_ORDER, STATUS_LABELS } from "../constants";

interface TaskCardProps {
  task: Task;
  onEdit: (task: Task) => void;
  onDelete: (id: string) => void;
  onMove: (task: Task, newStatus: TaskStatus) => void;
}

export function TaskCard({ task, onEdit, onDelete, onMove }: TaskCardProps) {
  const currentIndex = STATUS_ORDER.indexOf(task.status);
  const canMoveBack = currentIndex > 0;
  const canMoveForward = currentIndex < STATUS_ORDER.length - 1;

  return (
    <div
      className="task-card"
      draggable
      onDragStart={(e) => e.dataTransfer.setData("text/plain", task.id)}
    >
      <h4>{task.title}</h4>
      {task.description && <p>{task.description}</p>}
      <div className="task-card-actions">
        <button
          className="task-card-nav-btn"
          disabled={!canMoveBack}
          onClick={() => onMove(task, STATUS_ORDER[currentIndex - 1])}
          title={canMoveBack ? `Mover para ${STATUS_LABELS[STATUS_ORDER[currentIndex - 1]]}` : ""}
        >
          ←
        </button>
        <button className="task-card-edit-btn" onClick={() => onEdit(task)}>
          Editar
        </button>
        <button className="task-card-delete-btn" onClick={() => onDelete(task.id)}>
          Excluir
        </button>
        <button
          className="task-card-nav-btn"
          disabled={!canMoveForward}
          onClick={() => onMove(task, STATUS_ORDER[currentIndex + 1])}
          title={canMoveForward ? `Mover para ${STATUS_LABELS[STATUS_ORDER[currentIndex + 1]]}` : ""}
        >
          →
        </button>
      </div>
    </div>
  );
}
