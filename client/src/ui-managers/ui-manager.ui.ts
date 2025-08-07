//We might not need this at all
export function toggleElementVisibility(element: HTMLElement) {
  if (element.className.split(" ").find((className) => className === "hide")) {
    element.className = element.className.replace("hide", "");
  } else {
    element.className = `${element.className} hide`;
  }
}

export function hideElement(element: HTMLElement) {
  if (element.className.split(" ").find((className) => className === "hide")) {
    return;
  } else {
    element.className = `${element.className} hide`;
  }
}

export function showElement(element: HTMLElement) {
  if (element.className.split(" ").find((className) => className === "hide")) {
    element.className = element.className.replace("hide", "");
  } else {
    return;
  }
}

export function enableLoading() {
  const statisticsElement = document.querySelector<HTMLDivElement>(
    "#statistic-information"
  );
  const loadingElement =
    document.querySelector<HTMLDivElement>("#loading-section");

  hideElement(statisticsElement as HTMLElement);
  showElement(loadingElement as HTMLElement);
}

export function disableLoading() {
  const statisticsElement = document.querySelector<HTMLElement>(
    "#statistic-information"
  );
  const loadingElement =
    document.querySelector<HTMLElement>("#loading-section");

  showElement(statisticsElement as HTMLElement);
  hideElement(loadingElement as HTMLElement);
}
