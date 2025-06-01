import Competition from "./components/Competition";
import Round from "./components/Round";
import MatchList from "./components/MatchList";
import Header from "./components/Header";

export default function App() {
  return (
    <main>
      <Header />
      <Competition />
      <Round />
      <MatchList />
    </main>
  );
}
