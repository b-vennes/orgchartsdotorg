import { empty } from "./extras.ts";

export interface Chart {
  id: string;
  name: string;
}

export function isChart(chart: unknown): chart is Chart {
  const knownChart = chart as Chart;
  return !empty(knownChart?.id) && !empty(knownChart?.name);
}

export function isCharts(charts: unknown): charts is Array<Chart> {
  return Array.isArray(charts) &&
    charts.every(isChart);
}
