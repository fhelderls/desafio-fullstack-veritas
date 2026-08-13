import { useState, useEffect, useCallback } from "react";
import type { Task, TaskStatus } from "./types";
import { STATUS_ORDER } from "./constants";
import { fetchTasks, createTask, updateTask, deleteTask } from "./api";
import { Column } from "./components/Column";
import { TaskFormModal } from "./components/TaskFormModal";
import { TaskDetailModal } from "./components/TaskDetailModal";
import "./App.css";

function App() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingTask, setEditingTask] = useState<Task | null>(null);
  const [viewingTask, setViewingTask] = useState<Task | null>(null);

  const loadTasks = useCallback(async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await fetchTasks();
      setTasks(data);
    } catch {
      setError("Não foi possível carregar as tarefas. Verifique se o backend está rodando.");
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    // eslint-disable-next-line react-hooks/set-state-in-effect -- fetch inicial ao montar; sem lib de data-fetching no escopo do desafio
    loadTasks();
  }, [loadTasks]);

  function handleOpenCreate() {
    setEditingTask(null);
    setIsModalOpen(true);
  }

  function handleOpenView(task: Task) {
    setViewingTask(task);
  }

  function handleCloseView() {
    setViewingTask(null);
  }

  function handleOpenEdit(task: Task) {
    setViewingTask(null);
    setEditingTask(task);
    setIsModalOpen(true);
  }

  function handleCloseModal() {
    setIsModalOpen(false);
    setEditingTask(null);
  }

  async function handleSave(title: string, description: string, status: TaskStatus) {
    try {
      if (editingTask) {
        await updateTask(editingTask.id, title, description, status);
      } else {
        await createTask(title, description, status);
      }
      setIsModalOpen(false);
      setEditingTask(null);
      await loadTasks();
    } catch {
      setError("Não foi possível salvar a tarefa.");
    }
  }

  async function handleDelete(id: string) {
    try {
      await deleteTask(id);
      setViewingTask(null);
      await loadTasks();
    } catch {
      setError("Não foi possível excluir a tarefa.");
    }
  }

  async function handleMove(task: Task, newStatus: TaskStatus) {
    try {
      await updateTask(task.id, task.title, task.description, newStatus);
      await loadTasks();
    } catch {
      setError("Não foi possível mover a tarefa.");
    }
  }

  async function handleChangeStatus(task: Task, newStatus: TaskStatus) {
    await handleMove(task, newStatus);
    setViewingTask(null);
  }

  async function handleDropTask(taskId: string, newStatus: TaskStatus) {
    const task = tasks.find((t) => t.id === taskId);
    if (!task || task.status === newStatus) return;
    await handleMove(task, newStatus);
  }

  return (
    <div className="app">
      <header className="app-header">
        <h1>Mini Kanban</h1>
        <button onClick={handleOpenCreate}>+ Nova Tarefa</button>
      </header>

      {loading && <p>Carregando...</p>}
      {error && <p className="error">{error}</p>}

      <div className="board">
        {STATUS_ORDER.map((status) => (
          <Column
            key={status}
            status={status}
            tasks={tasks.filter((task) => task.status === status)}
            onOpen={handleOpenView}
            onDropTask={handleDropTask}
          />
        ))}
      </div>

      {viewingTask && (
        <TaskDetailModal
          task={viewingTask}
          onClose={handleCloseView}
          onEdit={handleOpenEdit}
          onDelete={handleDelete}
          onChangeStatus={handleChangeStatus}
        />
      )}

      {isModalOpen && (
        <TaskFormModal
          initialData={editingTask}
          onClose={handleCloseModal}
          onSave={handleSave}
        />
      )}
    </div>
  );
}

export default App;
