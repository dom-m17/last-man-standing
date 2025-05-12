import "./App.css";

function App() {

  return (
    <>
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
              Hello
        </button>
    </>
  );
}

export default App;
