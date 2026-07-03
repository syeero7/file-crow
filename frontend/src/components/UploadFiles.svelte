<script lang="ts">
  import { useWebsocket } from "@/lib/websocket.svelte";
  import { registerFile } from "@/lib/api";

  const ws = useWebsocket();

  function register(files: File[]) {
    files.forEach(async (file) => {
      const { name, size } = file;
      if (!name || !size) return;
      const { id } = await registerFile({ type: "register", name, size });
      ws.pendingFiles.set(id, file);
    });
  }
</script>

<section>
  <h2>Upload Files</h2>

  <form
    onsubmit={(e) => {
      e.preventDefault();

      const form = e.currentTarget;
      const formData = new FormData(form);
      const files = formData.getAll("files") as File[];
      if (!files.length) return;

      register(files);
      form.reset();
    }}
  >
    <div>
      <input type="file" name="files" multiple aria-label="select files to upload" />
    </div>
    <button type="submit">Upload</button>
  </form>
</section>

<style>
  form {
    padding: 0.5em;
    border-radius: var(--radius);
    corner-shape: squircle;
    display: grid;
    gap: 1.5em;

    div {
      --border: 2px solid hsla(from var(--foreground) h s l / 0.2);

      border: var(--border);
      box-shadow: 0 2px 2px hsla(from var(--foreground) h s l / 0.1);
      border-radius: var(--radius);
      corner-shape: squircle;
      overflow: hidden;

      input {
        width: 100%;
        cursor: pointer;
        color: var(--foreground);
        transition: background-color 0.25s ease-in-out;

        &::file-selector-button {
          border: none;
          font: inherit;
          cursor: pointer;
          padding: 0.75em 0.9em;
          border-right: var(--border);
          background-color: transparent;
          margin-right: 1em;
          color: var(--foreground);
        }

        &:hover,
        &:focus-visible {
          background-color: var(--background-a10);
        }
      }
    }

    button {
      padding: 0.75em;
      border-radius: var(--radius);
      max-width: 200px;
      width: 100%;
      margin-inline: auto;
      font-weight: 600;
      border: 2px solid transparent;
      transition-property: background-color, border-color, filter;
      transition-duration: 0.25s;
      color: var(--background);
      background: var(--primary);

      &:hover,
      &:focus-visible {
        background: var(--primary-a10);
        border-color: var(--primary-a10);
        filter: drop-shadow(0 0 0.25em var(--primary-a10));
      }
    }
  }
</style>
