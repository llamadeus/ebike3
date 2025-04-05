import { customerSchema } from "~/adapter/in/dto/customer";
import { invokeService } from "~/infrastructure/service";
import { isNullish } from "~/infrastructure/utils/helpers";
import type { QueryResolvers } from "~/schema/types.generated";


export const customer: NonNullable<QueryResolvers["customer"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "ADMIN") {
    throw new Error("Not authorized");
  }

  const data = await invokeService("customers", {
    endpoint: `GET /customers/${_arg.id}`,
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: customerSchema,
  });

  return {
    id: data.id,
    name: data.name,
    position: {
      x: data.positionX,
      y: data.positionY,
    },
    creditBalance: data.creditBalance,
    activeRental: ! isNullish(data.activeRental)
      ? {
        id: data.activeRental.id,
        vehicleId: data.activeRental.vehicleId,
        customerId: data.activeRental.customerId,
        vehicleType: data.activeRental.vehicleType,
        start: data.activeRental.start,
        cost: data.activeRental.cost,
      }
      : null,
    lastLogin: data.lastLogin,
  };
};
