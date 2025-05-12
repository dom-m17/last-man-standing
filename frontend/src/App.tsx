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
      <button onClick={() => fetch(
          "http://localhost:8080/query", {
            method: "POST",
            headers: {
              "Content-Type": "application/json"
            },
            body: JSON.stringify({
              query: `
                query {
                  hello
                }
              `
            })
          })
            .then(res => res.json())
            .then(data => console.log(data))}>
        </button>
    </>
  );
}

export default App;
