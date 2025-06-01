import MenuButton from "./MenuButton"
import styles from "./Header.module.css"

export default function Header() {
  return (
    <header>
      <h1 className={styles.header}>Last Man Standing</h1>
      <MenuButton></MenuButton>
    </header>
  );
}
