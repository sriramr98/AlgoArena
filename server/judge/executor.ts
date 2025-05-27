import axios from "axios";

export interface ExecutionResult {
  success: boolean;
  message?: string;
  stdout?: string;
  stderr?: string;
  output?: string;
  code?: number;
  signal?: string | null;
}

const getSupportedVersion = (lang: string): string | null => {
  const supportedVersions: Record<string, string> = {
    python: "3.12.0",
    javascript: "20.11.1",
  };
  return supportedVersions[lang] || null;
}

const getLangExtension = (lang: string): string | null => {
  const langExtensions: Record<string, string> = {
    python: "py",
    javascript: "js",
  };
  return langExtensions[lang] || null;
}

const execute = async (code: string, language: string): Promise<ExecutionResult> => {
  const PISTON_HOST = process.env.PISTON_API_URL || "http://localhost:2000";
  const EXECUTE_URL = `${PISTON_HOST}/api/v2/execute`;

  const payload = {
    language,
    version: getSupportedVersion(language),
    files: [
      {
        name: `code.${getLangExtension(language)}`,
        content: code,
      },
    ],
    compile_timeout: 10000,
    run_timeout: 3000,
    compile_cpu_time: 10000,
    run_cpu_time: 3000,
    compile_memory_limit: -1,
    run_memory_limit: -1,
  };

  try {
    const resp = await axios.post(EXECUTE_URL, payload);
    if (resp.status !== 200) {
      return {
        success: false,
        message: resp.data?.message || 'Unable to execute code',
      };
    }

    const data = resp.data?.run || {};
    return {
      success: true,
      ...data
    };
  } catch (error: any) {
    return {
      success: false,
      message: error.message || 'Unable to execute code',
    };
  }
}

export default execute;