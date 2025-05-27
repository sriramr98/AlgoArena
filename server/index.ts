import express, { Request, Response } from "express";
import cors from "cors";
import fs from "fs";
import codeJudge from "./judge"; // Original executor
import { problems } from "./problems";

const app = express();
const PORT = process.env.PORT || 5000;

// Middleware
app.use(express.json());
app.use(cors());

// Routes
app.get("/api/problems", (_req: Request, res: Response) => {
  // Only send basic info for the problems list
  const problemsList = problems.map(({ id, title, difficulty }) => ({
    id,
    title,
    difficulty,
  }));
  res.json(problemsList);
});

app.get("/api/problems/:id", (req: Request, res: Response) => {
  const problem = problems.find((p) => p.id === req.params.id);
  if (!problem) {
    return res.status(404).json({ error: "Problem not found" });
  }
  res.json(problem);
});

app.get("/api/problems/:id/stub", (req: Request, res: Response) => {
  // Get the problem stub from the code_templates folder and return it. The language is passed as a query parameter
  const { id } = req.params;
  const { language } = req.query;
  if (!language) {
    return res.status(400).json({ error: "Language is required" });
  }
  const problem = problems.find((p) => p.id === id);
  if (!problem) {
    return res.status(404).json({ error: "Problem not found" });
  }

  try {
    const codeTemplate = fs.readFileSync(
      `./code_templates/${problem.id}/stub/${language}`,
      "utf8",
    );
    res.json({ codeTemplate });
  } catch (error) {
    console.error(error);
    return res
      .status(400)
      .json({
        error: `Language ${language} is not supported for this problem`,
      });
  }
});

app.get("/api/problems/:id/testcases", (req: Request, res: Response) => {
  const problem = problems.find((p) => p.id === req.params.id);
  if (!problem) {
    return res.status(404).json({ error: "Problem not found" });
  }
  res.json(problem.testCases.slice(0, 2)); // Return only the first two test cases for preview
});

interface SubmitRequest {
  code: string;
  language: string;
  problemId: string;
}

// Code submission and evaluation endpoint
app.post("/api/submit", async (req: Request, res: Response) => {
  try {
    const { code, language, problemId } = req.body as SubmitRequest;
    // preview denotes whether only first few test cases should be run
    const preview = req.query.preview === "true";

    // Validate inputs
    if (!code || !language || !problemId) {
      return res
        .status(400)
        .json({
          error: "Missing required fields: code, language, or problemId",
        });
    }

    // Find the problem
    const problem = problems.find((p) => p.id === problemId);
    if (!problem) {
      return res.status(404).json({ error: "Problem not found" });
    }

    // Execute the code
    const results = await codeJudge(code, problem, language, preview ? 2 : problem.testCases.length);

    // Return results
    res.json(results);
  } catch (error: any) {
    console.error("Submission error:", error);
    res
      .status(500)
      .json({
        error: error.message || "An error occurred processing your submission",
      });
  }
});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});