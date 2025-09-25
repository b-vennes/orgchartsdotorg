"use client";

export default function UploadButton(
  props: {
    onSelected: (filename: string) => void;
  },
): React.JSX.Element {
  const fileSelectedHandler = (e: any) => {
    const filename = e.target.files[0].name;

    props.onSelected(filename);
  };

  return (
    <div>
      <label
        htmlFor="upload-file"
        className="
        px-4 py-1
        rounded-sm
        bg-purple-400
        hover:bg-purple-500 hover:cursor-pointer
        active:bg-purple-600
      "
      >
        Upload Org Chart
      </label>
      <input
        id="upload-file"
        type="file"
        name="Upload"
        style={{ display: "none" }}
        onChange={fileSelectedHandler}
      />
    </div>
  );
}
