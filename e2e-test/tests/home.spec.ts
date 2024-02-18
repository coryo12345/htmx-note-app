import { test, expect } from "@playwright/test";
import { HomePage } from "./home.page";

const URL = "http://127.0.0.1:8080";

test.describe("Homepage", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto(URL);
    await page.waitForSelector("main#note-app");
    await page.waitForSelector("#note-form")
  });

  test("has form to create notes", async ({ page }) => {
    const homepage = new HomePage(page);
    await expect(homepage.noteForm.el).toBeVisible();
    await expect(homepage.noteForm.valueInput).toBeVisible();
    await expect(homepage.noteForm.authorInput).toBeVisible();
  });

  test("creating note displays new card which can be deleted", async ({ page }, a) => {
    const homepage = new HomePage(page);
    let notes = await homepage.getAllNotes();
    const originalLength = notes.length;

    const noteValue = `${a.project.name} ${a.testId} ${a.title}`;

    // add a new note
    await homepage.noteForm.authorInput.fill("automated test user");
    await homepage.noteForm.valueInput.fill(noteValue);
    await homepage.noteForm.submitBtn.click();

    // wait for network request & new dom elements to be inserted
    await page.waitForTimeout(200);

    // make sure new note exists
    notes = await homepage.getAllNotes();
    await expect(notes.length).toBe(originalLength + 1);
    const addedNote = await page.locator('.note-item').filter({hasText: noteValue});
    await expect(addedNote).toBeVisible();

    // delete note
    const deleteBtn = addedNote.locator('button[type="submit"]');
    await deleteBtn.click();

    // verify it is deleted
    await expect(addedNote).not.toBeVisible();

    // verify it is deleted after refresh
    await page.reload();
    await page.waitForSelector("main#note-app");
    await expect(addedNote).not.toBeVisible();
  });
});
