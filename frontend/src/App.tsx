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
                  getUser(input: "user_f8d7ba95-92dd-45d8-a56f-5491562966d5") {
                      id
                      username
                      firstName
                      lastName
                      email
                      phoneNumber
                      favouriteTeam
                  }
                }
              `
            })
          })
            .then(res => res.json())
            .then(data => console.log(data))}>
              Get Dom
        </button>
    </>
  );
}

export default App;
