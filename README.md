
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
- Persistência em arquivo (`tasks.json`) — os dados sobrevivem a um restart do backend
- Título obrigatório no formulário
- Backend rejeita status fora do enum (HTTP 400)
- Tela de carregamento e de erro quando o backend está fora do ar
- CORS via middleware manual
- Testes unitários do backend (store e handlers)
- Dockerfile multi-stage + docker-compose pros dois serviços

## Diagrama de fluxo

[`docs/user-flow.png`](docs/user-flow.png)

## Decisões técnicas

Persistência em arquivo JSON (`tasks.json`) em vez de banco de dados — simples o suficiente pro escopo do desafio, sem precisar rodar/configurar um banco separado. Cada escrita reescreve o arquivo inteiro, o que é aceitável pro volume de dados esperado aqui, mas não escalaria pra um volume grande de tarefas ou escritas concorrentes intensas.

Fetch direto com `fetch` + `useState`/`useEffect`, sem lib de data-fetching (React Query, SWR) — o projeto é pequeno o suficiente pra não precisar.

Setas de mover usam `←`/`→` estilizados em CSS em vez de uma lib de ícones, pra não trazer uma dependência só por isso.

Dockerfile multi-stage nos dois serviços — a imagem final do backend não carrega o compilador Go, só o binário; a do frontend não carrega Node.js, só os assets estáticos servidos via nginx. Isso deixa as imagens finais bem menores que um build single-stage.

## Limitações conhecidas e melhorias futuras

- Sem autenticação — qualquer um com acesso à API edita as tarefas
- Drag-and-drop não funciona em touch (celular/tablet); as setas cobrem esse caso
- Testes cobrem só o backend (store e handlers) — sem testes automatizados no frontend
- `tasks.json` reescreve o arquivo inteiro a cada mutação — não seria uma estratégia de persistência adequada com volume alto de escritas concorrentes; um banco de verdade resolveria isso
- Sem paginação ou busca — não escalaria pra uma lista grande de tarefas
