import { authExchange } from "@urql/exchange-auth";
import { cacheExchange, createClient, fetchExchange, SSRExchange } from "urql";
import { isNullish } from "~/utils/value";


/**
 * Creates a new urql client.
 *
 * @param ssr The SSR exchange to use.
 * @returns The urql client.
 */
export function makeClient(ssr: SSRExchange) {
  const url = typeof window == "undefined"
    ? process.env.GRAPHQL_ENDPOINT_INTERNAL
    : process.env.GRAPHQL_ENDPOINT_EXTERNAL;

  return createClient({
    url: url ?? "http://localhost:4000/graphql",
    exchanges: [
      cacheExchange,
      ssr,
      authExchange(async (utils) => {
        let jwt: string | null = null;

        if (typeof window == "undefined") {
          const { cookies } = await import("next/headers");
          const cookieStore = await cookies();

          jwt = cookieStore.get("jwt")?.value ?? null;
        }

        return {
          addAuthToOperation: (operation) => {
            if (isNullish(jwt)) {
              return operation;
            }

            return utils.appendHeaders(operation, {
              Authorization: `Bearer ${jwt}`,
            });
          },
          didAuthError: () => false,
          refreshAuth: () => Promise.resolve(),
        };
      }),
      fetchExchange,
    ],
    fetchOptions: { credentials: "include" },
    requestPolicy: "cache-and-network",
    suspense: true,
  });
}
