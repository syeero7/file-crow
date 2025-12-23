import { registerFile } from "./api";
import { getHTMLElement } from "./utils";

const uploadForm = getHTMLElement("[data-upload-form]");

uploadForm.addEventListener("submit", (e) => {
  e.preventDefault();
  const form = e.currentTarget as HTMLFormElement;
  const formData = new FormData(form);
  const files = formData.getAll("files") as File[];

  files.forEach(async ({ name, size }) => {
    const id = crypto.randomUUID();
    await registerFile({ type: "register", id, name, size });
  });

  form.reset();
});
