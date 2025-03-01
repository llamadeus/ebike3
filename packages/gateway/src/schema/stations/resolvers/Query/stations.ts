import { GraphQLError } from "graphql/error";
import { z } from "zod";
import { stationSchema } from "~/adapter/in/dto/station";
import { invokeService } from "~/infrastructure/service";
import type { QueryResolvers } from "~/schema/types.generated";


export const stations: NonNullable<QueryResolvers["stations"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session === null) {
    throw new GraphQLError("Not authenticated");
  }

  const data = await invokeService("stations", {
    endpoint: "GET /stations",
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: z.array(stationSchema),
  });

  return data.map((station) => ({
    id: station.id,
    name: station.name,
    position: {
      x: station.positionX,
      y: station.positionY,
    },
    createdAt: station.createdAt,
    updatedAt: station.updatedAt,
  }));
};
