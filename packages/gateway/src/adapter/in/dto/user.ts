import { z } from "zod";


export const userSchema = z.object({
  id: z.string(),
  username: z.string(),
  role: z.enum(["ADMIN", "CUSTOMER"]),
  sessionId: z.string(),
  lastLogin: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
});
