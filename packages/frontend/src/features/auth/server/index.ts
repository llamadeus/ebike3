import { graphql } from "~/gql";
import { getClient } from "~/urql/server";


const authQueryDocument = graphql(`
  query ServerAuth {
    auth {
      id
      role
      username
      lastLogin
    }
  }
`);

/**
 * Queries the current user from the server.
 */
export async function queryAuth() {
  return getClient().query(authQueryDocument, {});
}
