import { defineConfig } from "vite";
import tailwindcss from "@tailwindcss/vite";
import { ViteMinifyPlugin } from "vite-plugin-minify";

const PORT = 8080;

export default defineConfig(({ mode }) => {
  return {
    plugins: [tailwindcss(), ViteMinifyPlugin()],
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
