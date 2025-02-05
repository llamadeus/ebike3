import { readFile } from "fs/promises";
import { SessionService } from "~/domain/service/session";
import { makeYogaServer } from "~/infrastructure/server";


const port = typeof process.env.APP_PORT != "undefined" ? Number(process.env.APP_PORT) : 4000;

const jwtPrivateKey = await readFile(process.env.JWT_PRIVATE_KEY_PATH, "utf-8");
const jwtPublicKey = await readFile(process.env.JWT_PUBLIC_KEY_PATH, "utf-8");
const sessionService = new SessionService(jwtPrivateKey, jwtPublicKey);

const yoga = makeYogaServer({
  sessionService,
});

const server = Bun.serve({
  port,
  fetch: yoga as never,
});

console.info(`🚀 Server is running on ${new URL(yoga.graphqlEndpoint, `http://${server.hostname}:${server.port}`)}`);
