import svelteConfig from "./svelte.config.js";
import { defineConfig } from "eslint/config";
import globals from "globals";
import ts from "typescript-eslint";
import svelte from "eslint-plugin-svelte";

export default defineConfig(
  ts.configs.recommended,
  ts.configs.recommended,
  svelte.configs.recommended,
  {
    languageOptions: {
      globals: {
        ...globals.browser,
      },
    },
  },
  {
    files: ["**/*.svelte", "**/*.svelte.ts", "**/*.svelte.js"],
    languageOptions: {
      parserOptions: {
        projectService: true,
        extraFileExtensions: [".svelte"],
        parser: ts.parser,
        // Note: `eslint --cache` will fail with non-serializable properties.
        // In those cases, please remove the non-serializable properties.
        // svelteConfig: {
        //   ...svelteConfig,
        //   }
        svelteConfig,
      },
    },
  },
  {
    rules: {},
  },
);
