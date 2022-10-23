import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'

const path = require("path");
const toPath = (filePath: string) => path.join(process.cwd(), filePath);

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@wailsapp/runtime': toPath("wailsjs/runtime/runtime.d.ts")
    }
  }
})
