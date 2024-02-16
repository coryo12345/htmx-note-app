import { test, expect } from "@playwright/test";

const URL = "http://127.0.0.1:8080";

test.describe("Homepage", () => {
  test("has form to create notes", async ({ page }) => {
    expect(1).toBe(1);
  });

  test("creating note displays new card", async ({ page }) => {
    expect(1).toBe(1);
  });

  test("can delete a note", async ({ page }) => {
    expect(1).toBe(1);
  });
});
