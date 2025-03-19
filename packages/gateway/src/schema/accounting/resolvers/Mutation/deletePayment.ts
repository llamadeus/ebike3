import { GraphQLError } from "graphql/error";
import { paymentSchema } from "~/adapter/in/dto/payment";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const deletePayment: NonNullable<MutationResolvers["deletePayment"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "CUSTOMER") {
    throw new GraphQLError("Not authorized");
  }

  const data = await invokeService("accounting", {
    endpoint: "DELETE /payments/{id}",
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    input: {
      customerId: _ctx.session.id,
    },
    output: paymentSchema,
  });

  return {
    id: data.id,
    amount: data.amount,
    status: data.status,
    customer: null,
    createdAt: data.createdAt,
  };
};
