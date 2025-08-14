import { useEffect, useState } from "react";
import styles from "./Header.module.css";

const MenuButton = () => {
  const [isOpen, setIsOpen] = useState(false);

  // Optional: close on Escape for accessibility
  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if (e.key === "Escape") setIsOpen(false);
    };
    document.addEventListener("keydown", onKey);
    return () => document.removeEventListener("keydown", onKey);
  }, []);

  return (
    <>
      {/* Hamburger button (stays above the drawer) */}
      <button
        className={`${styles.menuButton} ${isOpen ? styles.active : ""}`}
        onClick={() => setIsOpen(!isOpen)}
        aria-label="Toggle menu"
        aria-expanded={isOpen}
        aria-controls="lms-side-menu"
      >
        <span className={styles.bar}></span>
        <span className={styles.bar}></span>
        <span className={styles.bar}></span>
      </button>

      {/* Backdrop */}
      <div
        className={`${styles.backdrop} ${isOpen ? styles.backdropOpen : ""}`}
        onClick={() => setIsOpen(false)}
      />

      {/* Side menu */}
      <nav
        id="lms-side-menu"
        className={`${styles.sideMenu} ${isOpen ? styles.sideMenuOpen : ""}`}
        role="navigation"
      >
        <a href="#" className={styles.menuLink}>
          Home
        </a>
        <a href="#" className={styles.menuLink}>
          Fixtures
        </a>
        <a href="#" className={styles.menuLink}>
          Standings
        </a>
        <a href="#" className={styles.menuLink}>
          Rules
        </a>
        <a href="#" className={styles.menuLink}>
          Logout
        </a>
      </nav>
    </>
  );
};

export default MenuButton;
