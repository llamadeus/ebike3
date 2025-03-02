import { GraphQLError } from "graphql/error";
import { stationSchema } from "~/adapter/in/dto/station";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const createStation: NonNullable<MutationResolvers["createStation"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "ADMIN") {
    throw new GraphQLError("Not authorized");
  }

  const data = await invokeService("stations", {
    endpoint: "PUT /stations",
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    input: {
      name: _arg.input.name,
      positionX: _arg.input.position.x,
      positionY: _arg.input.position.y,
    },
    output: stationSchema,
  });

  return {
    id: data.id,
    name: data.name,
    position: {
      x: data.positionX,
      y: data.positionY,
    },
    createdAt: data.createdAt,
    updatedAt: data.updatedAt,
  };
};
