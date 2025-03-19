import { GraphQLError } from "graphql/error";
import { z } from "zod";
import { paymentSchema } from "~/adapter/in/dto/payment.ts";
import { invokeService } from "~/infrastructure/service";
import type { QueryResolvers } from "~/schema/types.generated";


export const payments: NonNullable<QueryResolvers["payments"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "ADMIN") {
    throw new GraphQLError("Not authorized");
  }

  const payments = await invokeService("accounting", {
    endpoint: `GET /payments`,
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: z.array(paymentSchema),
  });

  return payments.map(payment => {
    return {
      id: payment.id,
      amount: payment.amount,
      status: payment.status,
      createdAt: payment.createdAt,
      customer: null,
    };
  });
};
