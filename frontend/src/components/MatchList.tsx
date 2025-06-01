import Match from "./Match";
import styles from "./MatchList.module.css";

async function getMatchesByMatchday(matchday: number) {
    try {
        let result = await fetch("http://localhost/8080/query", {
            method: "GET",
            headers: {
            "Content-Type": "application/json",
            },
            body: JSON.stringify({
                query: `
                    getMatchesByMatchday
                `
            })
        })
        let matches = await result.json()
        console.log(matches.homeTeam)
    } catch (err) {
        console.log("fail")
    }
}

export default function MatchList() {

  getMatchesByMatchday(1)

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
