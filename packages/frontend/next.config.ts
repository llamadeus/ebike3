import type { NextConfig } from "next";

export default {
  output: process.env["BUILD_STANDALONE"] === "true" ? "standalone" : "export",
} satisfies NextConfig;
