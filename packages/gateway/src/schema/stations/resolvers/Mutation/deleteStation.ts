import { stationSchema } from "~/adapter/in/dto/station";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const deleteStation: NonNullable<MutationResolvers["deleteStation"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "ADMIN") {
    throw new Error("Not authorized");
  }

  const data = await invokeService("stations", {
    endpoint: `DELETE /stations/${_arg.id}`,
    headers: {
      "X-Request-ID": _ctx.requestId,
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
