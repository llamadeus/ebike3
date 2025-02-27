import { defineConfig } from "@eddeee888/gcg-typescript-resolver-files";
import type { CodegenConfig } from "@graphql-codegen/cli";


export default {
  schema: "src/**/*.graphql",
  generates: {
    "src/schema": defineConfig({
      mergeSchema: false,
      typesPluginsConfig: {
        contextType: "../infrastructure/server#ResolverContext",
      },
    }),
  },
} satisfies CodegenConfig;
