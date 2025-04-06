import { authExchange } from "@urql/exchange-auth";
import { registerUrql } from "@urql/next/rsc";
import { cookies } from "next/headers";
import { cacheExchange, createClient, fetchExchange } from "urql";
import { isNullish } from "~/utils/value";


/**
 * Creates a new urql client.
 *
 * @returns The urql client.
 */
export function makeClient() {
  return createClient({
    url: process.env.GRAPHQL_ENDPOINT_INTERNAL ?? "http://localhost:4000/graphql",
    exchanges: [
      cacheExchange,
      authExchange(async (utils) => {
        const cookieStore = await cookies();
        const jwt = cookieStore.get("jwt")?.value ?? null;

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
    fetchOptions: { cache: "no-store" },
  });
}

export const { getClient } = registerUrql(makeClient);
