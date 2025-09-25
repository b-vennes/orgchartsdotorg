import * as actions from "./actions.ts";
import { type Chart } from "./models.ts";

export interface ChartsService {
  upload(): Promise<void>;
  all(): Promise<Array<Chart>>;
}

export class ChartsServiceImpl implements ChartsService {
  async upload(): Promise<void> {
    await actions.upload();
  }

  async all(): Promise<Array<Chart>> {
    return await actions.charts();
  }
}
