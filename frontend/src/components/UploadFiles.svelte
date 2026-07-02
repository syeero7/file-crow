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
  <h1>Upload Files</h1>

  <form
    onsubmit={(e) => {
      const form = e.currentTarget;
      const formData = new FormData(form);
      const files = formData.getAll("files") as File[];
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
</style>
