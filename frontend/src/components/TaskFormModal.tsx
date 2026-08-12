import { useState } from "react";
import type { Task, TaskStatus } from "../types";
import { STATUS_ORDER, STATUS_LABELS } from "../constants";

interface TaskFormModalProps {
  initialData?: Task | null;
  onClose: () => void;
  onSave: (title: string, description: string, status: TaskStatus) => void;
}

export function TaskFormModal({ initialData, onClose, onSave }: TaskFormModalProps) {
  const [title, setTitle] = useState(initialData?.title ?? "");
  const [description, setDescription] = useState(initialData?.description ?? "");
  const [status, setStatus] = useState<TaskStatus>(initialData?.status ?? "todo");

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    if (!title.trim()) return;
    onSave(title, description, status);
  }

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h3>{initialData ? "Editar Tarefa" : "Nova Tarefa"}</h3>
        <form onSubmit={handleSubmit}>
          <label>
            Título:
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
            />
          </label>
          <label>
            Descrição:
            <textarea
              value={description}
              onChange={(e) => setDescription(e.target.value)}
            />
          </label>
          <label>
            Status:
            <select
              value={status}
              onChange={(e) => setStatus(e.target.value as TaskStatus)}
            >
              {STATUS_ORDER.map((s) => (
                <option key={s} value={s}>
                  {STATUS_LABELS[s]}
                </option>
              ))}
            </select>
          </label>
          <div className="modal-actions">
            <button type="button" onClick={onClose}>
              Cancelar
            </button>
            <button type="submit">Salvar</button>
          </div>
        </form>
      </div>
    </div>
  );
}
