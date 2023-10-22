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
      return data.Delay === 0 ? "On Time" : "Delayed";
    } catch (error) {
      console.error("Failed to fetch prediction:", error);
      return "Error fetching prediction";
    }
  };

  return (
    <Container maxWidth="sm" style={{ height: '100vh', display: 'flex', flexDirection: 'column', justifyContent: 'center' }}>
      <Typography variant="h4" align="center" gutterBottom style={{ color: '#f4f1f1' }}>
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
            InputProps={{ style: { color: 'white' } }}
            InputLabelProps={{ style: { color: 'white' } }}
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
            InputProps={{ style: { color: 'white' } }}
            InputLabelProps={{ style: { color: 'white' } }}
          />
        </div>
        <Button variant="contained" color="primary" type="submit" style={{ marginTop: '10px' }}>
          Predict
        </Button>
      </form>
      {result && (
        <Typography 
        variant="h5" 
        align="center" 
        style={{
            marginTop: '20px', 
            border: '1px solid', 
            padding: '10px', 
            borderRadius: '8px',
            borderColor: result === "Delayed" ? 'red' : 'green',
            color: result === "Delayed" ? 'red' : '#90EE90',
        }}
    >
          Flight Prediction: {result}
      </Typography>
    
      )}
    </Container>
  );
}

export default App;
