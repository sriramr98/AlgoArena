# LeetCode Clone

A simple LeetCode clone with a React frontend and Express.js backend.

## Features

- List of coding problems with difficulty indicators
- Problem detail page with description and constraints
- Advanced code editor with:
  - Syntax highlighting for JavaScript, Python, and C++
  - Monaco Editor (same as VS Code)
  - Language selector
  - Code submission capability
  - Smart code templates based on problem type
- Code execution with test cases:
  - Built-in executors for JavaScript, Python, and C++
  - Optional Judge0 CE integration for more languages and better isolation

## Project Structure

- `/client` - React frontend
- `/server` - Express.js backend
  - `/server/utils` - Code execution utilities

## Getting Started

### Prerequisites

- Node.js
- npm
- Docker and Docker Compose (for Judge0 CE integration)

### Installation

1. Clone the repository
2. Install dependencies:

```bash
# Install root dependencies
npm install

# Install server dependencies
cd server
npm install

# Install client dependencies
cd ../client
npm install
```

### Running the Application

```bash
# From the root directory
npm run dev
```

This will concurrently start:

- Frontend on http://localhost:3000
- Backend on http://localhost:5000

### Using Judge0 CE for Code Execution (Optional)

For improved code execution with better language support:

1. Start Judge0 CE using Docker:

```bash
cd server
docker-compose -f docker-compose.judge0.yml up -d
```

2. Check if Judge0 is running:

```bash
npm run check-judge0
```

3. Switch to the Judge0-based executor:

```bash
npm run toggle-executor
```

4. Restart the server:

```bash
npm run dev
```

## API Endpoints

- `GET /api/problems` - Get list of all problems
- `GET /api/problems/:id` - Get problem details by ID
- `POST /api/submit` - Submit code for evaluation

## Technologies Used

- Frontend: React, React Router, Axios, Monaco Editor
- Backend: Express.js, Node.js
- Code Execution: VM2 (JavaScript), PythonShell (Python), Judge0 CE (optional)

## Troubleshooting

For Monaco Editor issues, run the included fix script:

```bash
# From the project root
npm run fix-codemirror
```

Or manually fix it with these steps:

```bash
# Clean up node_modules
cd client
rm -rf node_modules

# Remove specific problematic packages
npm uninstall @codemirror/basic-setup @codemirror/commands

# Install only the unified package
npm install codemirror@latest

# Reinstall all dependencies
npm install
```

Then update your imports in CodeEditor.js to use only the unified package:
