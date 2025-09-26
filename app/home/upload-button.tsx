"use client";

import type { ChangeEvent } from "react";
import { extractFile } from "lib/inputs.ts";
import { Effect, Either, pipe } from "effect";
import { empty } from "lib/extras";

export default function UploadButton(
  props: {
    onSelected: (filename: string) => void;
  },
): React.JSX.Element {
  const fileSelectedHandler = (e: ChangeEvent<HTMLInputElement>) => {
    const filenameEffect = pipe(
      Effect.succeed(e),
      Effect.flatMap(extractFile),
      Effect.map(file => file.name),
      Effect.either
    )

    Either.match(
      Effect.runSync(filenameEffect),
      {
        onLeft: (error) => console.error(error),
        onRight: (filename) => props.onSelected(filename),
      }
    );
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
