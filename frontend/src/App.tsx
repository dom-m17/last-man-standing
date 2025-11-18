import Header from "./components/Header";
import HomePage from "./pages/HomePage";
import { createBrowserRouter, RouterProvider, Outlet } from "react-router-dom";
import SignUpPage from "./pages/SignUpPage";
import NotFoundPage from "./pages/NotFoundPage";

import "./index.css";

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <>
        <Header />
        <Outlet />
      </>
    ),
    errorElement: <NotFoundPage />,
    children: [
      { index: true, element: <HomePage /> },
      { path: "signup", element: <SignUpPage /> },
    ],
  },
]);

export default function App() {
  return <RouterProvider router={router} />;
}
