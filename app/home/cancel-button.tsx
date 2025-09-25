export default function CancelButton(props: {
  onClick: () => void;
}) {
  return (
    <div>
      <button
        type="reset"
        className="
          px-4 py-1
          bg-red-200 rounded-sm
          hover:cursor-pointer hover:bg-red-300
          active:bg-red-400
        "
        onClick={props.onClick}
      >
        Cancel
      </button>
    </div>
  );
}
