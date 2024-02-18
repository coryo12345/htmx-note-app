import { Locator, Page } from "playwright/test";

export class HomePage {
  private page: Page;
  noteForm: {
    el: Locator;
    valueInput: Locator;
    authorInput: Locator;
    submitBtn: Locator;
  };
  notes: Locator;

  constructor(page: Page) {
    this.page = page;
    this.noteForm = {
      el: page.locator("#note-form"),
      valueInput: page.locator("#value"),
      authorInput: page.locator("#author"),
      submitBtn: page.locator('#note-form button[type="submit"]'),
    };
    this.notes = page.locator("#note-container .note-item");
  }

  async getAllNotes(): Promise<Locator[]> {
    return this.notes.all();
  }
}
