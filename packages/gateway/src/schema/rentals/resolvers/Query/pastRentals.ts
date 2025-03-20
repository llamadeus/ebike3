import { GraphQLError } from "graphql/error";
import { z } from "zod";
import { rentalSchema } from "~/adapter/in/dto/rental";
import { invokeService } from "~/infrastructure/service";
import type { QueryResolvers } from "~/schema/types.generated";


export const pastRentals: NonNullable<QueryResolvers["pastRentals"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "CUSTOMER") {
    throw new GraphQLError("Not authorized");
  }

  const data = await invokeService("rentals", {
    endpoint: `GET /customers/${_ctx.session.id}/rentals/past`,
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: z.array(rentalSchema),
  });

  return data.map(rental => {
    return {
      id: rental.id,
      customerId: rental.customerId,
      vehicleId: rental.vehicleId,
      start: rental.start,
      end: rental.end,
      cost: rental.cost,
      createdAt: rental.createdAt,
      updatedAt: rental.updatedAt,
    };
  });
};
