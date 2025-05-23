import React from "react";
import {
  List,
  ListItem,
  ListItemText,
  Divider,
  Button,
  Typography,
} from "@mui/material";

export default function ExerciseList({
  exercises,
  onStartExercise,
  currentExerciseIndex,
  timerRunning,
  timerSeconds,
  intervalType,
}) {
  return (
    <>
      <Typography variant="h5" gutterBottom>
        Exercises
      </Typography>
      {exercises.length === 0 ? (
        <Typography>No exercises added yet.</Typography>
      ) : (
        <List>
          {exercises.map((ex, index) => (
            <React.Fragment key={index}>
              <ListItem
                secondaryAction={
                  <>
                    {currentExerciseIndex === index && timerRunning ? (
                      <Typography sx={{ mr: 2 }}>
                        Time Left: {timerSeconds}s{" "}
                        {ex.type === "Stretch" ? `(${intervalType})` : ""}
                      </Typography>
                    ) : null}
                    <Button
                      variant="contained"
                      size="small"
                      onClick={() => onStartExercise(index)}
                      disabled={timerRunning}
                    >
                      Start
                    </Button>
                  </>
                }
              >
                <ListItemText
                  primary={`${ex.name} (${ex.type})`}
                  secondary={
                    ex.type === "Stretch"
                      ? `Hold: ${ex.holdDuration}s, Rest: ${ex.restDuration}s, Limit: ${ex.limit}`
                      : `Reps per set: ${ex.repsPerSet}, Total sets: ${ex.totalSets}`
                  }
                />
              </ListItem>
              {index < exercises.length - 1 && <Divider />}
            </React.Fragment>
          ))}
        </List>
      )}
    </>
  );
}
