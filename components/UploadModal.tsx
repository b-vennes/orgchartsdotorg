import { useSignal } from "@preact/signals";

import { Cause, Effect, Either, Exit, pipe } from "effect";
import { CancelButton } from "./CancelButton.tsx";
import { extractText } from "../inputs.ts";
import { type FileUploadsService } from "../services.ts";

const chunkSize = 2;

function onUploadClicked(
  service: FileUploadsService,
  fileRef: File,
  name: string,
): void {
  const steps = pipe(
    Effect.log("Uploading file."),
    Effect.flatMap(() => Effect.sync(() => crypto.randomUUID())),
    Effect.flatMap((id) =>
      pipe(
        Effect.succeed(Math.ceil(fileRef.size / chunkSize)),
        Effect.flatMap((parts) =>
          pipe(
            service.start(id, name, parts),
            Effect.as(parts),
          )
        ),
        Effect.map((parts) => {
          return { parts, id };
        }),
      )
    ),
    Effect.flatMap(({ parts, id }) =>
      pipe(
        Effect.promise(() => fileRef.bytes()),
        Effect.map((text) => {
          return {
            parts,
            id,
            text,
          };
        }),
      )
    ),
    Effect.flatMap(({ parts: partsCount, id, text }) => {
      let parts: { piece: number; content: string }[] = [];
      let currentText = text;

      for (let i = 0; i < partsCount; i++) {
        if (i === partsCount - 1) {
          parts = parts.concat([
            {
              piece: i,
              content: new TextDecoder().decode(currentText),
            },
          ]);
        } else {
          parts = parts.concat([
            {
              piece: i,
              content: new TextDecoder().decode(
                currentText.slice(0, chunkSize),
              ),
            },
          ]);
          currentText = currentText.slice(chunkSize);
        }
      }

      return pipe(
        Effect.all(
          parts.map((part) =>
            pipe(
              Effect.log(`>>>Computer is uploading piece ${part.piece}!>>>`),
              Effect.flatMap(() =>
                service.upload(id, part.piece, part.content)
              ),
            )
          ),
        ),
        Effect.as(id),
      );
    }),
    Effect.either,
    Effect.tap(
      Either.match({
        onLeft: () => Effect.void,
        onRight: (id) =>
          Effect.log(
            ">>> Computer done uploading file transfer //" + id + "//!>>>",
          ),
      }),
    ),
    Effect.flatMap(Either.match({
      onLeft: (error) => Effect.logError("Failed to upload file. " + error),
      onRight: () => Effect.void,
    })),
  );

  Effect.runCallback(
    steps,
    {
      onExit: Exit.match({
        onFailure: (cause) => console.log(Cause.pretty(cause)),
        onSuccess: () => {},
      }),
    },
  );
}

function onOrgNameChange(
  event: unknown,
  setOrgName: (name: string) => void,
): void {
  const getText = extractText(event);

  const steps = pipe(
    getText,
    Effect.either,
    Effect.andThen(Either.match({
      onLeft: (error) => Effect.logError(error),
      onRight: (value) => Effect.sync(() => setOrgName(value)),
    })),
  );

  Effect.runSync(steps);
}

export function UploadModal(
  props: {
    fileRef: File;
    onCancel: () => void;
    uploadsService: FileUploadsService;
  },
) {
  const orgName = useSignal("");

  return (
    <div className="
      p-4
      flex flex-col
      rounded-md
      bg-stone-100
      gap-2
      ">
      <div>
        <h2 className="text-lg">
          Uploading{" "}
          <span className="font-mono font-bold">{props.fileRef.name}</span>
        </h2>
      </div>
      <div>
        <input
          type="text"
          placeholder="My Org Name"
          className="border-1 rounded-lg px-3 py-1"
          onChange={(e) => onOrgNameChange(e, (name) => orgName.value = name)}
        />
      </div>
      <div className="flex flex-row gap-1">
        <button
          type="button"
          className="
            px-2
            py-1
            bg-purple-400
            rounded-sm
            hover:bg-purple-500 hover:cursor-pointer
            active:bg-purple-600"
          onClick={() =>
            onUploadClicked(
              props.uploadsService,
              props.fileRef,
              orgName.value,
            )}
        >
          Upload
        </button>
        <CancelButton onClick={props.onCancel} />
      </div>
    </div>
  );
}
