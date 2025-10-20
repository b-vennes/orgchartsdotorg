import { Effect } from "effect";
import { empty } from "./extras.ts";

export type FileInput = {
  files: FileList;
};

export function extractFile(
  event: unknown,
): Effect.Effect<File, string> {
  if (empty(event.currentTarget)) {
    return Effect.fail(
      "Expected event object to have a non-empty 'currentTarget' field.",
    );
  }

  const asFileInputEvent = event.target as unknown as FileInput;

  if (empty(asFileInputEvent.files)) {
    const elementStr = JSON.stringify(asFileInputEvent);
    return Effect.fail(
      `Expected event target to be a file input element. Got '${elementStr}'.`,
    );
  }

  return Effect.succeed(asFileInputEvent.files[0]);
}

export type TextInput = {
  value: string;
};

export function extractText(
  event: unknown,
): Effect.Effect<string, string> {
  if (empty(event.currentTarget)) {
    return Effect.fail(
      "Expected event object to have a non-empty 'currentTarget' field.",
    );
  }

  const asTextInput = event.currentTarget as TextInput;

  if (empty(asTextInput.value)) {
    return Effect.fail(
      "Expected currentTarget object to have a non-empty 'value' field.",
    );
  }

  return Effect.succeed(asTextInput.value);
}
