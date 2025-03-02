import { GraphQLError } from "graphql/error";
import { z } from "zod";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const updateCustomerPosition: NonNullable<MutationResolvers["updateCustomerPosition"]> = async (
  _parent,
  _arg,
  _ctx,
) => {
  if (_ctx.session?.role !== "CUSTOMER") {
    throw new GraphQLError("Not authorized");
  }

  await invokeService("customers", {
    endpoint: `PATCH /customers/${_ctx.session.id}/position`,
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    input: {
      positionX: _arg.position.x,
      positionY: _arg.position.y,
    },
    output: z.null(),
  });

  return true;
};
