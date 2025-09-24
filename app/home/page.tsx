import UploadButton from "./upload-button";

export default function Page() {
  return (
    <div
      className="
        m-2 p-2
        border-2 rounded-md border-purple-200
        flex flex-col gap-2
      "
    >
      <div>
        <h1 className="text-2xl">
          orgchartsdotorg
        </h1>
      </div>
      <div>
        <UploadButton/>
      </div>
    </div>
  );
}
