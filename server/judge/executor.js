const axios = require('axios');

const getSupportedVersion = (lang) => {
  const supportedVersions = {
    python: '3.12.0',
    javascript: '20.11.1',
  };
  return supportedVersions[lang] || null;
};

const getLangExtension = (lang) => {
  const langExtensions = {
    python: 'py',
    javascript: 'js',
  };
  return langExtensions[lang] || null;
};

const execute = async (code, language) => {
  const PISTON_HOST = process.env.PISTON_API_URL || 'http://localhost:2000';
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
    ...data,
  };
};

module.exports = execute;
