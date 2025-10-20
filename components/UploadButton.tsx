import { Effect, Either, pipe } from "effect";
import { extractFile } from "../inputs.ts";

export function UploadButton(
  props: {
    onSelected: (fileRef: File) => void;
  },
) {
  const fileSelectedHandler = (e) => {
    console.log(e);
    const steps = pipe(
      Effect.succeed(e),
      Effect.flatMap(extractFile),
      Effect.either,
      Effect.flatMap(Either.match({
        onLeft: (error) => Effect.logError(error),
        onRight: (fileRef) => Effect.sync(() => props.onSelected(fileRef)),
      })),
    );

    Effect.runSync(steps);
  };

  return (
    <div>
      <button type="submit" onClick={() => console.log("clicked!")}>
        Click me
      </button>
      <input
        type="file"
        onChange={(e) => {
          console.log("File chosen using on-change!");
        }}
        onInput={(e) => {
          console.log("File chosen using on-input!");
        }}
        onSelect={(e) => {
          console.log("File chosen using on-select!");
        }}
      />
    </div>
  );
}
