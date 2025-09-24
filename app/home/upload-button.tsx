"use client";

import { upload } from "./actions";

export default function UploadButton() {
  return (
    <button
      onClick={upload}
      className="
        px-4 py-1
        rounded-sm
        bg-purple-400
        hover:bg-purple-500 hover:cursor-pointer
        active:bg-purple-600
      "
    >Upload Chart</button>
  )
}
