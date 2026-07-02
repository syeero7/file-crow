import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import path from "path";

const PORT = 8080;

export default defineConfig(({ mode }) => {
  return {
    plugins: [svelte()],
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
    ...(mode.startsWith("dev") && {
      server: {
        proxy: {
          "/api": {
            target: `http://localhost:${PORT}`,
            changeOrigin: true,
            secure: false,
            rewrite: (path) => path.replace(/^\/api/, ""),
          },
        },
      },
    }),
  };
});
