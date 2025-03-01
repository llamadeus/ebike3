import { z } from "zod";


export const stationSchema = z.object({
  id: z.string(),
  name: z.string(),
  positionX: z.number(),
  positionY: z.number(),
  createdAt: z.string(),
  updatedAt: z.string(),
});
