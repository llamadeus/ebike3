import { GraphQLError } from "graphql/error";
import { z } from "zod";
import { customerSchema } from "~/adapter/in/dto/customer";
import { invokeService } from "~/infrastructure/service";
import { isNullish } from "~/infrastructure/utils/helpers.ts";
import type { QueryResolvers } from "~/schema/types.generated";


export const customers: NonNullable<QueryResolvers["customers"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "ADMIN") {
    throw new GraphQLError("Not authorized");
  }

  const data = await invokeService("customers", {
    endpoint: "GET /customers",
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: z.array(customerSchema),
  });

  return data.map((customer) => ({
    id: customer.id,
    name: customer.name,
    position: {
      x: customer.positionX,
      y: customer.positionY,
    },
    creditBalance: customer.creditBalance,
    activeRental: ! isNullish(customer.activeRental)
      ? {
        id: customer.activeRental.id,
        vehicleId: customer.activeRental.vehicleId,
        customerId: customer.activeRental.customerId,
        vehicleType: customer.activeRental.vehicleType,
        start: customer.activeRental.start,
        cost: customer.activeRental.cost,
      }
      : null,
    lastLogin: customer.lastLogin,
  }));
};
