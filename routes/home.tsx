import { define } from "../utils.ts";

import { useSignal } from "@preact/signals";

import { ChartsView } from "../components/ChartsView.tsx";
import { UploadButton } from "../components/UploadButton.tsx";
import { UploadModal } from "../components/UploadModal.tsx";
import type { Chart } from "../models.ts";
import { FileUploadsServiceImpl } from "../services.ts";

type UploadState =
  | { state: "selected"; fileRef: File }
  | { state: "nothing" };

function nothingUploadState(): UploadState {
  return {
    state: "nothing",
  };
}

function selectedUploadState(fileRef: File): UploadState {
  return {
    state: "selected",
    fileRef,
  };
}

export default define.page(function Home() {
  const uploading = useSignal(
    nothingUploadState(),
  );

  const uploadSelected = (fileRef: File) => {
    console.log("selected!");
    uploading.value = selectedUploadState(fileRef);
  };

  const cancelUpload = () => {
    uploading.value = nothingUploadState();
  };

  const exampleCharts: Array<Chart> = [
    {
      id: "1",
      name: "Test Org 1",
    },
  ];

  const service = new FileUploadsServiceImpl();

  const uploadSection = uploading.value.state === "nothing"
    ? <UploadButton onSelected={uploadSelected} />
    : (
      <UploadModal
        fileRef={uploading.value.fileRef}
        onCancel={cancelUpload}
        uploadsService={service}
      />
    );

  return (
    <div className="
        m-2 p-2
        border-2 rounded-md border-purple-200
        flex flex-col gap-2
        bg-slate-100
      ">
      <div>
        <h1 className="text-2xl">
          orgchartsdotorg
        </h1>
      </div>
      <div>
        {uploadSection}
      </div>
      <ChartsView
        charts={exampleCharts}
        onSelected={(id) => {
          console.log("Selected ID: " + id);
        }}
      />
    </div>
  );
});
