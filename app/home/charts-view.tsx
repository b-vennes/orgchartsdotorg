"use client";

import { useState } from "react";

import Charts from "./charts.tsx";
import type { Chart } from "lib/models.ts";

type Selected = { id: string | null };

function noneSelected(): Selected {
  return { id: null };
}

function selectedId(id: string): Selected {
  return { id };
}

export default function ChartsView(props: {
  charts: Array<Chart>;
}) {
  const [selectedState, setSelectedState] = useState(noneSelected);

  const selectStateHandler = (id: string) => setSelectedState(selectedId(id));

  return (
    <div className="
      grid grid-cols-3 gap-2
      bg-white border-1 rounded-sm
      p-2
    ">
      <div>
        <Charts charts={props.charts} onSelected={selectStateHandler} />
      </div>
      <div className="
            border-1 rounded-lg
            col-span-2
            flex items-center justify-center
          ">
        {selectedState.id === null
          ? (
            <p className="text-center">
              Select an uploaded org chart to view current state.
            </p>
          )
          : (
            <div>
              <p>More details needed!</p>
            </div>
          )}
      </div>
    </div>
  );
}
