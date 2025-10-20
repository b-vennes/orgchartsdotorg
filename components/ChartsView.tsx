import { useSignal } from "@preact/signals";

import { Charts } from "./Charts.tsx";
import type { Chart } from "../models.ts";

type Selected = { id: string | null };

function noneSelected(): Selected {
  return { id: null };
}

function selectedId(id: string): Selected {
  return { id };
}

export function ChartsView(props: {
  charts: Array<Chart>;
}) {
  const selected = useSignal(noneSelected());

  const selectStateHandler = (id: string) => {
    selected.value = selectedId(id);
  };

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
        {selected.value.id === null
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
