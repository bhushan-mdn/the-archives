<script lang="ts">
  import { saveAs } from "file-saver";
  import ApiClient from "./ApiClient";
  import config from "./config";

  export let id: number;
  export let name: string;
  let downloading = false;

  export const downloadFile = async () => {
    console.log(id, name);
    try {
      // ApiClient
      //   .get(`${config.API_URL}/blob/${id}`, {
      //     responseType: 'blob',
      //   })
      //   .then((blob: any) => {
      //     saveAs(blob, name);
      //   })
      //   .catch((error) => {
      //     console.error(error);
      //   })
      saveAs(`${config.API_URL}/blob/${id}`, name);
      downloading = true;
    } catch (error) {
      console.error(error);
    }
    setTimeout(() => {
      downloading = false;
    }, 1500);
    // .then((response: any) => response.blob())
    // .then(response => {
    // });
  };
</script>

<main>
  <!-- <button
    type="button"
    class="btn btn-outline-primary"
    on:click={downloadFile}
  >
    {#if downloading}
      <span
        class="spinner-border spinner-border-sm"
        role="status"
        aria-hidden="true"
      />
      <span class="visually-hidden">Loading...</span>
    {:else}
      <i class="bi bi-download" />
    {/if}
  </button> -->
  <button class="dropdown-item d-flex gap-2" on:click={downloadFile}>
    <i class="bi bi-download" />
    Download
  </button>
</main>
