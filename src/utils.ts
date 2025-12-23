export function getHTMLElement(selector: string) {
  const el = document.querySelector(selector);
  if (!el || !(el instanceof HTMLElement)) {
    throw new Error(`failed to query ${selector} selector`);
  }
  return el;
}
