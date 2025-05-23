import React, { useState,useEffect } from "react";
import { Container, Typography, Paper } from "@mui/material";
import ExerciseForm from "./components/ExerciseForm";
import ExerciseList from "./components/ExerciseList";
import Timer from "./components/Timer";
import Summary from "./components/Summary";

function App() {
  const [exercises, setExercises] = useState([]);
  const [sessionComplete, setSessionComplete] = useState(false);

  const [currentExerciseIndex, setCurrentExerciseIndex] = useState(null);
  const [timerSeconds, setTimerSeconds] = useState(0);
  const [timerRunning, setTimerRunning] = useState(false);
  const [intervalType, setIntervalType] = useState("hold");
  const [completedReps, setCompletedReps] = useState(0);
  const [completedSets, setCompletedSets] = useState(0);

  const addExercise = (exercise) => setExercises((exs) => [...exs, exercise]);
  useEffect(() => {
    const stored = localStorage.getItem("exercises");
    if (stored) {
      setExercises(JSON.parse(stored));
    }
  }, []);

  // Save exercises whenever they change
  useEffect(() => {
    localStorage.setItem("exercises", JSON.stringify(exercises));
  }, [exercises]);
  const startTimer = (index) => {
    if (timerRunning) {
      alert("Timer already running!");
      return;
    }
    setCurrentExerciseIndex(index);
    const ex = exercises[index];
    if (ex.type === "Stretch") {
      setIntervalType("hold");
      setTimerSeconds(ex.holdDuration);
      setCompletedReps(0);
    } else {
      setTimerSeconds(ex.repsPerSet);
      setCompletedSets(0);
    }
    setTimerRunning(true);
  };

  return (
    <Container maxWidth="sm" sx={{ mt: 5 }}>
      <Typography variant="h4" gutterBottom>
        Exercise Timer
      </Typography>

      <ExerciseForm onAddExercise={addExercise} />

      <Paper sx={{ p: 2 }}>
        <ExerciseList
          exercises={exercises}
          onStartExercise={startTimer}
          currentExerciseIndex={currentExerciseIndex}
          timerRunning={timerRunning}
          timerSeconds={timerSeconds}
          intervalType={intervalType}
        />
      </Paper>
      {sessionComplete ? (
        <Summary
          exercise={exercises[currentExerciseIndex]}
          onReset={() => {
            setSessionComplete(false);
            setCurrentExerciseIndex(null);
            setCompletedReps(0);
            setCompletedSets(0);
            setIntervalType("hold");
          }}
        />
      ) : (
        <Timer
          timerRunning={timerRunning}
          setTimerRunning={setTimerRunning}
          timerSeconds={timerSeconds}
          setTimerSeconds={setTimerSeconds}
          currentExerciseIndex={currentExerciseIndex}
          setCurrentExerciseIndex={setCurrentExerciseIndex}
          exercises={exercises}
          intervalType={intervalType}
          setIntervalType={setIntervalType}
          completedReps={completedReps}
          setCompletedReps={setCompletedReps}
          completedSets={completedSets}
          setCompletedSets={setCompletedSets}
          setSessionComplete={setSessionComplete}
        />
      )}
    </Container>
  );
}

export default App;
