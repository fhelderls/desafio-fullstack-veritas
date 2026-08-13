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
