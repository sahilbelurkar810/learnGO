import React from "react";
import { Typography, Button, Paper, Stack } from "@mui/material";

const Summary = ({ exercise, onReset }) => {
  return (
    <Paper elevation={3} sx={{ p: 4, mt: 4 }}>
      <Typography variant="h5" gutterBottom>
        Exercise Summary
      </Typography>
      <Typography variant="subtitle1">Name: {exercise.name}</Typography>
      <Typography variant="subtitle1">Type: {exercise.type}</Typography>

      {exercise.type === "stretch" ? (
        <>
          <Typography>Hold: {exercise.holdDuration}s</Typography>
          <Typography>Rest: {exercise.restDuration}s</Typography>
          <Typography>Intervals: {exercise.limit}</Typography>
        </>
      ) : (
        <>
          <Typography>Reps per Set: {exercise.repsPerSet}</Typography>
          <Typography>Total Sets: {exercise.totalSets}</Typography>
        </>
      )}

      <Stack direction="row" spacing={2} mt={3}>
        <Button variant="outlined" onClick={onReset}>
          Back to Home
        </Button>
      </Stack>
    </Paper>
  );
};

export default Summary;
