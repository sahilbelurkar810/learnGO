import React, { useEffect } from "react";
import { Button, Typography, Stack, LinearProgress } from "@mui/material";

const Timer = ({
  timerRunning,
  setTimerRunning,
  timerSeconds,
  setTimerSeconds,
  currentExerciseIndex,
  setCurrentExerciseIndex,
  exercises,
  intervalType,
  setIntervalType,
  completedReps,
  setCompletedReps,
  completedSets,
  setCompletedSets,
  setSessionComplete,
}) => {
  const [paused, setPaused] = React.useState(false);
  const beepSound = new Audio("/sounds/beep.mp3");
  const completeSound = new Audio("/sounds/complete.mp3");
  
  useEffect(() => {
    let timer;
    if (timerRunning && !paused && timerSeconds > 0) {
      timer = setInterval(() => {
        setTimerSeconds((prev) => prev - 1);
      }, 1000);
    }

    if (timerSeconds === 0 && timerRunning && !paused) {
      handleTimerEnd();
    }

    return () => clearInterval(timer);
  }, [timerRunning, paused, timerSeconds]);

  const handlePause = () => setPaused(true);
  const handleResume = () => setPaused(false);
  const handleStop = () => {
    setPaused(false);
    setTimerRunning(false);
    setTimerSeconds(0);
  };

  const handleTimerEnd = () => {
    const currentExercise = exercises[currentExerciseIndex];
    if (!currentExercise) return;

    if (currentExercise.type === "Stretch") {
      if (intervalType === "hold") {
        beepSound.play(); // Play beep when switching to rest
        setIntervalType("rest");
        setTimerSeconds(currentExercise.restDuration);
      } else {
        const nextSet = completedSets + 1;
        if (nextSet < currentExercise.limit) {
          beepSound.play(); // Play beep when switching back to hold
          setIntervalType("hold");
          setCompletedSets(nextSet);
          setTimerSeconds(currentExercise.holdDuration);
        } else {
          // Reset for the next exercise
          setCompletedSets(0);
          setIntervalType("hold");
          if (currentExerciseIndex < exercises.length - 1) {
            setCurrentExerciseIndex(currentExerciseIndex + 1);
            setTimerSeconds(exercises[currentExerciseIndex + 1].holdDuration);
          } else {
            completeSound.play(); // Play complete sound at session end
            setTimerRunning(false);
            setSessionComplete(true);
          }
        }
      }
    } else if (currentExercise.type === "General") {
      const nextRep = completedReps + 1;
      if (nextRep < currentExercise.totalSets) {
        setCompletedReps(nextRep);
        setTimerSeconds(currentExercise.repsPerSet);
      } else {
        setCompletedReps(0);
        if (currentExerciseIndex < exercises.length - 1) {
          setCurrentExerciseIndex(currentExerciseIndex + 1);
          setTimerSeconds(exercises[currentExerciseIndex + 1].repsPerSet);
        } else {
          completeSound.play(); // Play complete sound at session end
          setTimerRunning(false);
          setSessionComplete(true);
        }
      }
    }
  };

  return (
    <div style={{ marginTop: 20 }}>
      <Typography variant="h6">Timer: {timerSeconds}s</Typography>
      <Stack direction="row" spacing={2} mt={2}>
        {timerRunning && !paused && (
          <Button variant="outlined" onClick={handlePause}>
            Pause
          </Button>
        )}
        {timerRunning && paused && (
          <Button variant="contained" onClick={handleResume}>
            Resume
          </Button>
        )}
        {timerRunning && (
          <Button variant="outlined" color="error" onClick={handleStop}>
            Stop
          </Button>
        )}
      </Stack>
      {timerRunning && (
        <>
          <Typography variant="body1" mt={2}>
            Progress:
          </Typography>
          {exercises[currentExerciseIndex]?.type === "Stretch" && (
            <LinearProgress
              variant="determinate"
              value={
                (completedSets / exercises[currentExerciseIndex].limit) * 100
              }
            />
          )}
          {exercises[currentExerciseIndex]?.type === "General" && (
            <LinearProgress
              variant="determinate"
              value={
                (completedReps / exercises[currentExerciseIndex].totalSets) *
                100
              }
            />
          )}
        </>
      )}
    </div>
  );
};

export default Timer;
