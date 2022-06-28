/**
 * Small, dependency-free helpers shared by the HTML renderer and the
 * GitHub-style formatter.
 */

/** Escape a string for safe interpolation into HTML text/attributes. */
export function escapeHtml(input: string): string {
  return input
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#39;");
}

/** Escape a pipe so Markdown table cells are not broken. */
export function escapePipe(input: string): string {
  return input.replace(/\|/g, "\\|");
}

/** Format a ratio (0..1+) to two decimals. */
export function ratio(n: number): string {
  return n.toFixed(2);
