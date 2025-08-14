import MenuButton from "./MenuButton";
import styles from "./Header.module.css";

export default function Header() {
  return (
    <header className={styles.headerContainer}>
      <h1 className={styles.title}>Last Man Standing</h1>
      <MenuButton />
    </header>
  );
}
