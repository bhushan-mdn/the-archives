<script lang="ts">
  export let type: string;

  const iconClasses: { [key: string]: string } = {
    // Media
    image: "bi-file-earmark-image",
    audio: "bi-file-earmark-music",
    video: "bi-file-earmark-play",
    // Documents
    "application/pdf": "bi-file-earmark-pdf",
    "application/msword": "bi-file-earmark-word",
    "application/vnd.ms-word": "bi-file-earmark-word",
    "application/vnd.oasis.opendocument.text": "bi-file-earmark-word",
    "application/vnd.openxmlformats-officedocument.wordprocessingml":
      "bi-file-earmark-word",
    "application/vnd.ms-excel": "bi-file-earmark-excel",
    "application/vnd.openxmlformats-officedocument.spreadsheetml":
      "bi-file-earmark-excel",
    "application/vnd.oasis.opendocument.spreadsheet": "bi-file-earmark-excel",
    "application/vnd.ms-powerpoint": "bi-file-earmark-ppt",
    "application/vnd.openxmlformats-officedocument.presentationml":
      "bi-file-earmark-ppt",
    "application/vnd.oasis.opendocument.presentation": "bi-file-earmark-ppt",
    "text/plain": "bi-file-earmark-text",
    "text/html": "bi-filetype-html",
    "application/json": "bi-filetype-json",
    "application/x-yaml": "bi-filetype-yml",
    // Archives
    "application/gzip": "bi-file-earmark-zip",
    "application/zip": "bi-file-earmark-zip",
  };

  const iconFromMime = (mimeType: string) => {

    const matches = Object.entries(iconClasses).find(([k]) =>
      mimeType.startsWith(k)
    );

    if (matches === undefined) {
      return "bi-file-earmark";
    } else {
      return matches[1];
    }
  };
</script>

<main>
  <i class={`bi ${iconFromMime(type)}`} />
</main>
