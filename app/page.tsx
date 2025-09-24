import Image from "next/image";

export default function Page() {
  return (
    <div className="m-2 flex flex-col gap-2">
      <h1>Hello, Next.js!</h1>
      <Image
        src="/skyscrapers.jpg"
        alt="scene"
        width={100}
        height={180}
        className="border-2 border-black rounded-md"
      />
    </div>
  );
}
