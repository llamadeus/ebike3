import { z } from "zod";


export const creditBalanceSchema = z.object({
  customerId: z.string(),
  creditBalance: z.number(),
});
