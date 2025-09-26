"use client";

import { useState, type JSX } from "react";

import ChartsView from "./charts-view.tsx";
import UploadButton from "./upload-button.tsx";
import UploadModal from "./upload-modal.tsx";
import type { Chart } from "lib/models.ts";

type UploadState = { state: "nothing" | "selected"; filename: string };

function nothingUploadState(): UploadState {
  return {
    state: "nothing",
    filename: "",
  };
}

function selectedUploadState(filename: string): UploadState {
  return {
    state: "selected",
    filename: filename,
  };
}

export default function Page() {
  const [uploadingState, setUploadingState] = useState(
    nothingUploadState(),
  );

  const uploadSelectedHandler = (name: string) =>
    setUploadingState(selectedUploadState(name));

  const cancelHandler = () => {
    setUploadingState(nothingUploadState());
  };

  const exampleCharts: Array<Chart> = [
    {
      id: "1",
      name: "Test Org 1",
    },
  ];

  const uploadSection: JSX.Element = uploadingState.state === "nothing" ?
    <UploadButton onSelected={uploadSelectedHandler} /> :
    <UploadModal
      filename={uploadingState.filename}
      onCancel={cancelHandler}
    />;

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
