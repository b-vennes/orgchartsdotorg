"use server";

export async function upload() {
  "use server";

  console.log("Uploading!");

  await fetch("http://localhost:5050/upload", {
    method: "POST"
  });
}
