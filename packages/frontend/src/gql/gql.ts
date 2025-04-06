/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 * Learn more about it here: https://the-guild.dev/graphql/codegen/plugins/presets/preset-client#reducing-bundle-size
 */
const documents = {
    "\n  query AdminView {\n    auth {\n      id\n      role\n      username\n      lastLogin\n    }\n    users {\n      id\n      role\n      username\n      lastLogin\n    }\n    vehicles {\n      id\n      position {\n        x\n        y\n      }\n      battery\n      createdAt\n      activeRental {\n        id\n        customerId\n        start\n        cost\n      }\n    }\n    payments {\n      id\n      amount\n      status\n      createdAt\n      customer {\n        id\n        name\n      }\n    }\n    customers {\n      id\n      name\n      creditBalance\n      lastLogin\n      position {\n        x\n        y\n      }\n      activeRental {\n        id\n        start\n        vehicleId\n        vehicleType\n      }\n    }\n  }\n": types.AdminViewDocument,
    "\n  query CustomerView {\n    auth {\n      id\n      role\n      username\n      lastLogin\n    }\n    creditBalance\n    activeRental {\n      id\n      start\n      end\n      cost\n    }\n    pastRentals {\n      id\n      start\n      end\n      cost\n    }\n    availableVehicles {\n      id\n      battery\n      position {\n        x\n        y\n      }\n    }\n    transactions {\n      __typename\n      ... on Payment {\n        id\n        amount\n        status\n        createdAt\n      }\n      ... on Expense {\n        id\n        amount\n        rentalId\n        createdAt\n      }\n    }\n  }\n": types.CustomerViewDocument,
    "\n  mutation CreateVehicle($input: CreateVehicleInput!) {\n    createVehicle(input: $input) {\n      id\n    }\n  }\n": types.CreateVehicleDocument,
    "\n  mutation ConfirmPayment($id: ID!) {\n    confirmPayment(id: $id) {\n      id\n    }\n  }\n": types.ConfirmPaymentDocument,
    "\n  mutation RejectPayment($id: ID!) {\n    rejectPayment(id: $id) {\n      id\n    }\n  }\n": types.RejectPaymentDocument,
    "\n  mutation deleteVehicle($id: ID!) {\n    deleteVehicle(id: $id) {\n      id\n    }\n  }\n": types.DeleteVehicleDocument,
    "\n  mutation Login($username: String!, $password: String!) {\n    login(username: $username, password: $password) {\n      __typename\n    }\n  }\n": types.LoginDocument,
    "\n  mutation RegisterAdmin($username: String!, $password: String!) {\n    registerAdmin(username: $username, password: $password) {\n      id\n    }\n  }\n": types.RegisterAdminDocument,
    "\n  mutation RegisterCustomer($username: String!, $password: String!) {\n    registerCustomer(username: $username, password: $password) {\n      id\n    }\n  }\n": types.RegisterCustomerDocument,
    "\n  mutation Logout {\n    logout\n  }\n": types.LogoutDocument,
    "\n  query ServerAuth {\n    auth {\n      id\n      role\n      username\n      lastLogin\n    }\n  }\n": types.ServerAuthDocument,
    "\n  mutation CreatePayment($amount: Int!) {\n    createPayment(amount: $amount) {\n      id\n    }\n  }\n": types.CreatePaymentDocument,
    "\n  mutation StopRental($id: ID!) {\n    stopRental(id: $id) {\n      id\n    }\n  }\n": types.StopRentalDocument,
    "\n  mutation StartRental($vehicleId: ID!) {\n    startRental(vehicleId: $vehicleId) {\n      id\n    }\n  }\n": types.StartRentalDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query AdminView {\n    auth {\n      id\n      role\n      username\n      lastLogin\n    }\n    users {\n      id\n      role\n      username\n      lastLogin\n    }\n    vehicles {\n      id\n      position {\n        x\n        y\n      }\n      battery\n      createdAt\n      activeRental {\n        id\n        customerId\n        start\n        cost\n      }\n    }\n    payments {\n      id\n      amount\n      status\n      createdAt\n      customer {\n        id\n        name\n      }\n    }\n    customers {\n      id\n      name\n      creditBalance\n      lastLogin\n      position {\n        x\n        y\n      }\n      activeRental {\n        id\n        start\n        vehicleId\n        vehicleType\n      }\n    }\n  }\n"): (typeof documents)["\n  query AdminView {\n    auth {\n      id\n      role\n      username\n      lastLogin\n    }\n    users {\n      id\n      role\n      username\n      lastLogin\n    }\n    vehicles {\n      id\n      position {\n        x\n        y\n      }\n      battery\n      createdAt\n      activeRental {\n        id\n        customerId\n        start\n        cost\n      }\n    }\n    payments {\n      id\n      amount\n      status\n      createdAt\n      customer {\n        id\n        name\n      }\n    }\n    customers {\n      id\n      name\n      creditBalance\n      lastLogin\n      position {\n        x\n        y\n      }\n      activeRental {\n        id\n        start\n        vehicleId\n        vehicleType\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query CustomerView {\n    auth {\n      id\n      role\n      username\n      lastLogin\n    }\n    creditBalance\n    activeRental {\n      id\n      start\n      end\n      cost\n    }\n    pastRentals {\n      id\n      start\n      end\n      cost\n    }\n    availableVehicles {\n      id\n      battery\n      position {\n        x\n        y\n      }\n    }\n    transactions {\n      __typename\n      ... on Payment {\n        id\n        amount\n        status\n        createdAt\n      }\n      ... on Expense {\n        id\n        amount\n        rentalId\n        createdAt\n      }\n    }\n  }\n"): (typeof documents)["\n  query CustomerView {\n    auth {\n      id\n      role\n      username\n      lastLogin\n    }\n    creditBalance\n    activeRental {\n      id\n      start\n      end\n      cost\n    }\n    pastRentals {\n      id\n      start\n      end\n      cost\n    }\n    availableVehicles {\n      id\n      battery\n      position {\n        x\n        y\n      }\n    }\n    transactions {\n      __typename\n      ... on Payment {\n        id\n        amount\n        status\n        createdAt\n      }\n      ... on Expense {\n        id\n        amount\n        rentalId\n        createdAt\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation CreateVehicle($input: CreateVehicleInput!) {\n    createVehicle(input: $input) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation CreateVehicle($input: CreateVehicleInput!) {\n    createVehicle(input: $input) {\n      id\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation ConfirmPayment($id: ID!) {\n    confirmPayment(id: $id) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation ConfirmPayment($id: ID!) {\n    confirmPayment(id: $id) {\n      id\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation RejectPayment($id: ID!) {\n    rejectPayment(id: $id) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation RejectPayment($id: ID!) {\n    rejectPayment(id: $id) {\n      id\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation deleteVehicle($id: ID!) {\n    deleteVehicle(id: $id) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation deleteVehicle($id: ID!) {\n    deleteVehicle(id: $id) {\n      id\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation Login($username: String!, $password: String!) {\n    login(username: $username, password: $password) {\n      __typename\n    }\n  }\n"): (typeof documents)["\n  mutation Login($username: String!, $password: String!) {\n    login(username: $username, password: $password) {\n      __typename\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation RegisterAdmin($username: String!, $password: String!) {\n    registerAdmin(username: $username, password: $password) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation RegisterAdmin($username: String!, $password: String!) {\n    registerAdmin(username: $username, password: $password) {\n      id\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation RegisterCustomer($username: String!, $password: String!) {\n    registerCustomer(username: $username, password: $password) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation RegisterCustomer($username: String!, $password: String!) {\n    registerCustomer(username: $username, password: $password) {\n      id\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation Logout {\n    logout\n  }\n"): (typeof documents)["\n  mutation Logout {\n    logout\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query ServerAuth {\n    auth {\n      id\n      role\n      username\n      lastLogin\n    }\n  }\n"): (typeof documents)["\n  query ServerAuth {\n    auth {\n      id\n      role\n      username\n      lastLogin\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation CreatePayment($amount: Int!) {\n    createPayment(amount: $amount) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation CreatePayment($amount: Int!) {\n    createPayment(amount: $amount) {\n      id\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation StopRental($id: ID!) {\n    stopRental(id: $id) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation StopRental($id: ID!) {\n    stopRental(id: $id) {\n      id\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation StartRental($vehicleId: ID!) {\n    startRental(vehicleId: $vehicleId) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation StartRental($vehicleId: ID!) {\n    startRental(vehicleId: $vehicleId) {\n      id\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;