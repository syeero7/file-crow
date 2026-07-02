<script lang="ts">
  import { useWebsocket } from "@/lib/websocket.svelte";
  import { getDownloadURL } from "@/lib/api";

  const ws = useWebsocket();
</script>

<section>
  <h2>Download Files</h2>
  <ul>
    {#each ws.files as entry}
      {const [id, file] = entry}

      <li>
        <label for={id}>{file.name}</label>
        <span>{file.size}</span>
        <span>{file.status}</span>

        {#snippet svg()}
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 -960 960 960">
            <path
              d="M480-315.33 284.67-510.67l47.33-48L446.67-444v-356h66.66v356L628-558.67l47.33 48L480-315.33ZM226.67-160q-27 0-46.84-19.83Q160-199.67 160-226.67V-362h66.67v135.33h506.66V-362H800v135.33q0 27-19.83 46.84Q760.33-160 733.33-160H226.67Z"
            />
          </svg>
        {/snippet}

        {#if file.status !== "pending"}
          {const { percentage } = file.progress}

          <progress {id} value={percentage}>"{percentage} %"</progress>
          <a href={null}>
            {@render svg()}
          </a>
        {:else}
          <progress {id}></progress>
          <a download title="download file" href={getDownloadURL(id)}>
            {@render svg()}
          </a>
        {/if}
      </li>
    {/each}
  </ul>
</section>

<style>
</style>
