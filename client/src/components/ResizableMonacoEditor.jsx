import { useEffect, useRef } from 'react';
import './CodeEditor.css';
import Editor from '@monaco-editor/react';

export const ResizableMonacoEditor = ({
  currentLanguage,
  codeValue,
  //   handleEditorDidMount,
  setCodeValue,
}) => {
  const containerRef = useRef(null);
  const editorRef = useRef(null);

  const options = {
    selectOnLineNumbers: true,
    roundedSelection: false,
    readOnly: false,
    cursorStyle: 'line',
    automaticLayout: false,
    minimap: { enabled: false },
    scrollBeyondLastLine: false,
    lineNumbers: 'on',
    wordWrap: 'on',
    folding: true,
    lineDecorationsWidth: 7,
    fontSize: 18,
    fontFamily: "'Fira Code', 'Consolas', 'Monaco', monospace",
    suggestFontSize: 18,
    bracketPairColorization: { enabled: true },
    autoIndent: 'full',
    tabSize: 4,
  };

  const handleEditorBeforeMount = (monaco) => {
    monaco.editor.defineTheme('leetCodeDark', {
      base: 'vs-dark',
      inherit: true,
      rules: [],
      colors: {
        'editor.background': '#1e1e1e',
        'editor.lineHighlightBackground': '#2d2d2d',
        'editorLineNumber.foreground': '#858585',
        'editorLineNumber.activeForeground': '#c6c6c6',
        'editor.selectionBackground': '#264f78',
      },
    });
  };

  const handleEditorDidMount = (editor) => {
    editorRef.current = editor;

    setTimeout(() => {
      editor.layout();
    }, 0);
  };

  useEffect(() => {
    if (!containerRef.current) return;

    const observer = new ResizeObserver(() => {
      if (editorRef.current) {
        editorRef.current.layout();
      }
    });

    observer.observe(containerRef.current);

    return () => observer.disconnect();
  }, []);

  return (
    <div className="editor-container" ref={containerRef}>
      <Editor
        height="100%"
        width="100%"
        defaultLanguage={'javascript'}
        language={currentLanguage}
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
  );
};
