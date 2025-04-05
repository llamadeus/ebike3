import { z } from "zod";


export const rentalSchema = z.object({
  id: z.string(),
  customerId: z.string(),
  vehicleId: z.string(),
  start: z.string(),
  end: z.string().nullable(),
  cost: z.number(),
  createdAt: z.string(),
  updatedAt: z.string(),
});
