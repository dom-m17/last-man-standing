import "./App.css";
import Header from "@/components/Header";
import Matches from "@/components/Matches";

function App() {

  return (
    <>
      <Header />
      <div className="flex flex-col items-center justify-center min-h-svh">
        <Matches/>
      </div>
    </>
  );
}

export default App;
