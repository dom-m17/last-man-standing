import { useState } from "react";
import CourseGoalList from "./components/CourseGoalList";
import Header from "./components/Header";
import goalsImg from "./assests/goals.jpg";
import NewGoal from "./components/NewGoal";

export type CourseGoal = {
  title: string;
  description: string;
  id: number;
};

export default function App() {
  const [goals, setGoals] = useState<CourseGoal[]>([]);

  function handleAddGoal(goal: string, summary: string) {
    setGoals((prevGoals) => {
      const newGoal: CourseGoal = {
        id: Math.random(),
        title: goal,
        description: summary,
      };
      return [...prevGoals, newGoal];
    });
  }

  function handleDeleteGoal(id: Number) {
    setGoals(prevGoals => prevGoals.filter(goal => goal.id !== id))
  }

  return (
    <main>
      <Header img={{ src: goalsImg, alt: "A list of goals" }}>
        <h1>Your Course Goals</h1>
      </Header>
      <NewGoal onAddGoal={handleAddGoal}></NewGoal>
      <CourseGoalList goals={goals} onDelete={handleDeleteGoal}></CourseGoalList>
    </main>
  );
}
