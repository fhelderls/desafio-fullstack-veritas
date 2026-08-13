
import type { Task, TaskStatus } from "./types";

const API_BASE_URL = "http://localhost:3001";


export async function fetchTasks(): Promise<Task[]> {
  const response = await fetch(`${API_BASE_URL}/tasks`);
  if (!response.ok) {
    throw new Error("Falha ao buscar tarefas");
  }
  return response.json();
}


    export async function createTask(
        title: string,
        description: string,
        status: TaskStatus,
        dueDate: string
    ): Promise<Task> {
        const res = await fetch(`${API_BASE_URL}/tasks`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ title, description, status, due_date: dueDate }),
        });
        if (!res.ok) {
            throw new Error("Falha ao criar tarefa");
        }
        return res.json();
    }

    export async function updateTask(
        id: string,
        title: string,
        description: string,
        status: TaskStatus,
        dueDate: string
    ): Promise<Task> {
        const res = await fetch(`${API_BASE_URL}/tasks/${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ title, description, status, due_date: dueDate }),
        });
        if (!res.ok) {
            throw new Error("Falha ao atualizar tarefa");
        }
        return res.json();
    }

export async function deleteTask(id: string): Promise<void> {
    const res = await fetch(`${API_BASE_URL}/tasks/${id}`, {
        method: "DELETE",
    });
    if (!res.ok) {
        throw new Error("Falha ao deletar tarefa");
    }
}
