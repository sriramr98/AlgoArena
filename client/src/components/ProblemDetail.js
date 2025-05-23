import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import CodeEditor from './CodeEditor';
import './ProblemDetail.css';
import { Panel, PanelGroup, PanelResizeHandle } from 'react-resizable-panels';

const ProblemDetail = () => {
  const { id } = useParams();
  const [problem, setProblem] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  
  useEffect(() => {
    const fetchProblem = async () => {
      try {
        console.log('Fetching problem with id:', id);
        const response = await axios.get(`http://localhost:5000/api/problems/${id}`);
        console.log('Problem data received:', response.data);
        setProblem(response.data);
        setLoading(false);
      } catch (err) {
        setError('Failed to fetch problem details');
        setLoading(false);
        console.error('Error fetching problem:', err);
      }
    };

    fetchProblem();
  }, [id]);

  if (loading) {
    return <div className="loading">Loading problem...</div>;
  }

  if (error) {
    return <div className="error">{error}</div>;
  }

  if (!problem) {
    return <div className="error">Problem not found</div>;
  }

  return (
    <div className="problem-detail-container">
      <PanelGroup direction="horizontal">
        <Panel defaultSize={40} minSize={25} className="problem-description-panel">
          <div className="problem-description">
            <div className="problem-header">
              <h1>{problem.title}</h1>
              <span className={`difficulty ${problem.difficulty.toLowerCase()}`}>
                {problem.difficulty}
              </span>
            </div>
            
            <div className="problem-content">
              <div className="description">
                {problem.description.split('\\n\\n').map((paragraph, idx) => (
                  <p key={idx}>{paragraph}</p>
                ))}
              </div>

              <div className="examples">
                <h3>Examples:</h3>
                {problem.examples.map((example, idx) => (
                  <div className="example" key={idx}>
                    <div className="example-input">
                      <strong>Input:</strong> {example.input}
                    </div>
                    <div className="example-output">
                      <strong>Output:</strong> {example.output}
                    </div>
                    {example.explanation && (
                      <div className="example-explanation">
                        <strong>Explanation:</strong> {example.explanation}
                      </div>
                    )}
                  </div>
                ))}
              </div>

              <div className="constraints">
                <h3>Constraints:</h3>
                <ul>
                  {problem.constraints.map((constraint, idx) => (
                    <li key={idx}>{constraint}</li>
                  ))}
                </ul>
              </div>
            </div>
          </div>
        </Panel>
        
        <PanelResizeHandle className="resize-handle" />
        
        <Panel defaultSize={60} minSize={30} className="code-editor-panel">
          <div className="code-editor-container" style={{ height: '100%', width: '100%', display: 'flex', flexDirection: 'column', overflow: 'hidden' }}>
            <CodeEditor 
              language="javascript"
              problemId={id}
              key={id} // Add a key to force re-render when problem changes
            />
          </div>
        </Panel>
      </PanelGroup>
    </div>
  );
};

export default ProblemDetail;
