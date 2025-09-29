"use client";

import type { ChangeEvent } from "react";
import { Effect, Either, pipe } from "effect";
import { extractFile, type FileReference } from "@/lib/inputs.ts";

export default function UploadButton(
  props: {
    onSelected: (fileRef: FileReference) => void;
  },
): React.JSX.Element {
  const fileSelectedHandler = (e: ChangeEvent<HTMLInputElement>) => {
    const steps = pipe(
      Effect.succeed(e),
      Effect.andThen(extractFile),
      Effect.either,
      Effect.andThen(Either.match({
        onLeft: (error) => Effect.logError(error),
        onRight: (fileRef) => Effect.sync(() => props.onSelected(fileRef)),
      })),
    );

    Effect.runSync(steps);
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
