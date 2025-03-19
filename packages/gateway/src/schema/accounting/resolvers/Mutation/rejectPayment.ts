import { GraphQLError } from "graphql/error";
import { paymentSchema } from "~/adapter/in/dto/payment";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const rejectPayment: NonNullable<MutationResolvers["rejectPayment"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "ADMIN") {
    throw new GraphQLError("Not authorized");
  }

  const data = await invokeService("accounting", {
    endpoint: "PATCH /payments/{id}",
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    input: {
      status: "REJECTED",
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
