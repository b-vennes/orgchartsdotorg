declare module "@applib" {
  export function sampleTask(): Promise<void>;
  export function initializeUpload(
    urlBase: string,
    id: string,
    name: string,
    parts: number,
  ): Promise<void>;
}
