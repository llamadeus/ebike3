import type { CodegenConfig } from "@graphql-codegen/cli";


export default {
  schema: "../gateway/**/*.graphql",
  documents: ["src/**/*.ts", "src/**/*.tsx"],
  ignoreNoDocuments: true,
  generates: {
    "./src/gql/": {
      preset: "client",
    },
  },
} satisfies CodegenConfig;
