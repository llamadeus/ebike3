import type { QueryResolvers } from "~/schema/types.generated";


export const auth: NonNullable<QueryResolvers['auth']> = async (_parent, _arg, _ctx) => {
  if (_ctx.session === null) {
    return null;
  }

  return {
    id: _ctx.session.id,
    username: _ctx.session.username,
    role: _ctx.session.role,
    lastLogin: _ctx.session.lastLogin,
  };
};
