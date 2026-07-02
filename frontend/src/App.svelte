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
</style>
