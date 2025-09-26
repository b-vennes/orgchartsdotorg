import CancelButton from "./cancel-button.tsx";

export default function UploadModal(
  props: { filename: string; onCancel: () => void },
) {
  return (
    <div className="
      p-4
      flex flex-col
      rounded-md
      bg-stone-100
      gap-2
      ">
      <div>
        <h2 className="text-lg">Uploading <span className="font-mono font-bold">{props.filename}</span></h2>
      </div>
      <div>
        <input
          type="text"
          placeholder="My Org Name"
          className="border-1 rounded-lg px-3 py-1"
        />
      </div>
      <div className="flex flex-row gap-1">
        <button
          type="button"
          className="
            px-2
            py-1
            bg-purple-400
            rounded-sm
            hover:bg-purple-500 hover:cursor-pointer
            active:bg-purple-600"
        >
          Upload
        </button>
        <CancelButton onClick={props.onCancel} />
      </div>
    </div>
  );
}
