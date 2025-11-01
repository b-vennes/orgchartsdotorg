"use server";

import { type Chart, isCharts } from "./models.ts";
import * as applib from "@app-lib";

export async function startUpload(
  base: string,
  id: string,
  name: string,
  parts: number,
): Promise<void> {
  await applib.initializeUpload(base, id, name, parts);
}

export async function uploadPart(
  base: string,
  id: string,
  piece: number,
  content: string,
): Promise<void> {
  await fetch(
    base + "/upload-part",
    {
      method: "POST",
      body: JSON.stringify({
        id,
        piece,
        content,
      }),
    },
  );
}

export async function checkUploads(
  base: string,
): Promise<unknown> {
  const response = await fetch(
    base + "/upload-status",
    {
      method: "GET",
    },
  );

  return await response.json();
}

export async function charts(base: string): Promise<Array<Chart>> {
  const response = await fetch(
    base + "/charts",
    {
      method: "POST",
    },
  );

  const responseJson = await response.json();

  if (isCharts(responseJson)) {
    return responseJson;
  } else {
    throw new Error(
      `Expected response to be charts array.  Got '${responseJson}'.`,
    );
  }
}
