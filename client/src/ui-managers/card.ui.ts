//TODO: These can all be 1 function

export function updateCard(elementId: string, value: string) {
  const element = document.querySelector<HTMLSpanElement>(elementId)!;

  element.textContent = value;
}

export function updateAverageCard(value: string) {
  const element = document.querySelector<HTMLSpanElement>("#average-value")!;

  element.textContent = value;
}

export function updateDeviationCard(value: string) {
  const element = document.querySelector<HTMLSpanElement>("#deviation-value")!;

  element.textContent = value;
}

export function updateAverageWinCard(value: string) {
  const element =
    document.querySelector<HTMLSpanElement>("#average-win-value")!;

  element.textContent = value;
}

export function updateBestWinCard(value: string) {
  const element = document.querySelector<HTMLSpanElement>("#best-win-value")!;

  element.textContent = value;
}

export function updateWorstLossCard(value: string) {
  const element = document.querySelector<HTMLSpanElement>("#worst-loss-value")!;

  element.textContent = value;
}
