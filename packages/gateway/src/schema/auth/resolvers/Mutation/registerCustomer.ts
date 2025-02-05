import { GraphQLError } from "graphql/error";
import { authSchema } from "~/adapter/in/dto/auth";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const registerCustomer: NonNullable<MutationResolvers["registerCustomer"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session !== null) {
    throw new GraphQLError("Authenticated");
  }

  const data = await invokeService("auth", {
    endpoint: "POST /register",
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: authSchema,
    input: {
      username: _arg.username,
      password: _arg.password,
      role: "CUSTOMER",
    },
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
