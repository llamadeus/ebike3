import { z } from "zod";


export const paymentSchema = z.object({
  id: z.string(),
  customerId: z.string(),
  amount: z.number(),
  status: z.enum(["PENDING", "CONFIRMED", "REJECTED"]),
  createdAt: z.string(),
  updatedAt: z.string(),
});
