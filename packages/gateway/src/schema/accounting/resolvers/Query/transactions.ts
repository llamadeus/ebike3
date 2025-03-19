import { GraphQLError } from "graphql/error";
import { z } from "zod";
import { expenseSchema } from "~/adapter/in/dto/expense";
import { paymentSchema } from "~/adapter/in/dto/payment";
import { invokeService } from "~/infrastructure/service";
import type { QueryResolvers, ResolversTypes, ResolversUnionTypes } from "~/schema/types.generated";


export const transactions: NonNullable<QueryResolvers["transactions"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "CUSTOMER") {
    throw new GraphQLError("Not authorized");
  }

  const payments = await invokeService("accounting", {
    endpoint: `GET /customers/${_ctx.session.id}/payments`,
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: z.array(paymentSchema),
  });
  const expenses = await invokeService("accounting", {
    endpoint: `GET /customers/${_ctx.session.id}/expenses`,
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: z.array(expenseSchema),
  });

  const result: ResolversUnionTypes<ResolversTypes>["Transaction"][] = [];

  for (const payment of payments) {
    result.push({
      __typename: "Payment",
      id: payment.id,
      amount: payment.amount,
      status: payment.status,
      createdAt: payment.createdAt,
      customer: null,
    });
  }

  for (const expense of expenses) {
    result.push({
      __typename: "Expense",
      id: expense.id,
      amount: expense.amount,
      createdAt: expense.createdAt,
      customer: null,
    });
  }

  return result;
};
