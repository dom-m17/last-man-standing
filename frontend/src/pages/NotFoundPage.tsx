import { Link } from "react-router-dom";

export default function NotFoundPage() {
  return (
    <div className="flex flex-col gap-4">
      404 Not Found
      <Link to="/">Home from Link</Link>
    </div>
  );
}
