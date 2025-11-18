import { useEffect, useState } from "react";
import styles from "./MatchList.module.css";

interface Team {
  id: string;
  longName: string;
}

interface MatchType {
  id: string;
  homeTeam: Team;
  awayTeam: Team;
}

interface Rounds {
  id: string;
  roundNumber: number;
  matchday: number;
}

interface CompetitionRounds {
  id: string;
  startMatchday: number;
  rounds: Rounds[];
}

async function getCompetitions(): Promise<CompetitionRounds[]> {
  try {
    const result = await fetch("http://localhost:8080/query", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        query: `
          query MyQuery{
            listCompetitions {
                id
                  name
                  startMatchday
                  status
            }
          }
        `,
      }),
    });

    const json = await result.json();
    return json.data?.listCompetitions || [];
  } catch (err) {
    console.error("Error fetching competitions:", err);
    return [];
  }
}

// TODO: Abstract this - create a query builder function
async function getMatchesByMatchday(matchday: number): Promise<MatchType[]> {
  try {
    const result = await fetch("http://localhost:8080/query", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        query: `
          query MyQuery {
            getMatchesByMatchday(input: ${matchday}) {
              id
              homeTeam {
                id
                longName
              }
              awayTeam {
                id
                longName
              }
            }
          }
        `,
      }),
    });

    const json = await result.json();
    return json.data?.getMatchesByMatchday || [];
  } catch (err) {
    console.error("Error fetching matches:", err);
    return [];
  }
}

interface MatchListProps {
  roundNumber: number;
  competitionNumber: number;
}

export default function MatchList({
  roundNumber,
  competitionNumber,
}: MatchListProps) {
  const [matches, setMatches] = useState<MatchType[]>([]);
  const [competitions, setCompetitions] = useState<CompetitionRounds[]>([]);

  useEffect(() => {
    getCompetitions().then(setCompetitions);
  }, []);

  //! untested
  useEffect(() => {
    if (!competitions.length) return;
    const matchday =
      competitions[competitionNumber - 1].startMatchday + roundNumber - 1;
    getMatchesByMatchday(matchday).then(setMatches);
  }, [competitions, roundNumber, competitionNumber]);

  return (
    <div className={styles.MatchList}>
      {matches.map((match) => (
        <div key={match.id} className={styles.MatchItem}>
          <span className={`${styles.TeamName} ${styles.Home}`}>
            {match.homeTeam.longName}
          </span>
          <span>vs</span>
          <span className={`${styles.TeamName} ${styles.Away}`}>
            {match.awayTeam.longName}
          </span>
        </div>
      ))}
    </div>
  );
}
