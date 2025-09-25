import type { Chart } from "./models.ts";

export default function Charts(props: {
  charts: Array<Chart>;
  onSelected: (id: string) => void;
}) {
  function chartClickedHandler(id: string): () => void {
    return () => props.onSelected(id);
  }

  return (
    <div className="flex flex-col gap-1">
      <h2 className="text-lg">Your Charts</h2>
      {props.charts.map((c) => (
        <div
          key={c.id}
          className="
            border-1 rounded-md
            hover:cursor-pointer
          "
          onClick={chartClickedHandler(c.id)}
        >
          {c.name}
        </div>
      ))}
    </div>
  );
}
