import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import ProblemsList from './components/ProblemsList';
import ProblemDetail from './components/ProblemDetail';
import Navbar from './components/Navbar';

function App() {
  return (
    <Router>
      <div className="App">
        <Navbar />
        <Routes>
          <Route path="/" element={<ProblemsList />} />
          <Route path="/problem/:id" element={<ProblemDetail />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
