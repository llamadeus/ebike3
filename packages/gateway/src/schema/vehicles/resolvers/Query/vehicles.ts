import { GraphQLError } from "graphql/error";
import { z } from "zod";
import { vehicleSchema } from "~/adapter/in/dto/vehicle";
import { invokeService } from "~/infrastructure/service";
import type { QueryResolvers } from "~/schema/types.generated";


export const vehicles: NonNullable<QueryResolvers["vehicles"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "ADMIN") {
    throw new GraphQLError("Not authorized");
  }

  const data = await invokeService("vehicles", {
    endpoint: "GET /vehicles",
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
