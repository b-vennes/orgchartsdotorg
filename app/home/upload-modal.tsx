import { type ChangeEvent, useState } from "react";
import { Effect, Either, pipe } from "effect";
import CancelButton from "@/home/cancel-button.tsx";
import { extractText, type FileReference } from "@/lib/inputs.ts";
import { type FileUploadsService } from "@/lib/services.ts";

function onUploadClicked(
  service: FileUploadsService,
  fileRef: FileReference,
  name: string,
): void {
  const steps = pipe(
    Effect.sync(() => crypto.randomUUID()),
    Effect.andThen((id) =>
      pipe(
        Effect.succeed(Math.ceil(fileRef.size / 1000)),
        Effect.andThen((parts) => service.start(id, name, parts)),
      )
    ),
  );

  Effect.runFork(steps);
}

function onOrgNameChange(
  event: ChangeEvent<HTMLInputElement>,
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

export default function UploadModal(
  props: {
    fileRef: FileReference;
    onCancel: () => void;
    uploadsService: FileUploadsService;
  },
) {
  const [orgName, setOrgName] = useState("");

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
          onChange={(e) => onOrgNameChange(e, setOrgName)}
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
              orgName,
            )}
        >
          Upload
        </button>
        <CancelButton onClick={props.onCancel} />
      </div>
    </div>
  );
}
