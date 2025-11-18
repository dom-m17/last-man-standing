import MenuButton from "./MenuButton";
import ProfileButton from "./ProfileButton";
import styles from "./Header.module.css";
import { Link } from "react-router-dom";

export default function Header() {
  return (
    <header className={styles.headerContainer}>
      <Link to="/" className={styles.title}>
        LMS
      </Link>
      <ProfileButton isLoggedIn={false} />
      <MenuButton />
    </header>
  );
}
