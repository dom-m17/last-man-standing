import "./App.css";
import Button from "@/components/Button";
import Header from "@/components/Header";
import Matches from "@/components/Matches";

function App() {

  return (
    <>
      <Header />
      <div className="flex flex-col items-center justify-center min-h-svh">
        <Matches/>
        <Button>Click me</Button>
      </div>
    </>
  );
}

export default App;
