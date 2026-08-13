import { useState } from "react";
import type { Task, TaskStatus } from "../types";
import { STATUS_ORDER, STATUS_LABELS } from "../constants";

interface TaskDetailModalProps {
  task: Task;
  onClose: () => void;
  onEdit: (task: Task) => void;
  onDelete: (id: string) => void;
  onChangeStatus: (task: Task, newStatus: TaskStatus) => void;
}

export function TaskDetailModal({ task, onClose, onEdit, onDelete, onChangeStatus }: TaskDetailModalProps) {
  const [isStatusMenuOpen, setIsStatusMenuOpen] = useState(false);

  function handleSelectStatus(status: TaskStatus) {
    setIsStatusMenuOpen(false);
    if (status !== task.status) {
      onChangeStatus(task, status);
    }
  }

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h3>{task.title}</h3>
        {task.description && <p>{task.description}</p>}

        <div className="status-menu">
          <button
            type="button"
            className="status-menu-trigger"
            onClick={() => setIsStatusMenuOpen((open) => !open)}
          >
            {STATUS_LABELS[task.status]} ▾
          </button>
          {isStatusMenuOpen && (
            <ul className="status-menu-list">
              {STATUS_ORDER.map((status) => (
                <li key={status}>
                  <button
                    type="button"
                    disabled={status === task.status}
                    onClick={() => handleSelectStatus(status)}
                  >
                    {STATUS_LABELS[status]}
                  </button>
                </li>
              ))}
            </ul>
          )}
        </div>

        <div className="modal-actions">
          <button type="button" onClick={onClose}>
            Fechar
          </button>
          <button type="button" onClick={() => onEdit(task)}>
            Editar
          </button>
          <button type="button" className="task-card-delete-btn" onClick={() => onDelete(task.id)}>
            Excluir
          </button>
        </div>
      </div>
    </div>
  );
}
