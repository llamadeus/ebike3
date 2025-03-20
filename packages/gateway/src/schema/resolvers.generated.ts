/* This file was automatically generated. DO NOT UPDATE MANUALLY. */
    import type   { Resolvers } from './types.generated';
    import    { activeRental as Query_activeRental } from './rentals/resolvers/Query/activeRental';
import    { auth as Query_auth } from './auth/resolvers/Query/auth';
import    { availableVehicles as Query_availableVehicles } from './vehicles/resolvers/Query/availableVehicles';
import    { customers as Query_customers } from './customers/resolvers/Query/customers';
import    { pastRentals as Query_pastRentals } from './rentals/resolvers/Query/pastRentals';
import    { payments as Query_payments } from './accounting/resolvers/Query/payments';
import    { stations as Query_stations } from './stations/resolvers/Query/stations';
import    { transactions as Query_transactions } from './accounting/resolvers/Query/transactions';
import    { vehicles as Query_vehicles } from './vehicles/resolvers/Query/vehicles';
import    { confirmPayment as Mutation_confirmPayment } from './accounting/resolvers/Mutation/confirmPayment';
import    { createPayment as Mutation_createPayment } from './accounting/resolvers/Mutation/createPayment';
import    { createStation as Mutation_createStation } from './stations/resolvers/Mutation/createStation';
import    { createVehicle as Mutation_createVehicle } from './vehicles/resolvers/Mutation/createVehicle';
import    { deletePayment as Mutation_deletePayment } from './accounting/resolvers/Mutation/deletePayment';
import    { deleteStation as Mutation_deleteStation } from './stations/resolvers/Mutation/deleteStation';
import    { deleteVehicle as Mutation_deleteVehicle } from './vehicles/resolvers/Mutation/deleteVehicle';
import    { login as Mutation_login } from './auth/resolvers/Mutation/login';
import    { logout as Mutation_logout } from './auth/resolvers/Mutation/logout';
import    { registerAdmin as Mutation_registerAdmin } from './auth/resolvers/Mutation/registerAdmin';
import    { registerCustomer as Mutation_registerCustomer } from './auth/resolvers/Mutation/registerCustomer';
import    { rejectPayment as Mutation_rejectPayment } from './accounting/resolvers/Mutation/rejectPayment';
import    { startRental as Mutation_startRental } from './rentals/resolvers/Mutation/startRental';
import    { stopRental as Mutation_stopRental } from './rentals/resolvers/Mutation/stopRental';
import    { updateCustomerPosition as Mutation_updateCustomerPosition } from './customers/resolvers/Mutation/updateCustomerPosition';
import    { Auth } from './auth/resolvers/Auth';
import    { Customer } from './customers/resolvers/Customer';
import    { CustomerRental } from './customers/resolvers/CustomerRental';
import    { Expense } from './accounting/resolvers/Expense';
import    { Payment } from './accounting/resolvers/Payment';
import    { Rental } from './rentals/resolvers/Rental';
import    { Station } from './stations/resolvers/Station';
import    { Vec2d } from './base/resolvers/Vec2d';
import    { Vehicle } from './vehicles/resolvers/Vehicle';
    export const resolvers: Resolvers = {
      Query: { activeRental: Query_activeRental,auth: Query_auth,availableVehicles: Query_availableVehicles,customers: Query_customers,pastRentals: Query_pastRentals,payments: Query_payments,stations: Query_stations,transactions: Query_transactions,vehicles: Query_vehicles },
      Mutation: { confirmPayment: Mutation_confirmPayment,createPayment: Mutation_createPayment,createStation: Mutation_createStation,createVehicle: Mutation_createVehicle,deletePayment: Mutation_deletePayment,deleteStation: Mutation_deleteStation,deleteVehicle: Mutation_deleteVehicle,login: Mutation_login,logout: Mutation_logout,registerAdmin: Mutation_registerAdmin,registerCustomer: Mutation_registerCustomer,rejectPayment: Mutation_rejectPayment,startRental: Mutation_startRental,stopRental: Mutation_stopRental,updateCustomerPosition: Mutation_updateCustomerPosition },
      
      Auth: Auth,
Customer: Customer,
CustomerRental: CustomerRental,
Expense: Expense,
Payment: Payment,
Rental: Rental,
Station: Station,
Vec2d: Vec2d,
Vehicle: Vehicle
    }