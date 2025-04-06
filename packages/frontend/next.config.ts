import type { NextConfig } from "next";


const nextConfig: NextConfig = {
  poweredByHeader: false,
  output: process.env.BUILD_STANDALONE === "true" ? "standalone" : undefined,
};

export default nextConfig;
