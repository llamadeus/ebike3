import { z } from "zod";


export const vehicleSchema = z.object({
  id: z.string(),
  type: z.enum(["BIKE", "EBIKE", "ABIKE"]),
  positionX: z.number(),
  positionY: z.number(),
  battery: z.number(),
  available: z.boolean(),
  createdAt: z.string(),
  updatedAt: z.string(),
});
