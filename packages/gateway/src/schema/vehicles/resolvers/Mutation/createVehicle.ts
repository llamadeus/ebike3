import { GraphQLError } from "graphql/error";
import { vehicleSchema } from "~/adapter/in/dto/vehicle";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const createVehicle: NonNullable<MutationResolvers["createVehicle"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "ADMIN") {
    throw new GraphQLError("Not authorized");
  }

  const data = await invokeService("vehicles", {
    endpoint: "PUT /vehicles",
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    input: {
      type: _arg.input.type,
      positionX: _arg.input.position.x,
      positionY: _arg.input.position.y,
    },
    output: vehicleSchema,
  });

  return {
    id: data.id,
    type: data.type,
    position: {
      x: data.positionX,
      y: data.positionY,
    },
    battery: data.battery,
    available: data.available,
    activeRental: data.activeRental,
    createdAt: data.createdAt,
    updatedAt: data.updatedAt,
  };
};
