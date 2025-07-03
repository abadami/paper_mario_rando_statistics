export function updateAverageCard(value: string) {
  const element = document.querySelector<HTMLSpanElement>("#average-value")!;

  element.textContent = value;
}

export function updateDeviationCard(value: string) {
  const element = document.querySelector<HTMLSpanElement>("#deviation-value")!;

  element.textContent = value;
}
