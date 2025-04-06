"use client";

import { useEffect } from "react";
import { useQuery } from "urql";
import { Section } from "~/components/Section";
import { Badge } from "~/components/ui/badge";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "~/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "~/components/ui/tabs";
import { CreateStationButton } from "~/features/admin/components/CreateStationButton";
import { CreateVehicleButton } from "~/features/admin/components/CreateVehicleButton";
import { PaymentsTable } from "~/features/admin/components/PaymentsTable";
import { StationsTable } from "~/features/admin/components/StationsTable";
import { UsersTable } from "~/features/admin/components/UsersTable";
import { VehiclesTable } from "~/features/admin/components/VehiclesTable";
import { LogoutButton } from "~/features/auth/components/LogoutButton";
import { graphql } from "~/gql";
import { isNullish } from "~/utils/value";


const queryAdminViewDocument = graphql(`
  query AdminView {
    auth {
      id
      role
      username
      lastLogin
    }
    users {
      id
      role
      username
      lastLogin
    }
    stations {
      id
      name
      position {
        x
        y
      }
      createdAt
    }
    vehicles {
      id
      type
      position {
        x
        y
      }
      battery
      createdAt
      activeRental {
        id
        customerId
        start
        cost
      }
    }
    payments {
      id
      amount
      status
      createdAt
      customer {
        id
        name
      }
    }
    customers {
      id
      name
      creditBalance
      lastLogin
      position {
        x
        y
      }
      activeRental {
        id
        start
        vehicleId
        vehicleType
      }
    }
  }
`);

export default function Admin() {
  const [{ data }, refetch] = useQuery({ query: queryAdminViewDocument });
  const rentedVehicles = data?.vehicles?.filter(vehicle => vehicle.activeRental !== null) ?? [];
  const pendingPayments = data?.payments?.filter(payment => payment.status === "PENDING") ?? [];

  useEffect(() => {
    const interval = setInterval(() => {
      refetch();
    }, 1000);

    return () => clearInterval(interval);
  }, [refetch]);

  if (isNullish(data?.auth) || data.auth.role !== "ADMIN") {
    return null;
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle>EBike / Admin</CardTitle>
        <CardDescription>Welcome to the admin panel, <b>{data.auth.username}</b>.</CardDescription>
      </CardHeader>
      <CardContent className="flex flex-col gap-8">
        <Tabs defaultValue="stations">
          <TabsList>
            <TabsTrigger value="stations">
              Stations
            </TabsTrigger>
            <TabsTrigger value="vehicles">
              Vehicles
              {rentedVehicles.length > 0 && (
                <Badge className="ml-2 rounded-full py-0 px-1.5">
                  {rentedVehicles.length}
                </Badge>
              )}
            </TabsTrigger>
            <TabsTrigger value="payments">
              Payments
              {pendingPayments.length > 0 && (
                <Badge className="ml-2 rounded-full py-0 px-1.5">
                  {pendingPayments.length}
                </Badge>
              )}
            </TabsTrigger>
            <TabsTrigger value="users">Users</TabsTrigger>
          </TabsList>

          <div className="h-px mt-4 mb-8 -mx-4 bg-border"></div>

          <TabsContent value="stations">
            <Section caption="Stations" right={<CreateStationButton/>}>
              <StationsTable stations={data.stations}/>
            </Section>
          </TabsContent>

          <TabsContent value="vehicles">
            <Section caption="Vehicles" right={<CreateVehicleButton/>}>
              <VehiclesTable vehicles={data.vehicles}/>
            </Section>
          </TabsContent>

          <TabsContent value="payments">
            <Section caption="Payments">
              <PaymentsTable payments={data.payments}/>
            </Section>
          </TabsContent>

          <TabsContent value="users">
            <Section caption="Users">
              <UsersTable users={data.users} userId={data.auth.id}/>
            </Section>
          </TabsContent>
        </Tabs>
      </CardContent>

      <CardFooter className="flex justify-end gap-2 pt-6">
        <LogoutButton variant="destructive"/>
      </CardFooter>
    </Card>
  );
}
