import { GraphQLError } from "graphql/error";
import { userSchema } from "~/adapter/in/dto/user";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const login: NonNullable<MutationResolvers["login"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session !== null) {
    throw new GraphQLError("Authenticated");
  }

  const data = await invokeService("auth", {
    endpoint: "POST /login",
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    input: {
      username: _arg.username,
      password: _arg.password,
    },
    output: userSchema,
  });

  // Login the user
  _ctx.sessionService.createSession(_ctx.request, data);

  return {
    id: data.id,
    username: data.username,
    role: data.role,
    lastLogin: data.lastLogin,
  };
};
