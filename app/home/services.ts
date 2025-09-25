import { upload } from "./actions.ts";

export interface UploadService {
  upload(): Promise<void>;
}

export class UploadServiceImpl implements UploadService {
  async upload(): Promise<void> {
    await upload();
  }
}
