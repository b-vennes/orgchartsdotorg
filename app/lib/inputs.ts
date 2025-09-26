import { Effect } from "effect";
import type { ChangeEvent } from "react";
import { empty } from "./extras";

export type FileReference = {
    name: string;
};

export type FileInput = {
    files: { 0: FileReference };
};

export function extractFile(
    event: ChangeEvent<HTMLInputElement>,
): Effect.Effect<FileReference, string> {
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
