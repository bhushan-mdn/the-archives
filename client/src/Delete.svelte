<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import ApiClient from "./ApiClient";

  const dispatch = createEventDispatcher();

  const sendRefresh = () => {
    dispatch("message", {
      event: "refresh",
    });
  };


  export let id: number;
  export let name: string;
  let deleting: boolean = false;

  const deleteFile = () => {
    deleting = true;
    ApiClient.delete(`/file/${id}`)
      .then((result) => {
        console.log("Success:", result);
        sendRefresh();
      })
      .catch((error) => {
        console.log("Error:", error);
      });
    setTimeout(() => {
      deleting = false;
    }, 1500);
  };
</script>

<!-- <main> -->
<!-- <button
    type="button"
    class="btn btn-outline-danger"
    data-bs-toggle="modal"
    data-bs-target="#deleteModal"
  >
    <i class="bi bi-trash-fill" />
  </button> -->

<div
  class="modal fade"
  id="deleteModal"
  tabindex="-1"
  aria-labelledby="deleteModalLabel"
  aria-hidden="true"
>
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="deleteModalLabel">Delete file</h5>
        <button
          type="button"
          class="btn-close"
          data-bs-dismiss="modal"
          aria-label="Close"
        />
      </div>
      <div class="modal-body">
        <p>Are you sure you want to delete {name}?</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal"
          >Close</button
        >
        <button
          type="button"
          class="btn btn-danger"
          on:click={deleteFile}
          data-bs-dismiss="modal"
        >
          Delete
        </button>
      </div>
    </div>
  </div>
</div>
<!-- </main> -->
