<script lang="ts">
  import UploadFiles from "@/components/UploadFiles.svelte";
  import { initializeWebsocket } from "@/lib/websocket.svelte";

  initializeWebsocket();
</script>

<main>
  <UploadFiles />
</main>

<style>
</style>
