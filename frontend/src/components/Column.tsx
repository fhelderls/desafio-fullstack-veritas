import { useState } from "react";
import type { Task, TaskStatus } from "../types";
import { STATUS_LABELS } from "../constants";
import { TaskCard } from "./TaskCard";

interface ColumnProps {
  status: TaskStatus;
  tasks: Task[];
  onEdit: (task: Task) => void;
  onDelete: (id: string) => void;
  onMove: (task: Task, newStatus: TaskStatus) => void;
  onDropTask: (taskId: string, newStatus: TaskStatus) => void;
}

export function Column({ status, tasks, onEdit, onDelete, onMove, onDropTask }: ColumnProps) {
  const [isDragOver, setIsDragOver] = useState(false);

  return (
    <div className="column">
      <h3>
        {STATUS_LABELS[status]} ({tasks.length})
      </h3>
      <div
        className={`column-tasks${isDragOver ? " column-tasks-drag-over" : ""}`}
        onDragOver={(e) => {
          e.preventDefault();
          setIsDragOver(true);
        }}
        onDragLeave={() => setIsDragOver(false)}
        onDrop={(e) => {
          e.preventDefault();
          setIsDragOver(false);
          const taskId = e.dataTransfer.getData("text/plain");
          onDropTask(taskId, status);
        }}
      >
        {tasks.map((task) => (
          <TaskCard
            key={task.id}
            task={task}
            onEdit={onEdit}
            onDelete={onDelete}
            onMove={onMove}
          />
        ))}
      </div>
    </div>
  );
}
