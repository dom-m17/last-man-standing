import { useEffect, useRef, useState } from "react";
import { Link } from "react-router-dom";
import styles from "./Header.module.css";

interface ProfileButtonProps {
  isLoggedIn: boolean; // You can later wire this to your auth state
}

export default function ProfileButton({
  isLoggedIn = false,
}: ProfileButtonProps) {
  const [isOpen, setIsOpen] = useState(false);
  const menuRef = useRef<HTMLDivElement>(null);

  // Close dropdown when clicking outside
  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (menuRef.current && !menuRef.current.contains(e.target as Node)) {
        setIsOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  // Close dropdown with Escape key
  useEffect(() => {
    const handleKey = (e: KeyboardEvent) => {
      if (e.key === "Escape") setIsOpen(false);
    };
    document.addEventListener("keydown", handleKey);
    return () => document.removeEventListener("keydown", handleKey);
  }, []);

  return (
    <div className={styles.profileWrapper} ref={menuRef}>
      <button
        className={styles.profileButton}
        onClick={() => setIsOpen((prev) => !prev)}
        aria-haspopup="menu"
        aria-expanded={isOpen}
      >
        <img
          src="./user-avatar.png"
          alt="Profile"
          className={styles.profileAvatar}
        />
      </button>

      {isOpen && (
        <div className={styles.profileMenu} role="menu">
          {isLoggedIn ? (
            <Link
              className="bg-sky-900"
              to="/profile"
              role="menuitem"
              onClick={() => setIsOpen(false)}
            >
              View Profile
            </Link>
          ) : (
            <>
              <Link
                className="bg-sky-900"
                to="/login"
                role="menuitem"
                onClick={() => setIsOpen(false)}
              >
                Log In
              </Link>
              <Link
                className="bg-sky-900"
                to="/signup"
                role="menuitem"
                onClick={() => setIsOpen(false)}
              >
                Sign Up
              </Link>
            </>
          )}
        </div>
      )}
    </div>
  );
}
