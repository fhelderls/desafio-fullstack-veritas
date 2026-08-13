import { useState } from "react";
import type { Task, TaskStatus } from "../types";
import { STATUS_ORDER, STATUS_LABELS, ASSIGNEES } from "../constants";
import { getTodayISO } from "../utils";

interface TaskFormModalProps {
  initialData?: Task | null;
  onClose: () => void;
  onSave: (title: string, description: string, status: TaskStatus, dueDate: string, assignee: string) => void;
}

export function TaskFormModal({ initialData, onClose, onSave }: TaskFormModalProps) {
  const [title, setTitle] = useState(initialData?.title ?? "");
  const [description, setDescription] = useState(initialData?.description ?? "");
  const [status, setStatus] = useState<TaskStatus>(initialData?.status ?? "todo");
  const [dueDate, setDueDate] = useState(initialData?.due_date || getTodayISO());
  const [assignee, setAssignee] = useState(initialData?.assignee ?? "");

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    if (!title.trim()) return;
    onSave(title, description, status, dueDate, assignee);
  }

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <button type="button" className="modal-close-btn" onClick={onClose} aria-label="Fechar">
          ×
        </button>
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
            Data de vencimento:
            <input
              type="date"
              value={dueDate}
              onChange={(e) => setDueDate(e.target.value)}
            />
          </label>
          <label>
            Responsável:
            <select
              value={assignee}
              onChange={(e) => setAssignee(e.target.value)}
              required
            >
              <option value="" disabled></option>
              {ASSIGNEES.map((name) => (
                <option key={name} value={name}>
                  {name}
                </option>
              ))}
            </select>
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
            <button type="submit">Salvar</button>
          </div>
        </form>
      </div>
    </div>
  );
}
