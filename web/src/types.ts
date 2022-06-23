/**
 * Type definitions mirroring the deterministic JSON emitted by the
 * `patchsentinel` Go CLI (schema version 1). Keeping these in one place lets
 * the renderer and formatter share a single source of truth.
 *
 * This module has zero runtime dependencies.
 */

export type Verdict = "pass" | "warn" | "fail";
export type Severity = "info" | "warn" | "error";

export interface Summary {
  filesChanged: number;
  linesAdded: number;
  linesRemoved: number;
  productionAdded: number;
  testAdded: number;
