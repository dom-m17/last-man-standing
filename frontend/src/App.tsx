import "./App.css";
import { Button } from "@/components/ui/Button";
import Header from "@/components/ui/Header";

interface fixture {
  homeTeam: string
  awayTeam: string
}

const fixtures = [
  { homeTeam: "Arsenal", awayTeam: "Tottenham" },
  { homeTeam: "Manchester United", awayTeam: "Chelsea" },
  { homeTeam: "Liverpool", awayTeam: "Manchester City" },
  { homeTeam: "Newcastle", awayTeam: "Aston Villa" },
  { homeTeam: "West Ham", awayTeam: "Brighton" }
];

function Fixture({ homeTeam, awayTeam }: fixture) {
  return (
    <>
      <div>{homeTeam} v {awayTeam}</div>
    </>
  )
}

function App() {
  return (
    <>
      <Header />
      <div className="flex flex-col items-center justify-center min-h-svh">
        <Fixture {...fixtures[0]}/>
        <Fixture {...fixtures[1]}/>
        <Fixture {...fixtures[2]}/>
        <Fixture {...fixtures[3]}/>
        <Fixture {...fixtures[4]}/>
        <Button>Click me</Button>
      </div>
    </>
  );
}

export default App;
