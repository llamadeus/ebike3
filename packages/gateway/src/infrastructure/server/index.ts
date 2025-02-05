import { useCookies } from "@whatwg-node/server-plugin-cookies";
import { createSchema, createYoga, type YogaServerInstance } from "graphql-yoga";
import { nanoid } from "nanoid";
import { authSchema } from "~/adapter/in/dto/auth";
import type { Session, SessionService } from "~/domain/service/session";
import { invokeService } from "~/infrastructure/service";
import { resolvers } from "~/schema/resolvers.generated";
import { typeDefs } from "~/schema/typeDefs.generated";


/**
 * The context for the resolvers.
 */
export interface ResolverContext {
  /**
   * The request object.
   */
  request: Request;
  /**
   * The session of the user.
   */
  session: Session | null;
  /**
   * The request ID.
   */
  requestId: string;
  /**
   * The session service.
   */
  sessionService: SessionService;
}

interface Options {
  /**
   * The session service.
   */
  sessionService: SessionService;
}

/**
 * Gets the JWT from the request headers.
 * The JWT is expected to be in the format "Bearer <token>".
 * If the JWT is not found, it will be cheked for in the cookie.
 *
 * @param request The request object.
 * @returns The JWT from the request headers.
 */
async function getJWTFromRequest(request: Request): Promise<string | null> {
  const authorization = request.headers.get("authorization");
  if (authorization !== null) {
    const [type, token] = authorization.split(" ");
    if (typeof token == "undefined") {
      return null;
    }

    if (type?.toLowerCase() === "bearer") {
      return token;
    }
  }

  const cookieValue = await request.cookieStore?.get("jwt");

  return cookieValue?.value ?? null;
}

/**
 * Creates a new GraphQL server.
 *
 * @param options The options for the server.
 * @returns The GraphQL server.
 */
export function makeYogaServer(options: Options): YogaServerInstance<ResolverContext, Record<string, any>> {
  return createYoga<ResolverContext>({
    schema: createSchema<ResolverContext>({ typeDefs, resolvers }),
    graphiql: process.env.NODE_ENV === "development",
    landingPage: false,
    context: async (initialContext) => {
      const requestId = nanoid(24);
      const jwt = await getJWTFromRequest(initialContext.request);
      let session: Session | null = null;

      if (jwt !== null) {
        try {
          session = options.sessionService.verify(jwt);

          const auth = await invokeService("auth", {
            endpoint: "GET /auth",
            headers: {
              "X-Request-ID": requestId,
              "X-Session-ID": session.sessionId,
            },
            output: authSchema.nullable(),
          });

          if (auth === null) {
            options.sessionService.destroySession(initialContext.request);

            session = null;
          }
        }
        catch (err) {
          console.log(err);
        }
      }

      return {
        requestId,
        session,
        request: initialContext.request,
        sessionService: options.sessionService,
      };
    },
    plugins: [useCookies()],
  });
}
