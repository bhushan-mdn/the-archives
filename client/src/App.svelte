<script lang="ts">
  import FileRow from "./FileRow.svelte";
  import { onMount } from "svelte";
  import { apiData, fileList } from "./store";

  import Upload from "./Upload.svelte";
  import Delete from "./Delete.svelte";
  import ApiClient from "./ApiClient";
  import type { File } from "./File.interface";

  let favicon = "/favicon.png";
  // export const name;
  let file;
  let fileCount = 0;
  let refreshing = false;
  let selectAll = false;
  let selectPartial = false;
  let selected = false;
  let selectedFileCount = 0;
  let selectedFiles = [];
  let selectedSize = 0;

  let deleteFileID: number;
  let deleteFileName: string;

  onMount(() => {
    getFiles();
  });

  const handleMessage = (e: any) => {
    console.log(e.detail);

    switch (e.detail.event) {
      case "refresh":
        getFiles();
        break;
      case "delete":
        deleteFileID = e.detail.deleteFileID;
        deleteFileName = e.detail.deleteFileName;
        break;
      default:
        break;
    }
  };

  const getFiles = () => {
    console.log("refresh");
    refreshing = true;
    ApiClient.get("/files")
      .then((response: any) => response.data)
      .then((data) => {
        console.log(data);
        if (data["files"]) {
          apiData.set(data["files"]);
        } else {
          apiData.set([]);
        }
        fileCount = $fileList.length;
      })
      .catch((error) => {
        console.log(error);
        apiData.set([]);
        // return [];
      });
    setTimeout(() => {
      refreshing = false;
    }, 1000);
  };

  const onAllCheck = (e: Event) => {
    // selectAll = !selectAll;
    // selectedFiles = $fileList.map((li) => li.ID);
    $fileList.map((f) => (f.selected = selectAll));
    console.log($fileList.map((li) => li.selected));
    // if (selectAll) {
    //   $fileList.map((file: File) => {
    //     file.selected = true;
    //   });
    //   console.log($fileList);
    //   selectChange();
    // } else {
    //   $fileList.map((file: File) => {
    //     file.selected = false;
    //   });
    //   selectChange();
    // }
  };

  const selectChange = () => {
    console.log("selectChange");
    $fileList.every((file: File, index: number) => file.selected)
      ? (selectAll = true)
      : (selectAll = false);
    console.log("selectAll:", selectAll);
    console.log("selectPartial:", selectPartial);
    // const reduction = $fileList.map((file: File) => file.selected);
    const reduction = selectedFiles;
    console.log("fileList:", reduction);
    if (reduction.includes(true)) {
      selectPartial = true;
      // selectAll = false;
    } else {
      selectPartial = false;
      // selectAll = false;
    }
  };

  const onFileSelect = (e: Event, id: number) => {
    // console.log("file.selected:", file.selected);
    console.log("onFileSelect", e);
    if ($fileList.filter(f => f.ID == id)[0].selected == true) {
      selectedSize += ($fileList.filter(f => f.ID == id)[0].size)
    } else {
      selectedSize -= ($fileList.filter(f => f.ID == id)[0].size)
    }
    console.log(selectedSize);
    selectChange();
  };
</script>

<svelte:head>
  <link rel="icon" type="image/png" href={favicon} />
</svelte:head>

<main>
  <h1 class="title">The Archives</h1>
  <div>
    <button
    type="button"
    class="btn btn-primary"
    data-bs-toggle="modal"
    data-bs-target="#uploadModal"
  >
    Upload
  </button>
  {#if refreshing}
    <button class="btn btn-secondary" type="button" disabled>
      <span
        class="spinner-border spinner-border-sm"
        role="status"
        aria-hidden="true"
      />
      Refreshing...
    </button>
  {:else}
    <button type="button" class="btn btn-secondary" on:click={getFiles}>
      Refresh
    </button>
  {/if}
  </div>
  <div class="container m-auto">
    <table class="table table-striped align-middle col-md-6 text-start">
      <thead>
        <tr>
          <th scope="col">
            <input
              type="checkbox"
              name="selectAll"
              id="selectAll"
              bind:checked={selectAll}
              bind:indeterminate={selectPartial}
              on:change={onAllCheck}
            />
          </th>
          <th scope="col">
            {#if selected}
              {`${selectedFileCount} files`} <i class="bi bi-three-dots" />
              {"Actions"}
            {:else}
              {"Name"}
            {/if}
          </th>
          <th scope="col">
            {"Size"}
          </th>
          <th scope="col">
            {"Modified"}
          </th>
        </tr>
      </thead>
      {#if refreshing}
        <tbody>
          <tr>
            <td colspan="4" class="text-center">
              <div class="spinner-border m-5" role="status">
                <span class="visually-hidden">Loading...</span>
              </div>
            </td>
          </tr>
        </tbody>
      {:else if fileCount === 0}
        <tbody>
          <tr>
            <td colspan="4" class="text-center">
              <p class="m-5">No files uploaded.</p>
            </td>
          </tr>
        </tbody>
      {:else}
        <tbody>
          {#each $fileList as file (file.ID)}
            <tr class="align-center">
              <th scope="row">
                <input
                  type="checkbox"
                  name="selectFile"
                  id="selectFile"
                  bind:checked={file.selected}
                  on:change={(e) => onFileSelect(e, file.ID)}
                />
              </th>
              <FileRow {file} on:message={handleMessage} />
            </tr>
          {/each}
        </tbody>
      {/if}
    </table>
  </div>

  <Delete id={deleteFileID} name={deleteFileName} on:message={handleMessage} />
  <Upload on:message={handleMessage} />
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  .title {
    /* color: #ff3e00; */
    color: purple;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
