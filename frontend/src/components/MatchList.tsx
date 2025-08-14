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
  round: number;
}

export default function MatchList({ round }: MatchListProps) {
  const [matches, setMatches] = useState<MatchType[]>([]);

  useEffect(() => {
    getMatchesByMatchday(round).then(setMatches);
  }, [round]); // refetch when round changes

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
