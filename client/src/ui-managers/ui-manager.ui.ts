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
