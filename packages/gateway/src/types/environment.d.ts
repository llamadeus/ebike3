declare global {
  namespace NodeJS {
    interface ProcessEnv {
      NODE_ENV: "development" | "production";
      CI?: boolean;
      npm_package_version: string;

      // Values from `.env`
      APP_PORT: "true" | string;
      JWT_PRIVATE_KEY_PATH: string;
      JWT_PUBLIC_KEY_PATH: string;
    }
  }
}

// If this file has no import/export statements (i.e. is a script)
// convert it into a module by adding an empty export statement.
export {};
