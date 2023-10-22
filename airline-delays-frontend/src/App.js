import React, { useState } from 'react';
import { Container, TextField, Button, Typography } from '@mui/material';
import './App.css';

function App() {
  const [city, setCity] = useState('');
  const [time, setTime] = useState('');
  const [result, setResult] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetchBackend(city, time);
    setResult(response);
  };

  const fetchBackend = async (city, time) => {
    try {
      const response = await fetch('http://localhost:8080/predict-delay', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          city: city,
          time: time,
        }),
      });
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      const data = await response.json();
  
      // Assuming your Go response has a field 'Delay' that indicates delay status
      // (e.g., 0 means not delayed, 1 means delayed)
      return data.Delay === 0 ? "Not Delayed" : "Delayed";
    } catch (error) {
      console.error("Failed to fetch prediction:", error);
      return "Error fetching prediction";
    }
  };
  

  return (
    <Container maxWidth="sm">
      <Typography variant="h4" align="center" gutterBottom>
        Flight Delay Predictor
      </Typography>
      <form onSubmit={handleSubmit}>
        <div>
          <TextField 
            fullWidth
            label="Enter City"
            variant="outlined"
            margin="normal"
            value={city}
            onChange={e => setCity(e.target.value)}
          />
        </div>
        <div>
          <TextField 
            fullWidth
            label="Enter Time"
            variant="outlined"
            margin="normal"
            value={time}
            onChange={e => setTime(e.target.value)}
          />
        </div>
        <Button variant="contained" color="primary" type="submit">
          Predict
        </Button>
      </form>
      {result && <Typography variant="h5" align="center" style={{marginTop: '20px'}}>Flight is: {result}</Typography>}
    </Container>
  );
}

export default App;
