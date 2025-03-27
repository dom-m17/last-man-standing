export default function Button({
  ...props
}: React.ComponentProps<"button">) {

  return (
    <button>
      {props.children}
    </button>
  )
}

