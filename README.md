# Mini Kanban de Tarefas

Aplicação full-stack de gerenciamento de tarefas em formato Kanban. Backend em Go, frontend em React + TypeScript.

## Stack

Backend: Go 1.26, `net/http` da biblioteca padrão (sem framework), `sync.RWMutex` pra proteger o armazenamento em memória contra acesso concorrente.

Frontend: React 19, TypeScript, Vite, ESLint.

## Estrutura do projeto

```
backend/
  models/    definição da struct Task
  store/     armazenamento em memória, CRUD thread-safe
  handlers/  handlers HTTP, validação, JSON, status codes
  main.go    rotas e CORS

frontend/
  src/
    components/  Column, TaskCard, TaskFormModal
    api.ts       chamadas HTTP pro backend
    types.ts     tipos TypeScript compartilhados
    constants.ts STATUS_ORDER / STATUS_LABELS
    App.tsx      estado global e orquestração da UI
```

O backend ficou dividido em três pacotes em vez de um arquivo único porque cada camada não precisa saber como as outras funcionam: a struct não sabe como é guardada, o store não sabe como é exposto via HTTP, os handlers não sabem como os dados são armazenados por dentro. Fica mais fácil testar cada parte isolada e mexer numa sem quebrar a outra.

## Como executar

Backend:
```
cd backend
go run main.go
```
Sobe em `http://localhost:3001`.

Frontend:
```
cd frontend
npm install
npm run dev
```
Sobe em `http://localhost:5173`.

O frontend depende do backend rodando. Sem ele, a tela mostra erro — é o comportamento esperado.

## Endpoints

| Método | Rota          | O que faz                            |
|--------|---------------|----------------------------------------|
| GET    | `/tasks`      | Lista as tarefas                       |
| POST   | `/tasks`      | Cria uma tarefa                        |
| PUT    | `/tasks/{id}` | Atualiza título, descrição e status    |
| DELETE | `/tasks/{id}` | Remove uma tarefa                      |

Status aceitos: `todo`, `in_progress`, `done`.

## Funcionalidades

- CRUD completo de tarefas
- 3 colunas: A Fazer, Em Progresso, Concluídas
- Mover tarefas pelas setas do card ou arrastando (drag-and-drop via HTML5 Drag and Drop API)
- Título obrigatório no formulário
- Backend rejeita status fora do enum (HTTP 400)
- Tela de carregamento e de erro quando o backend está fora do ar
- CORS via middleware manual

## Diagrama de fluxo

[`docs/user-flow.png`](docs/user-flow.png)

## Decisões técnicas

Armazenamento em memória, sem banco — o desafio não pede persistência, então os dados somem a cada restart do backend.

Fetch direto com `fetch` + `useState`/`useEffect`, sem lib de data-fetching (React Query, SWR) — o projeto é pequeno o suficiente pra não precisar.

Setas de mover usam `←`/`→` estilizados em CSS em vez de uma lib de ícones, pra não trazer uma dependência só por isso.

## Limitações conhecidas e melhorias futuras

- Sem persistência: reinicia o backend, perde os dados
- Sem autenticação — qualquer um com acesso à API edita as tarefas
- Drag-and-drop não funciona em touch (celular/tablet); as setas cobrem esse caso
- Sem testes automatizados
- Sem paginação ou busca — não escalaria pra uma lista grande de tarefas
