import { defineConfig } from "vite"
import reactRefresh from "@vitejs/plugin-react-refresh"
import { resolve } from "path"

const rendererPath = resolve(__dirname, "./src/renderer")
const outDirRenderer = resolve(__dirname, "./app/renderer")

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [reactRefresh()],
  base: "./",
  root: rendererPath,
  build: {
    outDir: outDirRenderer,
    emptyOutDir: true,
  },
  resolve: {
    alias: [
      {
        find: "@renderer",
        replacement: resolve(__dirname, "src/renderer"),
      },
      {
        find: "@preload",
        replacement: resolve(__dirname, "src/preload"),
      },
      {
        find: "@common",
        replacement: resolve(__dirname, "src/common"),
      },
    ],
  },
})
