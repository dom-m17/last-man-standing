import Match from "./Match";
import styles from "./MatchList.module.css";

async function getMatchesByMatchday(matchday: number) {
  try {
    let result = await fetch("http://localhost:8080/query", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        query: `
          query {
            getMatchesByMatchday(input: ${matchday}) {
              id
            }
          }
        `,
      }),
    });
    let matches = await result.json();
    console.log(matches);
    return matches;
  } catch (err) {
    console.error("Error fetching matches:", err);
  }
}

export default function MatchList() {
  getMatchesByMatchday(1);

  return (
    <div>
      <ul className={styles.MatchList}>
        <Match
          homeTeamID={1}
          homeTeamName="Arsenal"
          awayTeamID={2}
          awayTeamName="Aston Villa"
        ></Match>
      </ul>
    </div>
  );
}
