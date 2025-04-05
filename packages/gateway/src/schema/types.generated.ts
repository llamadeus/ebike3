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

export type CreateStationInput = {
  name: Scalars['String']['input'];
  position: Vec2dInput;
};

export type CreateVehicleInput = {
  position: Vec2dInput;
  type: VehicleType;
};

export type Customer = {
  __typename?: 'Customer';
  activeRental?: Maybe<CustomerRental>;
  creditBalance: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  lastLogin?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  position: Vec2d;
};

export type CustomerRental = {
  __typename?: 'CustomerRental';
  cost: Scalars['Int']['output'];
  customerId: Scalars['ID']['output'];
  id: Scalars['ID']['output'];
  start: Scalars['String']['output'];
  vehicleId: Scalars['ID']['output'];
  vehicleType: VehicleType;
};

export type Expense = {
  __typename?: 'Expense';
  amount: Scalars['Int']['output'];
  createdAt: Scalars['String']['output'];
  customer?: Maybe<Customer>;
  id: Scalars['ID']['output'];
  rentalId: Scalars['ID']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  confirmPayment: Payment;
  createPayment: Payment;
  createStation: Station;
  createVehicle: Vehicle;
  deletePayment: Payment;
  deleteStation: Station;
  deleteVehicle: Vehicle;
  login: Auth;
  logout: Scalars['Boolean']['output'];
  registerAdmin: Auth;
  registerCustomer: Auth;
  rejectPayment: Payment;
  startRental: Rental;
  stopRental: Rental;
  updateCustomerPosition: Scalars['Boolean']['output'];
};


export type MutationconfirmPaymentArgs = {
  id: Scalars['ID']['input'];
};


export type MutationcreatePaymentArgs = {
  amount: Scalars['Int']['input'];
};


export type MutationcreateStationArgs = {
  input: CreateStationInput;
};


export type MutationcreateVehicleArgs = {
  input: CreateVehicleInput;
};


export type MutationdeletePaymentArgs = {
  id: Scalars['ID']['input'];
};


export type MutationdeleteStationArgs = {
  id: Scalars['ID']['input'];
};


export type MutationdeleteVehicleArgs = {
  id: Scalars['ID']['input'];
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


export type MutationrejectPaymentArgs = {
  id: Scalars['ID']['input'];
};


export type MutationstartRentalArgs = {
  vehicleId: Scalars['ID']['input'];
};


export type MutationstopRentalArgs = {
  id: Scalars['ID']['input'];
};


export type MutationupdateCustomerPositionArgs = {
  position: Vec2dInput;
};

export type Payment = {
  __typename?: 'Payment';
  amount: Scalars['Int']['output'];
  createdAt: Scalars['String']['output'];
  customer?: Maybe<Customer>;
  id: Scalars['ID']['output'];
  status: PaymentStatus;
};

export type PaymentStatus =
  | 'CONFIRMED'
  | 'PENDING'
  | 'REJECTED';

export type Query = {
  __typename?: 'Query';
  activeRental?: Maybe<Rental>;
  auth?: Maybe<Auth>;
  availableVehicles: Array<Vehicle>;
  customers: Array<Customer>;
  pastRentals: Array<Rental>;
  payments: Array<Payment>;
  stations: Array<Station>;
  transactions: Array<Transaction>;
  vehicles: Array<Vehicle>;
};


export type QuerypaymentsArgs = {
  status?: InputMaybe<PaymentStatus>;
};

export type Rental = {
  __typename?: 'Rental';
  cost?: Maybe<Scalars['Int']['output']>;
  end?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  start: Scalars['String']['output'];
};

export type Station = {
  __typename?: 'Station';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  position: Vec2d;
};

export type Transaction = Expense | Payment;

export type Vec2d = {
  __typename?: 'Vec2d';
  x: Scalars['Float']['output'];
  y: Scalars['Float']['output'];
};

export type Vec2dInput = {
  x: Scalars['Float']['input'];
  y: Scalars['Float']['input'];
};

export type Vehicle = {
  __typename?: 'Vehicle';
  available: Scalars['Boolean']['output'];
  battery: Scalars['Float']['output'];
  createdAt: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  position: Vec2d;
  type: VehicleType;
};

export type VehicleType =
  | 'ABIKE'
  | 'BIKE'
  | 'EBIKE';



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

/** Mapping of union types */
export type ResolversUnionTypes<_RefType extends Record<string, unknown>> = {
  Transaction: ( Omit<Expense, 'customer'> & { customer?: Maybe<_RefType['Customer']> } & { __typename: 'Expense' } ) | ( Omit<Payment, 'customer' | 'status'> & { customer?: Maybe<_RefType['Customer']>, status: _RefType['PaymentStatus'] } & { __typename: 'Payment' } );
};


/** Mapping between all available schema types and the resolvers types */
export type ResolversTypes = {
  Auth: ResolverTypeWrapper<Omit<Auth, 'role'> & { role: ResolversTypes['AuthRole'] }>;
  ID: ResolverTypeWrapper<Scalars['ID']['output']>;
  String: ResolverTypeWrapper<Scalars['String']['output']>;
  AuthRole: ResolverTypeWrapper<'ADMIN' | 'CUSTOMER'>;
  CreateStationInput: CreateStationInput;
  CreateVehicleInput: CreateVehicleInput;
  Customer: ResolverTypeWrapper<Omit<Customer, 'activeRental'> & { activeRental?: Maybe<ResolversTypes['CustomerRental']> }>;
  Int: ResolverTypeWrapper<Scalars['Int']['output']>;
  CustomerRental: ResolverTypeWrapper<Omit<CustomerRental, 'vehicleType'> & { vehicleType: ResolversTypes['VehicleType'] }>;
  Expense: ResolverTypeWrapper<Omit<Expense, 'customer'> & { customer?: Maybe<ResolversTypes['Customer']> }>;
  Mutation: ResolverTypeWrapper<{}>;
  Boolean: ResolverTypeWrapper<Scalars['Boolean']['output']>;
  Payment: ResolverTypeWrapper<Omit<Payment, 'customer' | 'status'> & { customer?: Maybe<ResolversTypes['Customer']>, status: ResolversTypes['PaymentStatus'] }>;
  PaymentStatus: ResolverTypeWrapper<'PENDING' | 'CONFIRMED' | 'REJECTED'>;
  Query: ResolverTypeWrapper<{}>;
  Rental: ResolverTypeWrapper<Rental>;
  Station: ResolverTypeWrapper<Station>;
  Transaction: ResolverTypeWrapper<ResolversUnionTypes<ResolversTypes>['Transaction']>;
  Vec2d: ResolverTypeWrapper<Vec2d>;
  Float: ResolverTypeWrapper<Scalars['Float']['output']>;
  Vec2dInput: Vec2dInput;
  Vehicle: ResolverTypeWrapper<Omit<Vehicle, 'type'> & { type: ResolversTypes['VehicleType'] }>;
  VehicleType: ResolverTypeWrapper<'BIKE' | 'EBIKE' | 'ABIKE'>;
};

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = {
  Auth: Auth;
  ID: Scalars['ID']['output'];
  String: Scalars['String']['output'];
  CreateStationInput: CreateStationInput;
  CreateVehicleInput: CreateVehicleInput;
  Customer: Omit<Customer, 'activeRental'> & { activeRental?: Maybe<ResolversParentTypes['CustomerRental']> };
  Int: Scalars['Int']['output'];
  CustomerRental: CustomerRental;
  Expense: Omit<Expense, 'customer'> & { customer?: Maybe<ResolversParentTypes['Customer']> };
  Mutation: {};
  Boolean: Scalars['Boolean']['output'];
  Payment: Omit<Payment, 'customer'> & { customer?: Maybe<ResolversParentTypes['Customer']> };
  Query: {};
  Rental: Rental;
  Station: Station;
  Transaction: ResolversUnionTypes<ResolversParentTypes>['Transaction'];
  Vec2d: Vec2d;
  Float: Scalars['Float']['output'];
  Vec2dInput: Vec2dInput;
  Vehicle: Vehicle;
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

export type CustomerResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Customer'] = ResolversParentTypes['Customer']> = {
  activeRental?: Resolver<Maybe<ResolversTypes['CustomerRental']>, ParentType, ContextType>;
  creditBalance?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  lastLogin?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  position?: Resolver<ResolversTypes['Vec2d'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type CustomerRentalResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['CustomerRental'] = ResolversParentTypes['CustomerRental']> = {
  cost?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  customerId?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  start?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  vehicleId?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  vehicleType?: Resolver<ResolversTypes['VehicleType'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type ExpenseResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Expense'] = ResolversParentTypes['Expense']> = {
  amount?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  createdAt?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  customer?: Resolver<Maybe<ResolversTypes['Customer']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  rentalId?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type MutationResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Mutation'] = ResolversParentTypes['Mutation']> = {
  confirmPayment?: Resolver<ResolversTypes['Payment'], ParentType, ContextType, RequireFields<MutationconfirmPaymentArgs, 'id'>>;
  createPayment?: Resolver<ResolversTypes['Payment'], ParentType, ContextType, RequireFields<MutationcreatePaymentArgs, 'amount'>>;
  createStation?: Resolver<ResolversTypes['Station'], ParentType, ContextType, RequireFields<MutationcreateStationArgs, 'input'>>;
  createVehicle?: Resolver<ResolversTypes['Vehicle'], ParentType, ContextType, RequireFields<MutationcreateVehicleArgs, 'input'>>;
  deletePayment?: Resolver<ResolversTypes['Payment'], ParentType, ContextType, RequireFields<MutationdeletePaymentArgs, 'id'>>;
  deleteStation?: Resolver<ResolversTypes['Station'], ParentType, ContextType, RequireFields<MutationdeleteStationArgs, 'id'>>;
  deleteVehicle?: Resolver<ResolversTypes['Vehicle'], ParentType, ContextType, RequireFields<MutationdeleteVehicleArgs, 'id'>>;
  login?: Resolver<ResolversTypes['Auth'], ParentType, ContextType, RequireFields<MutationloginArgs, 'password' | 'username'>>;
  logout?: Resolver<ResolversTypes['Boolean'], ParentType, ContextType>;
  registerAdmin?: Resolver<ResolversTypes['Auth'], ParentType, ContextType, RequireFields<MutationregisterAdminArgs, 'password' | 'username'>>;
  registerCustomer?: Resolver<ResolversTypes['Auth'], ParentType, ContextType, RequireFields<MutationregisterCustomerArgs, 'password' | 'username'>>;
  rejectPayment?: Resolver<ResolversTypes['Payment'], ParentType, ContextType, RequireFields<MutationrejectPaymentArgs, 'id'>>;
  startRental?: Resolver<ResolversTypes['Rental'], ParentType, ContextType, RequireFields<MutationstartRentalArgs, 'vehicleId'>>;
  stopRental?: Resolver<ResolversTypes['Rental'], ParentType, ContextType, RequireFields<MutationstopRentalArgs, 'id'>>;
  updateCustomerPosition?: Resolver<ResolversTypes['Boolean'], ParentType, ContextType, RequireFields<MutationupdateCustomerPositionArgs, 'position'>>;
};

export type PaymentResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Payment'] = ResolversParentTypes['Payment']> = {
  amount?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  createdAt?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  customer?: Resolver<Maybe<ResolversTypes['Customer']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  status?: Resolver<ResolversTypes['PaymentStatus'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type PaymentStatusResolvers = EnumResolverSignature<{ CONFIRMED?: any, PENDING?: any, REJECTED?: any }, ResolversTypes['PaymentStatus']>;

export type QueryResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Query'] = ResolversParentTypes['Query']> = {
  activeRental?: Resolver<Maybe<ResolversTypes['Rental']>, ParentType, ContextType>;
  auth?: Resolver<Maybe<ResolversTypes['Auth']>, ParentType, ContextType>;
  availableVehicles?: Resolver<Array<ResolversTypes['Vehicle']>, ParentType, ContextType>;
  customers?: Resolver<Array<ResolversTypes['Customer']>, ParentType, ContextType>;
  pastRentals?: Resolver<Array<ResolversTypes['Rental']>, ParentType, ContextType>;
  payments?: Resolver<Array<ResolversTypes['Payment']>, ParentType, ContextType, Partial<QuerypaymentsArgs>>;
  stations?: Resolver<Array<ResolversTypes['Station']>, ParentType, ContextType>;
  transactions?: Resolver<Array<ResolversTypes['Transaction']>, ParentType, ContextType>;
  vehicles?: Resolver<Array<ResolversTypes['Vehicle']>, ParentType, ContextType>;
};

export type RentalResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Rental'] = ResolversParentTypes['Rental']> = {
  cost?: Resolver<Maybe<ResolversTypes['Int']>, ParentType, ContextType>;
  end?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  start?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type StationResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Station'] = ResolversParentTypes['Station']> = {
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  position?: Resolver<ResolversTypes['Vec2d'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type TransactionResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Transaction'] = ResolversParentTypes['Transaction']> = {
  __resolveType?: TypeResolveFn<'Expense' | 'Payment', ParentType, ContextType>;
};

export type Vec2dResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Vec2d'] = ResolversParentTypes['Vec2d']> = {
  x?: Resolver<ResolversTypes['Float'], ParentType, ContextType>;
  y?: Resolver<ResolversTypes['Float'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type VehicleResolvers<ContextType = ResolverContext, ParentType extends ResolversParentTypes['Vehicle'] = ResolversParentTypes['Vehicle']> = {
  available?: Resolver<ResolversTypes['Boolean'], ParentType, ContextType>;
  battery?: Resolver<ResolversTypes['Float'], ParentType, ContextType>;
  createdAt?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  position?: Resolver<ResolversTypes['Vec2d'], ParentType, ContextType>;
  type?: Resolver<ResolversTypes['VehicleType'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type VehicleTypeResolvers = EnumResolverSignature<{ ABIKE?: any, BIKE?: any, EBIKE?: any }, ResolversTypes['VehicleType']>;

export type Resolvers<ContextType = ResolverContext> = {
  Auth?: AuthResolvers<ContextType>;
  AuthRole?: AuthRoleResolvers;
  Customer?: CustomerResolvers<ContextType>;
  CustomerRental?: CustomerRentalResolvers<ContextType>;
  Expense?: ExpenseResolvers<ContextType>;
  Mutation?: MutationResolvers<ContextType>;
  Payment?: PaymentResolvers<ContextType>;
  PaymentStatus?: PaymentStatusResolvers;
  Query?: QueryResolvers<ContextType>;
  Rental?: RentalResolvers<ContextType>;
  Station?: StationResolvers<ContextType>;
  Transaction?: TransactionResolvers<ContextType>;
  Vec2d?: Vec2dResolvers<ContextType>;
  Vehicle?: VehicleResolvers<ContextType>;
  VehicleType?: VehicleTypeResolvers;
};

export type DirectiveResolvers<ContextType = ResolverContext> = {
  loggedIn?: loggedInDirectiveResolver<any, any, ContextType>;
  notLoggedIn?: notLoggedInDirectiveResolver<any, any, ContextType>;
};
