import { GraphQLError } from "graphql/error";
import { rentalSchema } from "~/adapter/in/dto/rental";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const stopRental: NonNullable<MutationResolvers["stopRental"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "CUSTOMER") {
    throw new GraphQLError("Not authorized");
  }

  const data = await invokeService("rentals", {
    endpoint: `POST /rentals/${_arg.id}/stop`,
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    input: {
      customerId: _ctx.session.id,
    },
    output: rentalSchema,
  });

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
