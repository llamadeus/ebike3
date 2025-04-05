import { z } from "zod";
import { userSchema } from "~/adapter/in/dto/user";
import { invokeService } from "~/infrastructure/service";
import type { QueryResolvers } from "~/schema/types.generated";


export const users: NonNullable<QueryResolvers["users"]> = async (_parent, _arg, _ctx) => {
  if (_ctx.session?.role !== "ADMIN") {
    throw new Error("Not authorized");
  }

  const data = await invokeService("auth", {
    endpoint: "GET /users",
    headers: {
      "X-Request-ID": _ctx.requestId,
    },
    output: z.array(userSchema),
  });

  return data.map((user) => ({
    id: user.id,
    username: user.username,
    role: user.role,
    lastLogin: user.lastLogin,
  }));
};
