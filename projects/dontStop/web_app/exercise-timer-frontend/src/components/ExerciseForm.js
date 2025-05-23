import React, { useState } from "react";
import {
  TextField,
  Button,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Box,
} from "@mui/material";

const exerciseTypes = ["Stretch", "General"];

export default function ExerciseForm({ onAddExercise }) {
  const [name, setName] = useState("");
  const [type, setType] = useState("Stretch");
  const [holdDuration, setHoldDuration] = useState("");
  const [restDuration, setRestDuration] = useState("");
  const [limit, setLimit] = useState("");
  const [repsPerSet, setRepsPerSet] = useState("");
  const [totalSets, setTotalSets] = useState("");

  const handleSubmit = () => {
    if (!name.trim()) return alert("Please enter exercise name");

    const newExercise = { name, type };

    if (type === "Stretch") {
      newExercise.holdDuration = parseInt(holdDuration);
      newExercise.restDuration = parseInt(restDuration);
      newExercise.limit = parseInt(limit);
    } else {
      newExercise.repsPerSet = parseInt(repsPerSet);
      newExercise.totalSets = parseInt(totalSets);
    }

    onAddExercise(newExercise);

    setName("");
    setHoldDuration("");
    setRestDuration("");
    setLimit("");
    setRepsPerSet("");
    setTotalSets("");
  };

  return (
    <Box sx={{ display: "flex", flexDirection: "column", gap: 2, mb: 4 }}>
      <TextField
        label="Exercise Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />

      <FormControl fullWidth>
        <InputLabel>Type</InputLabel>
        <Select
          value={type}
          label="Type"
          onChange={(e) => setType(e.target.value)}
        >
          {exerciseTypes.map((exType) => (
            <MenuItem key={exType} value={exType}>
              {exType}
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      {type === "Stretch" ? (
        <>
          <TextField
            label="Hold Duration (seconds)"
            type="number"
            value={holdDuration}
            onChange={(e) => setHoldDuration(e.target.value)}
          />
          <TextField
            label="Rest Duration (seconds)"
            type="number"
            value={restDuration}
            onChange={(e) => setRestDuration(e.target.value)}
          />
          <TextField
            label="Limit (repetitions)"
            type="number"
            value={limit}
            onChange={(e) => setLimit(e.target.value)}
          />
        </>
      ) : (
        <>
          <TextField
            label="Reps per Set"
            type="number"
            value={repsPerSet}
            onChange={(e) => setRepsPerSet(e.target.value)}
          />
          <TextField
            label="Total Sets"
            type="number"
            value={totalSets}
            onChange={(e) => setTotalSets(e.target.value)}
          />
        </>
      )}

      <Button variant="contained" onClick={handleSubmit}>
        Add Exercise
      </Button>
    </Box>
  );
}
