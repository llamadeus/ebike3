import { GraphQLError } from "graphql/error";
import { z } from "zod";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const logout: NonNullable<MutationResolvers['logout']> = async (_parent, _arg, _ctx) => {
  if (_ctx.session === null) {
    throw new GraphQLError("Not authenticated");
  }

  await invokeService("auth", {
    endpoint: "POST /logout",
    headers: {
      "X-Request-ID": _ctx.requestId,
      "X-Session-ID": _ctx.session.sessionId,
    },
    output: z.object({
      status: z.string(),
    }),
  });

  // Logout the user
  _ctx.sessionService.destroySession(_ctx.request);

  return true;
};
