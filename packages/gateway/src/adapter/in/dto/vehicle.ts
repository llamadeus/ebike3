import { z } from "zod";


export const vehicleSchema = z.object({
  id: z.string(),
  type: z.enum(["BIKE", "EBIKE", "ABIKE"]),
  positionX: z.number(),
  positionY: z.number(),
  battery: z.number(),
  available: z.boolean(),
  activeRental: z.object({
    id: z.string(),
    vehicleId: z.string(),
    customerId: z.string(),
    vehicleType: z.enum(["BIKE", "EBIKE", "ABIKE"]),
    start: z.string(),
    cost: z.number(),
    createdAt: z.string(),
    updatedAt: z.string(),
  }).nullish(),
  createdAt: z.string(),
  updatedAt: z.string(),
});
