<script lang="ts">
  import prettyBytes from "pretty-bytes";
  import moment from "moment";
  import { createEventDispatcher, onMount } from "svelte";

  import Icon from "./Icon.svelte";
  import Download from "./Download.svelte";
  import type { File } from "./File.interface";
  import ApiClient from "./ApiClient";

  export let file: File;
  let renaming = false;

  let deleteFileID: number;
  let deleteFileName: string;

  let renameRef: HTMLInputElement;

  // onMount(() => {
  //   renameRef.focus();
  // });

  const dispatch = createEventDispatcher();

  const sendRefresh = () => {
    dispatch("message", {
      event: "refresh",
    });
  };

  const sendDelete = () => {
    dispatch("message", {
      event: "delete",
      deleteFileID,
      deleteFileName,
    });
  };

  const setDelete = (e: any, id: number, name: string) => {
    console.log("Delete event called", e);
    deleteFileID = id;
    deleteFileName = name;
    sendDelete();
  };

  const renameStart = (e: any) => {
    console.log("renameStart. Event:", e);
    renaming = true;
    setTimeout(() => {
      renameRef.focus();
    }, 250);
  };

  let renameValue: string = file.name;

  const renameFile = () => {
    if (renameValue != file.name && renameValue !== "") {
      console.log("File name changed");
      ApiClient.put(`/file/${file.ID}`, { name: renameValue })
        .then((resp) => {
          console.log(resp);
        })
        .catch((err) => {
          console.error(err);
        });
      renameValue = file.name;
      sendRefresh();
    } else {
      console.log("File name unchanged");
    }
  };

  const renameHandler = (e: any) => {
    if (e.keyCode == 13 || e.keyCode == 27) {
      console.log("renameHandler");
      renaming = false;
      renameFile();
    }
  };

  const onBlur = () => {
    console.log("onBlur");
    renaming = false;
    renameFile();
  };
</script>

<td>
  <div class="d-flex align-items-center justify-content-between">
    <div class="d-flex align-items-center justify-content-start gap-3 ">
      <Icon type={file.type} />
      {#if renaming}
        <input
          type="text"
          placeholder="File Name"
          aria-label="File Name"
          aria-describedby="basic-addon1"
          bind:this={renameRef}
          on:keydown={renameHandler}
          on:blur={onBlur}
          bind:value={renameValue}
        />
      {:else}
        {file.name}
      {/if}
    </div>
    <button
      class="btn actions-button"
      type="button"
      id="actionsMenuButton1"
      data-bs-toggle="dropdown"
      aria-expanded="false"
    >
      <i class="bi bi-three-dots" style="color: grey;" />
    </button>
    <ul class="dropdown-menu" aria-labelledby="actionsMenuButton1">
      <li>
        <button
          class="dropdown-item d-flex gap-2"
          on:click={(e) => renameStart(e)}
        >
          <i class="bi bi-pencil" />
          Rename
        </button>
      </li>
      <li>
        <Download id={file.ID} name={file.name} />
      </li>
      <li>
        <button
          class="dropdown-item d-flex gap-2"
          data-bs-toggle="modal"
          data-bs-target="#deleteModal"
          on:click={(e) => setDelete(e, file.ID, file.name)}
        >
          <i class="bi bi-trash3-fill" />
          Delete
        </button>
      </li>
    </ul>
  </div>
</td>
<!-- <td>{file.type}</td> -->
<td>{prettyBytes(file.size)}</td>
<td>{moment(file.UpdatedAt).fromNow()}</td>

<style>
  .actions-button:hover {
    outline: none;
    background-color: none;
    color: none;
    border-color: rgba(0, 0, 0, 0);
    /* color: black; */
  }

  .actions-button:focus {
    outline: none;
    background-color: none;
    color: none;
    border-color: rgba(0, 0, 0, 0);
    /* color: black; */
  }

  .actions-button:active {
    outline: none;
    background-color: none;
    color: none;
    border-color: rgba(0, 0, 0, 0);
    /* color: black; */
  }
</style>
