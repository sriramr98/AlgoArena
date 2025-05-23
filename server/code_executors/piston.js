const axios = require('axios');

const getSupportedVersion = (lang) => {
    const supportedVersions = {
        'python': '3.12.0',
        'javascript': '20.11.1',
    };
    return supportedVersions[lang] || null; 
}

const getLangExtension = (lang) => {
    const langExtensions = {
        'python': 'py',
        'javascript': 'js',
    };
    return langExtensions[lang] || null; 
}

const getRuntimes = async () => {
   const PISTON_HOST = process.env.PISTON_API_URL || 'http://localhost:2000';
    const RUNTIME_URL = `${PISTON_HOST}/api/v2/runtimes`;

    try {
        const response = await axios.get(RUNTIME_URL);
        return response.data;
    }
    catch (error) {
        console.error('Error fetching runtimes:', error);
        throw new Error('Unable to fetch runtimes');
    }
    
}


const callPiston = async (lang, code, stdin = "", args = []) => {
    const PISTON_HOST = process.env.PISTON_API_URL || 'http://localhost:2000';
    const EXECUTE_URL = `${PISTON_HOST}/api/v2/execute`;

    const payload = {
        language: lang,
        version: getSupportedVersion(lang),
        files: [
            {
                name: `code.${getLangExtension(lang)}`,
                content: code
            }
        ],
        stdin: stdin,
        args: args,
        compile_timeout: 10000,
        run_timeout: 3000,
        compile_cpu_time: 10000,
        run_cpu_time: 3000,
        compile_memory_limit: -1,
        run_memory_limit: -1
    }

    return axios.post(EXECUTE_URL, payload)
}

module.exports = callPiston