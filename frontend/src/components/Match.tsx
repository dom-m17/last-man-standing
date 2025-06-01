import styles from './Match.module.css'

type MatchProps = {
    homeTeamID: number
    homeTeamName: string
    awayTeamID: number
    awayTeamName: string
}

const Match = (props: MatchProps) => {
  return (
    <li className={styles.Match}>
      <button id={props.homeTeamID.toString()}>{props.homeTeamName}</button>
      <p>v</p>
      <button id={props.awayTeamID.toString()}>{props.awayTeamName}</button>
    </li>
  );
};

export default Match;
