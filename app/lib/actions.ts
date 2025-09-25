"use server";

import { type Chart, isCharts } from "./models.ts";

export async function upload(): Promise<void> {
  await fetch(
    "http://localhost:5050/upload",
    {
      method: "POST",
    },
  );
}

export async function charts(): Promise<Array<Chart>> {
  const response = await fetch(
    "http://localhost:5050/charts",
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
