# Mini Kanban Veritas

Aplicação full-stack de gerenciamento de tarefas em formato Kanban. Backend em Go, frontend em React + TypeScript.

## Stack

Backend: Go 1.26, `net/http` da biblioteca padrão (sem framework), `sync.RWMutex` pra proteger o armazenamento em memória contra acesso concorrente.

Frontend: React 19, TypeScript, Vite, ESLint.

## Estrutura do projeto

```
backend/
  models/      definição da struct Task
  store/       armazenamento em memória + persistência em JSON, CRUD thread-safe
  handlers/    handlers HTTP, validação, JSON, status codes
  main.go      rotas e CORS
  Dockerfile   build multi-stage do backend

frontend/
  src/
    components/  Column, TaskCard, TaskDetailModal, TaskFormModal
    api.ts       chamadas HTTP pro backend
    types.ts     tipos TypeScript compartilhados
    constants.ts STATUS_ORDER / STATUS_LABELS / ASSIGNEES
    utils.ts     formatação de data
    App.tsx      estado global e orquestração da UI
  Dockerfile     build multi-stage do frontend (assets estáticos + nginx)

docker-compose.yml   sobe backend e frontend juntos
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

Com Docker (backend + frontend juntos, build de produção):
```
docker compose up --build
```
Backend em `http://localhost:3001`, frontend em `http://localhost:8090`.

## Testes

Testes unitários do backend (store e handlers):
```
cd backend
go test ./...
```

## Endpoints

| Método | Rota          | O que faz                            |
|--------|---------------|----------------------------------------|
| GET    | `/tasks`      | Lista as tarefas                                            |
| POST   | `/tasks`      | Cria uma tarefa                                             |
| PUT    | `/tasks/{id}` | Atualiza título, descrição, status, data e responsável      |
| DELETE | `/tasks/{id}` | Remove uma tarefa                                           |

Status aceitos: `todo`, `in_progress`, `done`. Responsáveis aceitos: `Felipe`, `Hellen` (lista fixa de exemplo — não há cadastro de usuários no escopo deste desafio). Título e responsável são obrigatórios; data de vencimento é opcional e aceita o formato `AAAA-MM-DD`.

## Funcionalidades

- CRUD completo de tarefas
- 3 colunas: A Fazer, Em Progresso, Concluídas
- Clique no card abre um modal de detalhes, com menu cascata pra mudar o status, e atalhos pra editar ou excluir
- Mover entre colunas pelo menu cascata ou arrastando o card (drag-and-drop via HTML5 Drag and Drop API)
- Data de vencimento (sugere o dia de hoje por padrão) e responsável (obrigatório, lista fixa de exemplo)
- Persistência em arquivo (`tasks.json`) — os dados sobrevivem a um restart do backend
- Título e responsável obrigatórios no formulário
- Backend rejeita status, data ou responsável fora do esperado (HTTP 400)
- Tela cheia de carregamento e de erro (com botão de tentar novamente) quando o backend está fora do ar
- CORS via middleware manual
- Testes unitários do backend (store e handlers)
- Dockerfile multi-stage + docker-compose pros dois serviços

## Diagramas

Fluxo de uso: [`docs/user-flow.png`](docs/user-flow.png)

Fluxo de dados (requisição → handler → store → arquivo): [`docs/data-flow.png`](docs/data-flow.png)

## Decisões técnicas

Persistência em arquivo JSON (`tasks.json`) em vez de banco de dados — simples o suficiente pro escopo do desafio, sem precisar rodar/configurar um banco separado. Cada escrita reescreve o arquivo inteiro, o que é aceitável pro volume de dados esperado aqui, mas não escalaria pra um volume grande de tarefas ou escritas concorrentes intensas.

Fetch direto com `fetch` + `useState`/`useEffect`, sem lib de data-fetching (React Query, SWR) — o projeto é pequeno o suficiente pra não precisar.

Dockerfile multi-stage nos dois serviços — a imagem final do backend não carrega o compilador Go, só o binário; a do frontend não carrega Node.js, só os assets estáticos servidos via nginx. Isso deixa as imagens finais bem menores que um build single-stage.

Responsável como lista fixa (`Felipe`, `Hellen`) direto na `Task`, em vez de uma entidade `User` própria — cobre a necessidade de atribuição sem o custo de cadastro/autenticação de usuários, fora do escopo do desafio.

## Limitações conhecidas e melhorias futuras

- Sem autenticação — qualquer um com acesso à API edita as tarefas
- Drag-and-drop não funciona em touch (celular/tablet) — a API nativa do HTML5 não suporta toque sem código adicional; o menu cascata de status (clicando no card) cobre esse caso
- Testes cobrem só o backend (store e handlers) — sem testes automatizados no frontend
- `tasks.json` reescreve o arquivo inteiro a cada mutação — não seria uma estratégia de persistência adequada com volume alto de escritas concorrentes; um banco de verdade resolveria isso
- Sem paginação ou busca — não escalaria pra uma lista grande de tarefas
- Responsáveis são uma lista fixa de 2 nomes de exemplo, sem cadastro de usuários — evoluiria pra uma entidade `User` própria num escopo maior
