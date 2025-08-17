import { useState } from "react";
import CompetitionRound from "./components/CompetitionRound";
import MatchList from "./components/MatchList";
import Header from "./components/Header";

export default function App() {
  const [round, setRound] = useState(1);
  const [competition, setCompetition] = useState(1);

  return (
    <main>
      <Header />
      <CompetitionRound
        round={round}
        setRound={setRound}
        competition={competition}
        setCompetition={setCompetition}
      />
      <MatchList roundNumber={round} competitionNumber={competition} />
    </main>
  );
}
