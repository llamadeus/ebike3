import { z } from "zod";


export const customerSchema = z.object({
  id: z.string(),
  name: z.string(),
  positionX: z.number(),
  positionY: z.number(),
  creditBalance: z.number(),
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
  lastLogin: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
});
