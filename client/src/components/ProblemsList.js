import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import "./ProblemsList.css";

const ProblemsList = () => {
  const [problems, setProblems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchProblems = async () => {
      try {
        const response = await axios.get("http://localhost:5000/api/problems");
        setProblems([...response.data]);
        setLoading(false);
      } catch (err) {
        setError("Failed to fetch problems");
        setLoading(false);
        console.error("Error fetching problems:", err);
      }
    };

    fetchProblems();
  }, []);

  if (loading) {
    return <div className="loading">Loading problems...</div>;
  }

  if (error) {
    return <div className="error">{error}</div>;
  }

  return (
    <div className="problems-container">
      <h1>AlgoArena Problems</h1>
      <div className="problems-table">
        <div className="table-header">
          <div className="col-title-header">Title</div>
          <div className="col-difficulty">Difficulty</div>
        </div>
        {problems.map((problem, ind) => (
          <Link to={`/problem/${problem.id}`} key={ind} className="table-row">
            <div className="col-title">{problem.title}</div>
            <div className={`col-difficulty ${problem.difficulty.toLowerCase()}`}>{problem.difficulty}</div>
          </Link>
        ))}
      </div>
    </div>
  );
};

export default ProblemsList;
