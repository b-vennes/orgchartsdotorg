"use client";

import type { ChangeEvent } from "react";
import { empty } from "lib/extras.ts";

function hasFilesKey(obj: unknown): obj is { files: Array<unknown> } {
  const objWithFilesKey = obj as { files: Array<unknown> };

  return !empty(objWithFilesKey) &&
    !empty(objWithFilesKey.files) &&
    typeof objWithFilesKey.files === "object";
}

function hasNameKey(obj: unknown): obj is { name: string } {
  const objWithNameKey = obj as { name: string };

  return !empty(objWithNameKey) &&
    !empty(objWithNameKey.name);
}

export default function UploadButton(
  props: {
    onSelected: (filename: string) => void;
  },
): React.JSX.Element {
  const fileSelectedHandler = (e: ChangeEvent<HTMLInputElement>) => {
    if (!hasFilesKey(e.target)) {
      return;
    }

    const files = e.target.files;

    if (!hasNameKey(files[0])) {
      return;
    }

    const filename = files[0].name;

    props.onSelected(filename);
  };

  return (
    <div>
      <label
        htmlFor="upload-file"
        className="
        px-4 py-1
        rounded-sm
        bg-purple-400
        hover:bg-purple-500 hover:cursor-pointer
        active:bg-purple-600
      "
      >
        Upload Org Chart
      </label>
      <input
        id="upload-file"
        type="file"
        name="Upload"
        style={{ display: "none" }}
        onChange={fileSelectedHandler}
      />
    </div>
  );
}
