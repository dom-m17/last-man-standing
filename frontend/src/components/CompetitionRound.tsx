import styles from "./CompetitionRound.module.css";

interface CompetitionRoundProps {
  round: number;
  setRound: React.Dispatch<React.SetStateAction<number>>;
  competition: number;
  setCompetition: React.Dispatch<React.SetStateAction<number>>;
}

export default function CompetitionRound({
  round,
  setRound,
  competition,
  setCompetition,
}: CompetitionRoundProps) {
  return (
    <div className={styles.container}>
      {/* Competition Selector */}
      <div className={styles.selector}>
        <button
          className={styles.arrow}
          aria-label="Previous competition"
          onClick={() => setCompetition((prev) => Math.max(prev - 1, 1))}
        >
          ◀
        </button>
        <span className={styles.label}>Comp {competition}</span>
        <button
          className={styles.arrow}
          aria-label="Next competition"
          onClick={() => setCompetition((prev) => prev + 1)}
        >
          ▶
        </button>
      </div>

      {/* Round Selector */}
      <div className={styles.selector}>
        <button
          className={styles.arrow}
          aria-label="Previous round"
          onClick={() => setRound((prev) => Math.max(prev - 1, 1))}
        >
          ◀
        </button>
        <span className={styles.label}>Round {round}</span>
        <button
          className={styles.arrow}
          aria-label="Next round"
          onClick={() => setRound((prev) => prev + 1)}
        >
          ▶
        </button>
      </div>
    </div>
  );
}
