<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import ApiClient from "./ApiClient";

  const dispatch = createEventDispatcher();

  const sendRefresh = () => {
    dispatch("message", {
      event: "refresh",
    });
  };

  let files;
  let fileSizeExceeds = false;
  let progress = 0;
  let showSuccess = false;

  const onFileSelect = (e) => {
    if (e.target.files.length > 0) {
      console.log(files);
      const file = e.target.files[0];
      console.log(progress);
      // if (file["size"] > 15 * 1024 * 1024) {
      //   console.log("File size exceeds 15 MB");
      //   fileSizeExceeds = true;
      // } else {
      //   console.log("File size does not exceed");
      //   fileSizeExceeds = false;
      // }
    }
  };

  const uploadFile = () => {
    if (files[0]) {
      const formData: FormData = new FormData();
      formData.append("file", files[0]);
      ApiClient.post("/blob", formData, {
        onUploadProgress: (progressEvent: any) => {
          if (progressEvent.lengthComputable) {
            console.log(
              "Uploaded " + progressEvent.loaded + " of " + progressEvent.total
            );
            progress = Math.round(
              (progressEvent.loaded / progressEvent.total) * 100
            );
            console.log(`Uploaded ${progress} %`);
            showSuccess = true;
            setTimeout(() => {
              progress = 0;
              showSuccess = false;
              files = [];
            }, 1500);
          }
        },
      })
        .then((result) => {
          console.log("Success:", result);
          sendRefresh();
        })
        // .catch((error) => {
        //   console.error("Error:", error);
        // })
        .catch((error) => {
          console.error("Error:", error);
        });
    }
  };
</script>

<main>
  <div
    class="modal fade"
    id="uploadModal"
    tabindex="-1"
    aria-labelledby="uploadModalLabel"
    aria-hidden="true"
  >
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="uploadModalLabel">Upload file</h5>
          <button
            type="button"
            class="btn-close"
            data-bs-dismiss="modal"
            aria-label="Close"
          />
        </div>
        <div class="modal-body">
          <div class="mb-3">
            <label for="file-upload" class="form-label"
              >Choose file to upload</label
            >
            <input
              class="form-control"
              type="file"
              name="file"
              id="file-upload"
              bind:files
              on:change={(e) => onFileSelect(e)}
            />
          </div>
          {#if fileSizeExceeds}
            <p>File size should be less than 15MB</p>
          {/if}
          {#if showSuccess}
            <p>Successfully uploaded file.</p>
          {/if}
          <div class="progress">
            <div
              class="progress-bar progress-bar-striped progress-bar-animated"
              role="progressbar"
              aria-valuenow={progress}
              aria-valuemin={0}
              aria-valuemax={100}
              style={`width: ${progress}%`}
            />
          </div>
        </div>
        <div class="modal-footer">
          <button
            type="button"
            class="btn btn-secondary"
            data-bs-dismiss="modal">Close</button
          >
          <button
            type="button"
            class="btn btn-primary"
            on:click={uploadFile}
            disabled={fileSizeExceeds}>Upload</button
          >
        </div>
      </div>
    </div>
  </div>
</main>
