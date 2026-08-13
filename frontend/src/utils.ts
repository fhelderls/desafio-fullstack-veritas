// formatDate converte uma data ISO ("AAAA-MM-DD", vinda do backend/input
// type="date") para o formato brasileiro ("DD/MM/AAAA") exibido na UI.
// Usa split de string em vez de `new Date(iso)` de proposito: o construtor
// Date interpreta "AAAA-MM-DD" como meia-noite UTC, e ao exibir de volta no
// fuso horario local pode "voltar" um dia - manipular a string diretamente
// evita essa armadilha de fuso horario.
export function formatDate(isoDate: string): string {
  if (!isoDate) return "";
  const [year, month, day] = isoDate.split("-");
  return `${day}/${month}/${year}`;
}

// getTodayISO retorna a data de hoje no formato "AAAA-MM-DD", usando as
// partes locais do Date (ano/mes/dia) em vez de toISOString() - toISOString
// converte pra UTC, o que pode "voltar" um dia perto da meia-noite
// dependendo do fuso horario do navegador.
export function getTodayISO(): string {
  const now = new Date();
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, "0");
  const day = String(now.getDate()).padStart(2, "0");
  return `${year}-${month}-${day}`;
}
