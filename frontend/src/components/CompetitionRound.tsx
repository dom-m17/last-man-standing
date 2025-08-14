import styles from "./CompetitionRound.module.css";

interface CompetitionRoundProps {
  round: number;
  setRound: React.Dispatch<React.SetStateAction<number>>;
}

export default function CompetitionRound({
  round,
  setRound,
}: CompetitionRoundProps) {
  return (
    <div className={styles.container}>
      {/* Competition Selector - static for now */}
      <div className={styles.selector}>
        <button className={styles.arrow} aria-label="Previous competition">
          ◀
        </button>
        <span className={styles.label}>Comp 1</span>
        <button className={styles.arrow} aria-label="Next competition">
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
