import { GraphQLResolveInfo } from 'graphql';
import { ResolverContext } from '../infrastructure/server';
export type Maybe<T> = T | null | undefined;
export type InputMaybe<T> = T | null | undefined;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
export type EnumResolverSignature<T, AllowedValues = any> = { [key in keyof T]?: AllowedValues };
export type RequireFields<T, K extends keyof T> = Omit<T, K> & { [P in K]-?: NonNullable<T[P]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string | number; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type Auth = {
  __typename?: 'Auth';
  id: Scalars['ID']['output'];
  lastLogin?: Maybe<Scalars['String']['output']>;
  role: AuthRole;
  username: Scalars['String']['output'];
};

export type AuthRole =
  | 'ADMIN'
  | 'CUSTOMER';

export type Mutation = {
  __typename?: 'Mutation';
  login: Auth;
  logout: Scalars['Boolean']['output'];
  registerAdmin: Auth;
  registerCustomer: Auth;
};


export type MutationloginArgs = {
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
};


export type MutationregisterAdminArgs = {
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
};


export type MutationregisterCustomerArgs = {
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
};

export type Query = {
  __typename?: 'Query';
  auth?: Maybe<Auth>;
};



export type ResolverTypeWrapper<T> = Promise<T> | T;


export type ResolverWithResolve<TResult, TParent, TContext, TArgs> = {
  resolve: ResolverFn<TResult, TParent, TContext, TArgs>;
};
export type Resolver<TResult, TParent = {}, TContext = {}, TArgs = {}> = ResolverFn<TResult, TParent, TContext, TArgs> | ResolverWithResolve<TResult, TParent, TContext, TArgs>;

export type ResolverFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => Promise<TResult> | TResult;

export type SubscriptionSubscribeFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => AsyncIterable<TResult> | Promise<AsyncIterable<TResult>>;

export type SubscriptionResolveFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => TResult | Promise<TResult>;

export interface SubscriptionSubscriberObject<TResult, TKey extends string, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<{ [key in TKey]: TResult }, TParent, TContext, TArgs>;
  resolve?: SubscriptionResolveFn<TResult, { [key in TKey]: TResult }, TContext, TArgs>;
}

export interface SubscriptionResolverObject<TResult, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<any, TParent, TContext, TArgs>;
  resolve: SubscriptionResolveFn<TResult, any, TContext, TArgs>;
}

export type SubscriptionObject<TResult, TKey extends string, TParent, TContext, TArgs> =
  | SubscriptionSubscriberObject<TResult, TKey, TParent, TContext, TArgs>
  | SubscriptionResolverObject<TResult, TParent, TContext, TArgs>;

export type SubscriptionResolver<TResult, TKey extends string, TParent = {}, TContext = {}, TArgs = {}> =
  | ((...args: any[]) => SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>)
  | SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>;

export type TypeResolveFn<TTypes, TParent = {}, TContext = {}> = (
  parent: TParent,
  context: TContext,
  info: GraphQLResolveInfo
) => Maybe<TTypes> | Promise<Maybe<TTypes>>;

export type IsTypeOfResolverFn<T = {}, TContext = {}> = (obj: T, context: TContext, info: GraphQLResolveInfo) => boolean | Promise<boolean>;

export type NextResolverFn<T> = () => Promise<T>;

export type DirectiveResolverFn<TResult = {}, TParent = {}, TContext = {}, TArgs = {}> = (
  next: NextResolverFn<TResult>,
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => TResult | Promise<TResult>;



/** Mapping between all available schema types and the resolvers types */
export type ResolversTypes = {
  Auth: ResolverTypeWrapper<Omit<Auth, 'role'> & { role: ResolversTypes['AuthRole'] }>;
  ID: ResolverTypeWrapper<Scalars['ID']['output']>;
  String: ResolverTypeWrapper<Scalars['String']['output']>;
  AuthRole: ResolverTypeWrapper<'ADMIN' | 'CUSTOMER'>;
  Mutation: ResolverTypeWrapper<{}>;
  Boolean: ResolverTypeWrapper<Scalars['Boolean']['output']>;
  Query: ResolverTypeWrapper<{}>;
};

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = {
  Auth: Auth;
  ID: Scalars['ID']['output'];
  String: Scalars['String']['output'];
  Mutation: {};
  Boolean: Scalars['Boolean']['output'];
  Query: {};
};

export type loggedInDirectiveArgs = { };

export type loggedInDirectiveResolver<Result, Parent, ContextType = ResolverContext, Args = loggedInDirectiveArgs> = DirectiveResolverFn<Result, Parent, ContextType, Args>;

export type notLoggedInDirectiveArgs = { };

export type notLoggedInDirectiveResolver<Result, Parent, ContextType = ResolverContext, Args = notLoggedInDirectiveArgs> = DirectiveResolverFn<Result, Parent, ContextType, Args>;

export type AuthResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Auth'] = ResolversParentTypes['Auth']> = {
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  lastLogin?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  role?: Resolver<ResolversTypes['AuthRole'], ParentType, ContextType>;
  username?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type AuthRoleResolvers = EnumResolverSignature<{ ADMIN?: any, CUSTOMER?: any }, ResolversTypes['AuthRole']>;

export type MutationResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Mutation'] = ResolversParentTypes['Mutation']> = {
  login?: Resolver<ResolversTypes['Auth'], ParentType, ContextType, RequireFields<MutationloginArgs, 'password' | 'username'>>;
  logout?: Resolver<ResolversTypes['Boolean'], ParentType, ContextType>;
  registerAdmin?: Resolver<ResolversTypes['Auth'], ParentType, ContextType, RequireFields<MutationregisterAdminArgs, 'password' | 'username'>>;
  registerCustomer?: Resolver<ResolversTypes['Auth'], ParentType, ContextType, RequireFields<MutationregisterCustomerArgs, 'password' | 'username'>>;
};

export type QueryResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Query'] = ResolversParentTypes['Query']> = {
  auth?: Resolver<Maybe<ResolversTypes['Auth']>, ParentType, ContextType>;
};

export type Resolvers<ContextType = ResolverContext> = {
  Auth?: AuthResolvers<ContextType>;
  AuthRole?: AuthRoleResolvers;
  Mutation?: MutationResolvers<ContextType>;
  Query?: QueryResolvers<ContextType>;
};

export type DirectiveResolvers<ContextType = ResolverContext> = {
  loggedIn?: loggedInDirectiveResolver<any, any, ContextType>;
  notLoggedIn?: notLoggedInDirectiveResolver<any, any, ContextType>;
};
