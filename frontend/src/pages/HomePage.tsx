import CompetitionRound from "../components/CompetitionRound";
import MatchList from "../components/MatchList";
import { useState } from "react";

export default function HomePage() {
  const [round, setRound] = useState(1);
  const [competition, setCompetition] = useState(1);

  return (
    <>
      <CompetitionRound
        round={round}
        setRound={setRound}
        competition={competition}
        setCompetition={setCompetition}
      />
      <MatchList roundNumber={round} competitionNumber={competition} />
    </>
  );
}
