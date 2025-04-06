import { customerSchema } from "~/adapter/in/dto/customer";
import { invokeService } from "~/infrastructure/service";
import type { QueryResolvers } from "~/schema/types.generated";


export const position: NonNullable<QueryResolvers["position"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "CUSTOMER") {
    throw new Error("Not authorized");
  }

  const data = await invokeService("customers", {
    endpoint: `GET /customers/${_ctx.session.id}`,
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: customerSchema,
  });

  return {
    x: data.positionX,
    y: data.positionY,
  };
};
