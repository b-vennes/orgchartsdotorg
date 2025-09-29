import * as actions from "./actions.ts";
import { Effect, pipe } from "effect";

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
      Effect.andThen(startUpload),
      Effect.andThen(logEndMessage),
    );
  }

  upload(
    id: string,
    piece: number,
    content: string,
  ): Effect.Effect<void> {
    throw new Error("Method not implemented.");
  }
  status(): Effect.Effect<FileUploadsStatusResponse> {
    throw new Error("Method not implemented.");
  }
}
