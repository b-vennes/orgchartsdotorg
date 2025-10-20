"use client";

import { type JSX, useState } from "react";

import ChartsView from "@/home/charts-view.tsx";
import UploadButton from "@/home/upload-button.tsx";
import UploadModal from "@/home/upload-modal.tsx";
import type { Chart } from "@/lib/models.ts";
import { FileUploadsServiceImpl } from "@/lib/services.ts";

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

export default function Page() {
  const [uploadingState, setUploadingState] = useState(
    nothingUploadState(),
  );

  const uploadSelectedHandler = (fileRef: File) =>
    setUploadingState(selectedUploadState(fileRef));

  const cancelHandler = () => {
    setUploadingState(nothingUploadState());
  };

  const exampleCharts: Array<Chart> = [
    {
      id: "1",
      name: "Test Org 1",
    },
  ];

  const service = new FileUploadsServiceImpl();

  const uploadSection: JSX.Element = uploadingState.state === "nothing"
    ? <UploadButton onSelected={uploadSelectedHandler} />
    : (
      <UploadModal
        fileRef={uploadingState.fileRef}
        onCancel={cancelHandler}
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
      <ChartsView charts={exampleCharts} />
    </div>
  );
}
