import React, { useState, useRef, useEffect } from "react";
import Editor from "@monaco-editor/react";
import axios from "axios";
import "./CodeEditor.css";
import { Panel, PanelGroup, PanelResizeHandle } from "react-resizable-panels";

// Collapsible Array component for large arrays
const CollapsibleArray = ({ array }) => {
  const [isExpanded, setIsExpanded] = useState(false);

  // Helper to format single items in arrays
  const formatItem = (item) => {
    if (typeof item === "string") return `"${item}"`;
    if (Array.isArray(item)) {
      if (item.length <= 3) {
        return `[${item.map(formatItem).join(", ")}]`;
      } else {
        return `[${item.slice(0, 3).map(formatItem).join(", ")}, ...]`;
      }
    }
    if (typeof item === "object" && item !== null) return JSON.stringify(item);
    return String(item);
  };

  if (array.length <= 10) {
    // For small arrays, just show all elements in a single line
    return <span>[{array.map(formatItem).join(", ")}]</span>;
  }

  return (
    <div className="collapsible-array">
      {isExpanded ? (
        <>
          <pre className="expanded-array">
            [
            {array.map((item, i) => (
              <React.Fragment key={i}>
                {formatItem(item)}
                {i < array.length - 1 ? "," : ""}
                {i < array.length - 1 && <br />}
              </React.Fragment>
            ))}
            ]
          </pre>
          <button className="toggle-array-btn" onClick={() => setIsExpanded(false)}>
            Show Less
          </button>
        </>
      ) : (
        <>
          <span className="collapsed-array">[{array.slice(0, 7).map(formatItem).join(", ")}, ...]</span>
          <button className="toggle-array-btn" onClick={() => setIsExpanded(true)}>
            Show All ({array.length} items)
          </button>
        </>
      )}
    </div>
  );
};

// Helper to format output consistently across all views
const formatModalOutput = (output) => {
  const formattedValue = formatOutput(output);

  // Handle arrays with special formatting
  if (Array.isArray(formattedValue)) {
    // Use formatSingleLineArray for arrays to ensure consistent display
    return formatSingleLineArray(formattedValue);
  }

  // Handle objects
  if (typeof formattedValue === "object" && formattedValue !== null) {
    return JSON.stringify(formattedValue, null, 2);
  }

  // Handle primitive values
  if (typeof formattedValue === "string") {
    return `"${formattedValue}"`;
  }

  return String(formattedValue);
};

// Utility function to format test case inputs
const formatTestInput = (input) => {
  if (!input) return [];

  try {
    // If input is a string, try to parse it
    const parsedInput = typeof input === "string" ? JSON.parse(input) : input;

    return Object.entries(parsedInput).map(([key, valueObj]) => {
      let rawValue;
      let displayType = valueObj.type || typeof valueObj;

      // Extract the actual value
      if (valueObj.value !== undefined) {
        rawValue = valueObj.value;
      } else {
        rawValue = valueObj;
      }

      // Determine the type
      if (Array.isArray(rawValue)) {
        displayType = "array";
      } else if (typeof rawValue === "object" && rawValue !== null) {
        displayType = "object";
      } else if (typeof rawValue === "string") {
        displayType = "string";
      } else if (typeof rawValue === "number") {
        displayType = "number";
      } else if (typeof rawValue === "boolean") {
        displayType = "boolean";
      }

      return {
        name: key,
        rawValue: rawValue,
        value: rawValue,
        type: displayType
      };
    });
  } catch (error) {
    console.error("Error formatting test input:", error);
    return [];
  }
};

// Helper to format arrays as a single line
const formatSingleLineArray = (array) => {
  if (array.length <= 10) {
    return `[${array.map(formatArrayItem).join(", ")}]`;
  } else {
    // For longer arrays, show first 8 elements with ellipsis indicating how many more
    return `[${array.slice(0, 8).map(formatArrayItem).join(", ")}, ... ${array.length - 8} more]`;
  }
};

// Helper to format individual array items
const formatArrayItem = (item) => {
  if (typeof item === "string") return `"${item}"`;
  if (Array.isArray(item)) {
    if (item.length <= 3) {
      return `[${item.map(formatArrayItem).join(", ")}]`;
    } else {
      return `[${item.slice(0, 2).map(formatArrayItem).join(", ")}, ...]`;
    }
  }
  if (typeof item === "object" && item !== null) return JSON.stringify(item);
  return String(item);
};

// Utility function to format expected/actual outputs
const formatOutput = (output) => {
  if (output === null || output === undefined) return null;

  try {
    // If it's a string, try to parse it as JSON
    let parsedOutput = output;
    if (typeof output === "string") {
      try {
        parsedOutput = JSON.parse(output);
      } catch (e) {
        // If parsing fails, return the original string
        return output;
      }
    }

    // Check if the output has the structure { value: ..., type: ... }
    if (parsedOutput && typeof parsedOutput === "object" && "value" in parsedOutput && "type" in parsedOutput) {
      // Return just the value based on the type
      return parsedOutput.value;
    }

    return parsedOutput;
  } catch (error) {
    console.error("Error formatting output:", error);
    return output;
  }
};

// Helper function to format display of output values
const formatDisplayValue = (value) => {
  // Helper to format single items in arrays
  const formatItem = (item) => {
    if (typeof item === "string") return `"${item}"`;
    if (Array.isArray(item)) {
      if (item.length <= 3) {
        return `[${item.map(formatItem).join(", ")}]`;
      } else {
        return `[${item.slice(0, 3).map(formatItem).join(", ")}, ...]`;
      }
    }
    if (typeof item === "object" && item !== null) return JSON.stringify(item);
    return String(item);
  };

  try {
    // Format arrays on a single line for better readability
    if (Array.isArray(value)) {
      return formatSingleLineArray(value);
    }

    // Handle objects with pretty indentation
    if (typeof value === "object" && value !== null && !Array.isArray(value)) {
      return JSON.stringify(value, null, 2);
    }

    // Format strings with quotes
    if (typeof value === "string") {
      return `"${value}"`;
    }

    // For other primitive types, convert to string
    return String(value);
  } catch (error) {
    console.error("Error formatting display value:", error);
    return String(value);
  }
};

// Wrapper component for formatted output
const FormattedOutput = ({ value }) => {
  const formattedValue = formatOutput(value);

  // Determine the type of the value
  let valueType = typeof formattedValue;
  if (formattedValue === null) valueType = "null";
  else if (Array.isArray(formattedValue)) valueType = "array";
  else if (valueType === "object" && formattedValue !== null) valueType = "object";

  return (
    <div className="param-item output-item">
      <div className="param-header">
        <div className="param-name">Output</div>
        <div className="param-type output-type">{valueType}</div>
      </div>
      <div className="param-value">
        {Array.isArray(formattedValue) ? (
          <CollapsibleArray array={formattedValue} />
        ) : valueType === "object" ? (
          <pre>{JSON.stringify(formattedValue, null, 2)}</pre>
        ) : valueType === "string" ? (
          <pre>"{formattedValue}"</pre>
        ) : (
          <pre>{String(formattedValue)}</pre>
        )}
      </div>
    </div>
  );
};

const CodeEditor = ({ language = "javascript", problemId = "" }) => {
  const editorRef = useRef(null);

  const [currentLanguage, setCurrentLanguage] = useState(language);
  const [isEditorReady, setIsEditorReady] = useState(false);
  const [codeValue, setCodeValue] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [results, setResults] = useState(null);
  const [error, setError] = useState(null);
  const [previewTestCases, setPreviewTestCases] = useState([]);
  const [activeTestTab, setActiveTestTab] = useState(0);
  const [testResults, setTestResults] = useState({});
  const [showTestPanel, setShowTestPanel] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [modalMessage, setModalMessage] = useState("");
  const [modalSuccess, setModalSuccess] = useState(false);

  // Function to format language for Monaco Editor
  const getMonacoLanguage = (lang) => {
    switch (lang.toLowerCase()) {
      case "javascript":
        return "javascript";
      case "python":
        return "python";
      case "cpp":
        return "cpp";
      case "java":
        return "java";
      case "ruby":
        return "ruby";
      case "go":
        return "go";
      case "rust":
        return "rust";
      default:
        return "javascript";
    }
  };

  // Handle editor mount
  const handleEditorDidMount = (editor, monaco) => {
    editorRef.current = editor;
    setIsEditorReady(true);

    // Add command palette shortcut
    editor.addCommand(monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyP, () => {
      // Show command palette
      editor.trigger("", "editor.action.quickCommand");
    });

    // Set up a resize observer to manually handle editor layout
    const resizeObserver = new ResizeObserver(() => {
      // Debounce the layout call to prevent loops
      if (editorRef.current) {
        setTimeout(() => {
          try {
            editorRef.current.layout();
          } catch (err) {
            console.warn("Editor layout error:", err);
          }
        }, 100);
      }
    });

    // Observe the container element
    const container = document.querySelector(".editor-container");
    if (container) {
      resizeObserver.observe(container);

      // Store the observer in a ref for cleanup
      window.addEventListener("beforeunload", () => {
        resizeObserver.disconnect();
      });
    }
  };

  const handleLanguageChange = async (e) => {
    const newLanguage = e.target.value;
    setCurrentLanguage(newLanguage);

    // Clear any test results when language changes
    const initialResults = {};
    previewTestCases.forEach((_, index) => {
      initialResults[index] = null;
    });
    setTestResults(initialResults);
    setResults(null);
    setError(null);
  };

  const fetchLanguageStub = async (lang) => {
    try {
      const response = await axios.get(`http://localhost:5000/api/problems/${problemId}/stub`, {
        params: { language: lang }
      });

      if (response.data && response.data.codeTemplate) {
        setCodeValue(response.data.codeTemplate);
        return;
      }
    } catch (err) {
      console.error("Error fetching code template:", err);
      setCodeValue("");
    }
  };

  const fetchPreviewTestCases = async () => {
    try {
      const response = await axios.get(`http://localhost:5000/api/problems/${problemId}/testcases`);
      if (response.data) {
        setPreviewTestCases(response.data);
        // Initialize an empty result object for each test case
        const initialResults = {};
        response.data.forEach((_, index) => {
          initialResults[index] = null;
        });
        setTestResults(initialResults);
      }
    } catch (err) {
      console.error("Error fetching test cases:", err);
    }
  };

  const handleCodeExecution = async (isPreview) => {
    if (editorRef.current) {
      try {
        setIsSubmitting(true);
        setError(null);

        if (!isPreview) {
          // Clear results first and then wait a moment for the UI to stabilize
          setResults(null);
          await new Promise((resolve) => setTimeout(resolve, 100));
        }

        const code = editorRef.current.getValue();
        console.log(`${isPreview ? "Running" : "Submitting"} code:`, code);

        // Send the code to the backend for testing
        const response = await axios.post(`http://localhost:5000/api/submit?preview=${isPreview}`, {
          code,
          language: currentLanguage,
          problemId: problemId
        });

        // Small delay before updating state to allow UI to stabilize
        setTimeout(() => {
          if (isPreview) {
            // Update test results for each test case
            const newTestResults = { ...testResults };
            response.data.testResults.forEach((result, index) => {
              newTestResults[index] = result;
            });
            setTestResults(newTestResults);
          } else {
            // Set the results but don't display them in the UI
            setResults(response.data);

            // Show result in a modal instead
            const { passed, total, success, testResults } = response.data;

            if (success) {
              setModalMessage("You have successfully completed the problem!");
              setModalSuccess(true);
            } else {
              // Find the first failed test case to show as an example
              const failedTest = testResults.find((test) => !test.passed);

              setModalMessage({
                summary: `${passed}/${total} passed, ${total - passed}/${total} failing`,
                failedExample: failedTest
                  ? {
                      input: failedTest.input ? JSON.parse(failedTest.input) : null,
                      expected: failedTest.expectedOutput,
                      actual: failedTest.actualOutput
                    }
                  : null
              });

              setModalSuccess(false);
            }
            setShowModal(true);
          }

          setIsSubmitting(false);

          // Manually trigger editor layout after results are displayed
          if (editorRef.current) {
            setTimeout(() => {
              editorRef.current.layout();
            }, 200);
          }
        }, 100);
      } catch (err) {
        console.error(`Error ${isPreview ? "running" : "submitting"} code:`, err);
        const errorMessage =
          err.response?.data?.error || err.message || `Failed to ${isPreview ? "run" : "submit"} code`;

        if (isPreview) {
          setError(errorMessage);
        } else {
          // For submission errors, show in modal
          setModalMessage({
            summary: `Error: ${errorMessage}`,
            failedExample: null
          });
          setModalSuccess(false);
          setShowModal(true);
        }

        setIsSubmitting(false);
      }
    }
  };

  // Function to run code with preview mode (only runs a subset of test cases)
  const handleRun = () => {
    handleCodeExecution(true);
    // Reset to show the first test case tab when running
    setActiveTestTab(0);
  };

  // Function to submit code (runs all test cases)
  const handleSubmit = () => handleCodeExecution(false);

  useEffect(() => {
    // Fetch the initial code template when the component mounts
    fetchLanguageStub(currentLanguage);
  }, [currentLanguage]);

  useEffect(() => {
    // Fetch preview test cases when problemId changes
    if (problemId) {
      fetchPreviewTestCases();
    }
  }, [problemId]);

  // When a test case completes, auto-navigate to it if it failed
  useEffect(() => {
    // Find the first failed test case
    const failedTestIndex = Object.entries(testResults).find(([_, result]) => result && !result.passed)?.[0];

    if (failedTestIndex !== undefined) {
      setActiveTestTab(Number(failedTestIndex));
    }
  }, [testResults]);

  // Function to handle click outside modal
  const handleClickOutside = (e) => {
    if (e.target.classList.contains("result-modal-overlay")) {
      setShowModal(false);
    }
  };

  // Handle errors during code submission
  useEffect(() => {
    if (error) {
      setModalMessage({
        summary: `Error: ${error}`,
        failedExample: null
      });
      setModalSuccess(false);
      setShowModal(true);
    }
  }, [error]);

  // Handle layout updates when results or errors change
  useEffect(() => {
    if (editorRef.current) {
      // Manually trigger layout update when results or errors change
      setTimeout(() => {
        editorRef.current.layout();
      }, 300);
    }
  }, [results, error, testResults]);

  // Monaco editor options
  const options = {
    selectOnLineNumbers: true,
    roundedSelection: false,
    readOnly: false,
    cursorStyle: "line",
    // Disable automaticLayout - we'll handle layout manually to avoid ResizeObserver issues
    automaticLayout: false,
    minimap: { enabled: false },
    scrollBeyondLastLine: false,
    lineNumbers: "on",
    wordWrap: "on",
    folding: true,
    lineDecorationsWidth: 7,
    fontSize: 18,
    fontFamily: "'Fira Code', 'Consolas', 'Monaco', monospace",
    suggestFontSize: 18,
    bracketPairColorization: { enabled: true },
    autoIndent: "full",
    tabSize: 4
  };

  // Editor before mount callback
  const handleEditorBeforeMount = (monaco) => {
    // Define custom editor themes
    monaco.editor.defineTheme("leetCodeDark", {
      base: "vs-dark",
      inherit: true,
      rules: [],
      colors: {
        "editor.background": "#1e1e1e",
        "editor.lineHighlightBackground": "#2d2d2d",
        "editorLineNumber.foreground": "#858585",
        "editorLineNumber.activeForeground": "#c6c6c6",
        "editor.selectionBackground": "#264f78"
      }
    });
  };

  return (
    <div className="code-editor">
      <div className="editor-toolbar">
        <select value={currentLanguage} onChange={handleLanguageChange} className="language-selector">
          <option value="javascript">JavaScript</option>
          <option value="python">Python</option>
          <option value="cpp">C++</option>
        </select>
        <div className="editor-info">
          <span className="editor-status">{isEditorReady ? "Editor Ready" : "Loading..."}</span>
        </div>
      </div>

      <PanelGroup direction="vertical" className="editor-panels">
        <Panel defaultSize={showTestPanel ? 70 : 100} minSize={30} className="monaco-editor-panel">
          <div className="editor-container">
            <Editor
              height="100%"
              defaultLanguage={getMonacoLanguage(currentLanguage)}
              language={getMonacoLanguage(currentLanguage)}
              value={codeValue}
              theme="leetCodeDark"
              options={options}
              onMount={handleEditorDidMount}
              beforeMount={handleEditorBeforeMount}
              onChange={(newValue) => {
                setCodeValue(newValue);
              }}
              loading={<div className="editor-loading">Loading editor...</div>}
            />
          </div>
        </Panel>

        {showTestPanel && (
          <>
            <PanelResizeHandle className="resize-handle horizontal" />

            <Panel defaultSize={30} minSize={15} className="test-panel">
              <div className="test-results-container">
                {/* For Run button results - Tabbed Interface */}
                {previewTestCases.length > 0 && (
                  <div className="preview-test-tabs">
                    <div className="tab-headers">
                      {previewTestCases.map((_, index) => (
                        <div
                          key={index}
                          className={`tab-header ${activeTestTab === index ? "active" : ""} ${
                            testResults[index] ? (testResults[index].passed ? "passed" : "failed") : ""
                          }`}
                          onClick={() => setActiveTestTab(index)}
                        >
                          <span className="tab-number">Test #{index + 1}</span>
                          {testResults[index] && (
                            <span className="tab-status">
                              {testResults[index].passed ? (
                                <span className="status-icon pass">✓</span>
                              ) : (
                                <span className="status-icon fail">✗</span>
                              )}
                            </span>
                          )}
                        </div>
                      ))}
                    </div>

                    <div className="tab-content">
                      {previewTestCases.map((testCase, index) => (
                        <div key={index} className={`tab-panel ${activeTestTab === index ? "active" : ""}`}>
                          <div className="test-case-info">
                            <div className="test-input">
                              <h4>Input:</h4>
                              <div className="formatted-params">
                                {formatTestInput(testCase.input).map((param, paramIndex) => (
                                  <div key={paramIndex} className="param-item">
                                    <div className="param-header">
                                      <div className="param-name">{param.name}</div>
                                      <div className="param-type output-type">{param.type}</div>
                                    </div>
                                    <div className="param-value">
                                      {param.type === "array" ? (
                                        <CollapsibleArray array={param.rawValue} />
                                      ) : param.type === "object" ? (
                                        <pre>{JSON.stringify(param.rawValue, null, 2)}</pre>
                                      ) : param.type === "string" ? (
                                        <pre>"{param.rawValue}"</pre>
                                      ) : (
                                        <pre>{String(param.rawValue)}</pre>
                                      )}
                                    </div>
                                  </div>
                                ))}
                              </div>
                            </div>

                            <div className="test-expected">
                              <h4>Expected Output:</h4>
                              <FormattedOutput value={testCase.expected} />
                            </div>

                            {!testResults[index] && !isSubmitting && (
                              <div className="no-results-message">
                                <p>Click the "Run" button to test your code against this test case.</p>
                              </div>
                            )}

                            {isSubmitting && activeTestTab === index && (
                              <div className="test-loading">
                                <div className="loader"></div>
                                <span>Running test case...</span>
                              </div>
                            )}

                            {testResults[index] && (
                              <div className={`test-result ${testResults[index].passed ? "passed" : "failed"}`}>
                                <h4>Result: {testResults[index].passed ? "PASSED" : "FAILED"}</h4>

                                {testResults[index].actualOutput && (
                                  <div className="test-actual">
                                    <h4>Your Output:</h4>
                                    <FormattedOutput value={testResults[index].actualOutput} />
                                  </div>
                                )}

                                {testResults[index].error && (
                                  <div className="test-error">
                                    <h4>Error:</h4>
                                    <pre>{testResults[index].error}</pre>
                                  </div>
                                )}
                              </div>
                            )}
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>
                )}

                {/* For Submit button results - Show errors only */}
                {error && (
                  <div className="test-results error">
                    <h3>Error</h3>
                    <div className="error-message">{error}</div>
                  </div>
                )}
              </div>
            </Panel>
          </>
        )}
      </PanelGroup>

      <div className="editor-footer">
        <div className="footer-left">
          <button className="toggle-panel-button" onClick={() => setShowTestPanel(!showTestPanel)}>
            {showTestPanel ? "Hide Test Panel" : "Show Test Panel"}
          </button>
        </div>
        <div className="footer-right">
          <button className="run-button" onClick={handleRun} disabled={isSubmitting || !isEditorReady}>
            {isSubmitting ? "Running..." : "Run"}
          </button>
          <button className="submit-button" onClick={handleSubmit} disabled={isSubmitting || !isEditorReady}>
            {isSubmitting ? "Submitting..." : "Submit"}
          </button>
        </div>
      </div>

      {/* Modal for submission results */}
      {showModal && (
        <div className="result-modal-overlay" onClick={handleClickOutside}>
          <div className={`result-modal ${modalSuccess ? "success" : "failure"}`}>
            <div className="modal-content">
              <div className="modal-icon">{modalSuccess ? "✅" : "❌"}</div>

              {typeof modalMessage === "string" ? (
                <div className="modal-message">{modalMessage}</div>
              ) : (
                <>
                  <div className="modal-message">{modalMessage.summary}</div>

                  {/* Display Failed Test Example */}
                  {modalMessage.failedExample && (
                    <div className="modal-test-example">
                      <h4>Failed Test Case Example:</h4>

                      {/* Input */}
                      {modalMessage.failedExample.input && (
                        <div className="modal-test-input">
                          <div className="modal-test-label">Input:</div>
                          <div className="modal-test-data">
                            {Object.entries(modalMessage.failedExample.input).map(([key, val]) => {
                              const value = val.value !== undefined ? val.value : val;
                              const formattedValue = Array.isArray(value)
                                ? formatSingleLineArray(value)
                                : typeof value === "object" && value !== null
                                ? JSON.stringify(value)
                                : typeof value === "string"
                                ? `"${value}"`
                                : String(value);

                              return (
                                <div key={key} className="modal-param">
                                  <span className="modal-param-name">{key}:</span>
                                  <span className="modal-param-value">
                                    {Array.isArray(value) ? (
                                      <CollapsibleArray array={value} />
                                    ) : typeof value === "object" && value !== null ? (
                                      <pre>{JSON.stringify(value, null, 2)}</pre>
                                    ) : typeof value === "string" ? (
                                      <pre>"{value}"</pre>
                                    ) : (
                                      <pre>{formattedValue}</pre>
                                    )}
                                  </span>
                                </div>
                              );
                            })}
                          </div>
                        </div>
                      )}

                      {/* Expected Output */}
                      <div className="modal-test-expected">
                        <div className="modal-test-label">Expected:</div>
                        <div className="modal-test-data">
                          {(() => {
                            const value = formatOutput(modalMessage.failedExample.expected);
                            return Array.isArray(value) ? (
                              <CollapsibleArray array={value} />
                            ) : typeof value === "object" && value !== null ? (
                              <pre>{JSON.stringify(value, null, 2)}</pre>
                            ) : typeof value === "string" ? (
                              <pre>"{value}"</pre>
                            ) : (
                              <pre>{String(value)}</pre>
                            );
                          })()}
                        </div>
                      </div>

                      {/* Actual Output */}
                      <div className="modal-test-actual">
                        <div className="modal-test-label">Your Output:</div>
                        <div className="modal-test-data">
                          {(() => {
                            const value = formatOutput(modalMessage.failedExample.actual);
                            return Array.isArray(value) ? (
                              <CollapsibleArray array={value} />
                            ) : typeof value === "object" && value !== null ? (
                              <pre>{JSON.stringify(value, null, 2)}</pre>
                            ) : typeof value === "string" ? (
                              <pre>"{value}"</pre>
                            ) : (
                              <pre>{String(value)}</pre>
                            );
                          })()}
                        </div>
                      </div>
                    </div>
                  )}
                </>
              )}

              <button className="modal-close-button" onClick={() => setShowModal(false)}>
                Close
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default CodeEditor;
