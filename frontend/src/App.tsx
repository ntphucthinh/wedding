import { Box, Button, Container, Paper, Typography } from '@mui/material';

function App() {
  return (
    <Container maxWidth="sm" sx={{ py: 6 }}>
      <Paper elevation={3} sx={{ p: 4, textAlign: 'center' }}>
        <Typography variant="h4" gutterBottom>
          Wedding Frontend
        </Typography>
        <Typography variant="body1" color="text.secondary" sx={{ mb: 3 }}>
          React + Vite + MUI + Refine + Yup + tsx starter template.
        </Typography>
        <Button variant="contained" color="primary">
          Start building
        </Button>
      </Paper>
    </Container>
  );
}

export default App;
