import { vehicleSchema } from "~/adapter/in/dto/vehicle";
import { invokeService } from "~/infrastructure/service";
import type { MutationResolvers } from "~/schema/types.generated";


export const deleteVehicle: NonNullable<MutationResolvers["deleteVehicle"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "ADMIN") {
    throw new Error("Not authorized");
  }

  const data = await invokeService("vehicles", {
    endpoint: `DELETE /vehicles/${_arg.id}`,
    headers: {
      "X-Request-ID": _ctx.requestId,
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
