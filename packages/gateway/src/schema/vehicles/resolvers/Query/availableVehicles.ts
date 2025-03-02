import { GraphQLError } from "graphql/error";
import { z } from "zod";
import { vehicleSchema } from "~/adapter/in/dto/vehicle.ts";
import { invokeService } from "~/infrastructure/service";
import type { QueryResolvers } from "~/schema/types.generated";


export const availableVehicles: NonNullable<QueryResolvers["availableVehicles"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session === null) {
    throw new GraphQLError("Not authenticated");
  }

  const data = await invokeService("vehicles", {
    endpoint: "GET /vehicles/available",
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: z.array(vehicleSchema),
  });

  return data.map((vehicle) => ({
    id: vehicle.id,
    type: vehicle.type,
    position: {
      x: vehicle.positionX,
      y: vehicle.positionY,
    },
    battery: vehicle.battery,
    available: vehicle.available,
    createdAt: vehicle.createdAt,
    updatedAt: vehicle.updatedAt,
  }));
};
