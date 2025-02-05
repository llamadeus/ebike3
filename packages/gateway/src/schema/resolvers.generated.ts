/* This file was automatically generated. DO NOT UPDATE MANUALLY. */
    import type   { Resolvers } from './types.generated';
    import    { auth as Query_auth } from './auth/resolvers/Query/auth';
import    { login as Mutation_login } from './auth/resolvers/Mutation/login';
import    { logout as Mutation_logout } from './auth/resolvers/Mutation/logout';
import    { registerAdmin as Mutation_registerAdmin } from './auth/resolvers/Mutation/registerAdmin';
import    { registerCustomer as Mutation_registerCustomer } from './auth/resolvers/Mutation/registerCustomer';
import    { Auth } from './auth/resolvers/Auth';
    export const resolvers: Resolvers = {
      Query: { auth: Query_auth },
      Mutation: { login: Mutation_login,logout: Mutation_logout,registerAdmin: Mutation_registerAdmin,registerCustomer: Mutation_registerCustomer },
      
      Auth: Auth
    }