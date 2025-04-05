import { creditBalanceSchema } from "~/adapter/in/dto/credit-balance";
import { invokeService } from "~/infrastructure/service";
import type { QueryResolvers } from "~/schema/types.generated";


export const creditBalance: NonNullable<QueryResolvers["creditBalance"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "CUSTOMER") {
    throw new Error("Not authorized");
  }

  const data = await invokeService("accounting", {
    endpoint: `GET /customers/${_ctx.session.id}/credit-balance`,
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: creditBalanceSchema,
  });

  return data.creditBalance;
};
