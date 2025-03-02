/* This file was automatically generated. DO NOT UPDATE MANUALLY. */
    import type   { Resolvers } from './types.generated';
    import    { auth as Query_auth } from './auth/resolvers/Query/auth';
import    { availableVehicles as Query_availableVehicles } from './vehicles/resolvers/Query/availableVehicles';
import    { customers as Query_customers } from './customers/resolvers/Query/customers';
import    { stations as Query_stations } from './stations/resolvers/Query/stations';
import    { vehicles as Query_vehicles } from './vehicles/resolvers/Query/vehicles';
import    { createStation as Mutation_createStation } from './stations/resolvers/Mutation/createStation';
import    { createVehicle as Mutation_createVehicle } from './vehicles/resolvers/Mutation/createVehicle';
import    { deleteStation as Mutation_deleteStation } from './stations/resolvers/Mutation/deleteStation';
import    { deleteVehicle as Mutation_deleteVehicle } from './vehicles/resolvers/Mutation/deleteVehicle';
import    { login as Mutation_login } from './auth/resolvers/Mutation/login';
import    { logout as Mutation_logout } from './auth/resolvers/Mutation/logout';
import    { registerAdmin as Mutation_registerAdmin } from './auth/resolvers/Mutation/registerAdmin';
import    { registerCustomer as Mutation_registerCustomer } from './auth/resolvers/Mutation/registerCustomer';
import    { updateCustomerPosition as Mutation_updateCustomerPosition } from './customers/resolvers/Mutation/updateCustomerPosition';
import    { Auth } from './auth/resolvers/Auth';
import    { Customer } from './customers/resolvers/Customer';
import    { CustomerRental } from './customers/resolvers/CustomerRental';
import    { Station } from './stations/resolvers/Station';
import    { Vec2d } from './base/resolvers/Vec2d';
import    { Vehicle } from './vehicles/resolvers/Vehicle';
    export const resolvers: Resolvers = {
      Query: { auth: Query_auth,availableVehicles: Query_availableVehicles,customers: Query_customers,stations: Query_stations,vehicles: Query_vehicles },
      Mutation: { createStation: Mutation_createStation,createVehicle: Mutation_createVehicle,deleteStation: Mutation_deleteStation,deleteVehicle: Mutation_deleteVehicle,login: Mutation_login,logout: Mutation_logout,registerAdmin: Mutation_registerAdmin,registerCustomer: Mutation_registerCustomer,updateCustomerPosition: Mutation_updateCustomerPosition },
      
      Auth: Auth,
Customer: Customer,
CustomerRental: CustomerRental,
Station: Station,
Vec2d: Vec2d,
Vehicle: Vehicle
    }