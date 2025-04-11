import Link from "next/link";
import styles from "../styles/header.module.css";

export default function Header() {
  return (
    <header className={styles.header}>
      <Link href="/">Home</Link>
      <Link href="/register">Register</Link>
      <Link href="/login">Login</Link>
    </header>
  );
}
