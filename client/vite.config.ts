import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

const serverUrl =
  process.env.BACKEND_API_URL === undefined ||
  process.env.BACKEND_API_URL === ''
    ? 'http://localhost:8080'
    : process.env.BACKEND_API_URL;

export default defineConfig({
  plugins: [react()],
  root: '.',
  server: {
    port: 3000,
    host: true, // needed for Docker container
    proxy: {
      '/api': {
        target: serverUrl,
        changeOrigin: true,
      },
    },
  },
});
