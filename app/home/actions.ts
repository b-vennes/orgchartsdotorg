"use server";

import type { Chart } from "./models.ts";

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

  const responseBody = JSON.parse(response.body);

  return responseBody.charts;
}
