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
