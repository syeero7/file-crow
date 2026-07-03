<script lang="ts">
  import UploadFiles from "@/components/UploadFiles.svelte";
  import DownloadFiles from "@/components/DownloadFiles.svelte";
  import { initializeWebsocket } from "@/lib/websocket.svelte";

  initializeWebsocket();
</script>

<main>
  <h1>File Crow</h1>
  <UploadFiles />
  <DownloadFiles />
</main>

<style>
  h1 {
    max-width: fit-content;
    margin-inline: auto;
  }

  main {
    max-width: 64em;
    margin-inline: auto;
    padding: 0.5em;
  }
</style>
