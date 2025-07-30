//TODO: These can all be 1 function

export function updateCard(elementId: string, value: string) {
  const element = document.querySelector<HTMLSpanElement>(elementId)!;

  element.textContent = value;
}
