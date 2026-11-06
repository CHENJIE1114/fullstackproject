import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  return {
    plugins: [vue()],
    server: {
      host: env.VITE_HOST || '127.0.0.1',
      port: Number(env.VITE_PORT || '5188'),
      strictPort: true,
      proxy: {
        '/api': env.VITE_API_PROXY_TARGET || 'http://127.0.0.1:8091',
        '/ai': env.VITE_AI_PROXY_TARGET || 'http://127.0.0.1:3000',
      },
    },
  }
})
