import { GraphQLError } from "graphql/error";
import { rentalSchema } from "~/adapter/in/dto/rental";
import { invokeService } from "~/infrastructure/service";
import type { QueryResolvers } from "~/schema/types.generated";


export const activeRental: NonNullable<QueryResolvers["activeRental"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "CUSTOMER") {
    throw new GraphQLError("Not authorized");
  }

  const data = await invokeService("rentals", {
    endpoint: `GET /customers/${_ctx.session.id}/rentals/active`,
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: rentalSchema.nullable(),
  });

  if (data === null) {
    return null;
  }

  return {
    id: data.id,
    customerId: data.customerId,
    vehicleId: data.vehicleId,
    start: data.start,
    end: data.end,
    cost: data.cost,
    createdAt: data.createdAt,
    updatedAt: data.updatedAt,
  };
};
