import * as actions from "./actions.ts";
import { Effect, pipe } from "effect";
import type { Chart } from "./models.ts";

export type FileStatus = {
  key: string;
  files: {
    id: string;
    piece: number;
    content: string;
  };
};

export type FileUploadsStatusResponse = FileStatus[];

export interface FileUploadsService {
  start(id: string, name: string, parts: number): Effect.Effect<void, string>;
  upload(
    id: string,
    piece: number,
    content: string,
  ): Effect.Effect<void, string>;
  status(): Effect.Effect<FileUploadsStatusResponse, string>;
}

export interface ChartsService {
  all(): Effect.Effect<Chart[], string>;
}

const realBase = "http://localhost:5050";

export class FileUploadsServiceImpl implements FileUploadsService {
  start(id: string, name: string, parts: number): Effect.Effect<void> {
    const startMessage =
      `Computer is starting a file transfer of ${parts} parts for the file named ${name}.`;
    const endMessage = `Computer has started a file transfer.`;

    const logStartMessage = Effect.log(">>>" + startMessage + ">>>");

    const startUpload = Effect.promise<void>(
      () => actions.startUpload(realBase, id, name, parts),
    );

    const logEndMessage = Effect.log(">>>" + endMessage + ">>>");

    return pipe(
      logStartMessage,
      Effect.flatMap(() => startUpload),
      Effect.flatMap(() => logEndMessage),
    );
  }

  upload(
    id: string,
    piece: number,
    content: string,
  ): Effect.Effect<void> {
    const startMessage = `Computer is uploading piece number ${piece}!`;

    const endMessage = `Computer has finished uploading piece number ${piece}!`;

    return pipe(
      Effect.log(">>>" + startMessage + ">>>"),
      Effect.flatMap(() =>
        Effect.promise(() => actions.uploadPart(realBase, id, piece, content))
      ),
      Effect.flatMap(() => Effect.log(">>>" + endMessage + ">>>")),
    );
  }

  status(): Effect.Effect<FileUploadsStatusResponse> {
    throw new Error("Method not implemented.");
  }
}
