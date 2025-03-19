import { z } from "zod";


export const expenseSchema = z.object({
  id: z.string(),
  customerId: z.string(),
  rentalId: z.string(),
  amount: z.number(),
  createdAt: z.string(),
  updatedAt: z.string(),
});
