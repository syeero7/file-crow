<script lang="ts">
  import { useWebsocket } from "@/lib/websocket.svelte";
  import { getDownloadURL } from "@/lib/api";

  const ws = useWebsocket();
</script>

<section>
  <h2>Download Files</h2>
  <ul>
    {#each ws.files as entry (entry[0])}
      {@const [id, file] = entry}

      <li>
        <div class="item">
          <label for={id}>{file.name}</label>
          <span class="size">{file.size}</span>
          <span class="status" data-status={file.status}>{file.status}</span>

          {#if file.status !== "pending"}
            {@const { percentage } = file.progress}

            <progress {id} max={100} value={percentage}>{percentage} &#37;</progress>
          {:else}
            <progress {id}></progress>
          {/if}
        </div>

        {#snippet svg()}
          <svg viewBox="0 -960 960 960" xmlns="http://www.w3.org/2000/svg">
            <path
              d="M480-315.33 284.67-510.67l47.33-48L446.67-444v-356h66.66v356L628-558.67l47.33 48L480-315.33ZM226.67-160q-27 0-46.84-19.83Q160-199.67 160-226.67V-362h66.67v135.33h506.66V-362H800v135.33q0 27-19.83 46.84Q760.33-160 733.33-160H226.67Z"
            />
          </svg>
        {/snippet}

        {#if file.status !== "pending"}
          <a href={null}>
            {@render svg()}
          </a>
        {:else}
          <a download href={getDownloadURL(id)} title="download file">
            {@render svg()}
          </a>
        {/if}
      </li>
    {/each}
  </ul>
</section>

<style>
  ul {
    list-style: none;
    padding: 0 0.25em 0.5em;
  }

  li {
    display: flex;
    gap: 0.5em;
    border-radius: var(--radius);
    corner-shape: squircle;
    background-color: var(--background-a10);
    padding: 0.25em 0.5em;

    &:not(:last-of-type) {
      margin-bottom: 0.5em;
    }
  }

  .item {
    flex: 1;
    font-size: 0.9em;
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    grid-template-rows: repeat(3, 1fr);
    pointer-events: none;
    gap: 0.2em;

    label {
      grid-area: 1/1/1/5;
      font-weight: 500;
      white-space: nowrap;
      overflow-x: auto;
      margin-left: 0.25em;
      scrollbar-width: thin;
      scrollbar-color: hsla(from var(--text) h s l / 0.2) var(--primary);
    }

    .size {
      grid-area: 2/1/2/2;
      font-size: 0.9em;
      margin-top: 0.2em;
      margin-left: 0.25em;
    }

    --status-color: var(--info);
    --text-transform: unset;

    .status {
      grid-area: 2/2/2/3;
      font-size: 0.9em;
      margin-top: 0.2em;
      margin-left: 1em;
      font-weight: 500;
      color: var(--status-color);
      text-transform: var(--text-transform);
      letter-spacing: 1px;
    }

    progress {
      grid-area: 3/1/3/5;
      width: calc(100% - 0.25em);
      margin-inline: 0.25em;
      height: 10px;
      margin-top: 0.2em;

      @media (prefers-color-scheme: dark) {
        accent-color: var(--status-color);
      }
    }

    &:has([data-status="pending"]) {
      --status-color: var(--warning);
      --text-transform: capitalize;
    }

    &:has([data-status="failed"]) {
      --status-color: var(--danger);
      --text-transform: capitalize;
    }

    &:has([data-status="done"]) {
      --status-color: var(--success);
      --text-transform: capitalize;
    }
  }

  a {
    --dlbtn-color: var(--primary-a30);

    align-self: center;
    width: 3em;
    height: 3em;
    display: flex;
    justify-content: center;
    align-items: center;
    border: 2px solid var(--dlbtn-color);
    border-radius: var(--radius);
    transition-property: border-color, filter;
    transition-duration: 0.25s, 0.1s;

    svg {
      width: 2.2em;
      height: 2.2em;
      pointer-events: none;
      fill: var(--dlbtn-color);
    }

    &[href]:hover,
    &[href]:focus-visible {
      --dlbtn-color: var(--accent-a10);
      filter: drop-shadow(0 0 0.5em var(--dlbtn-color));
    }

    &:not([href]) {
      cursor: not-allowed;
      filter: none;
      opacity: 0.6;
    }
  }
</style>
