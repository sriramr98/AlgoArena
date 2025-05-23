# Monaco Editor Example

This file provides examples of using Monaco Editor in a React application.

## Setup

To use Monaco Editor in a React project:

```bash
npm install @monaco-editor/react
```

## Basic Usage

```javascript
import React, { useRef } from 'react';
import Editor from '@monaco-editor/react';

function MonacoEditor() {
  const editorRef = useRef(null);

  function handleEditorDidMount(editor, monaco) {
    editorRef.current = editor; 
  }
  
  return (
    <Editor
      height="500px"
      defaultLanguage="javascript"
      defaultValue="// Some initial code"
      onMount={handleEditorDidMount}
    />
  );
}
```

## Features

Monaco Editor provides many features of VS Code:

- Syntax highlighting for 60+ languages
- IntelliSense and code completion
- Error and warning annotations
- Code folding
- Find and replace
- Custom themes
- Keyboard shortcuts

## Theme Customization

```javascript
// Before mounting the editor
function handleEditorBeforeMount(monaco) {
  monaco.editor.defineTheme('myCustomTheme', {
    base: 'vs-dark',
    inherit: true,
    rules: [],
    colors: {
      'editor.background': '#1e1e1e',
      'editor.lineHighlightBackground': '#2d2d2d',
    }
  });
}

// When rendering
<Editor
  theme="myCustomTheme"
  beforeMount={handleEditorBeforeMount}
  // other props...
/>
```
