"use client";

import { useEffect } from "react";
import { useQuery } from "urql";
import { Section } from "~/components/Section";
import { Badge } from "~/components/ui/badge";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "~/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "~/components/ui/tabs";
import { LogoutButton } from "~/features/auth/components/LogoutButton";
import { TransactionsTable } from "~/features/credit/components/TransactionsTable";
import { CreatePaymentButton } from "~/features/credit/components/CreatePaymentButton";
import { ActiveRental } from "~/features/rental/components/ActiveRental";
import { PastRentalsTable } from "~/features/rental/components/PastRentalsTable";
import { StartRental } from "~/features/rental/components/StartRental";
import { graphql } from "~/gql";
import { formatCurrency } from "~/utils/currency";
import { isNotNullish, isNullish } from "~/utils/value";


const queryCustomerViewDocument = graphql(`
  query CustomerView {
    auth {
      id
      role
      username
      lastLogin
    }
    creditBalance
    activeRental {
      id
      start
      end
      cost
    }
    pastRentals {
      id
      start
      end
      cost
    }
    availableVehicles {
      id
      battery
      position {
        x
        y
      }
    }
    transactions {
      __typename
      ... on Payment {
        id
        amount
        status
        createdAt
      }
      ... on Expense {
        id
        amount
        rentalId
        createdAt
      }
    }
  }
`);

export default function Customer() {
  const [{ data }, refetch] = useQuery({ query: queryCustomerViewDocument });

  useEffect(() => {
    const interval = setInterval(() => {
      refetch();
    }, 1000);

    return () => clearInterval(interval);
  }, [refetch]);

  if (isNullish(data?.auth) || data.auth.role !== "CUSTOMER") {
    return null;
  }

  return (
    <Card>
      <CardHeader className="flex flex-1">
        <CardTitle>EBike / Customer</CardTitle>
        <CardDescription>
          Logged in as: <b>{data.auth.username}</b>.
          Your credit: <b
          className={data.creditBalance < 0
            ? "text-destructive-foreground"
            : "text-foreground"}
        >{formatCurrency(data.creditBalance)}</b>.
        </CardDescription>
      </CardHeader>

      <CardContent className="flex flex-col gap-8">
        <Tabs defaultValue="rental">
          <TabsList>
            <TabsTrigger value="rental">
              Rental
              {isNotNullish(data.activeRental) && (
                <Badge className="ml-2 rounded-full size-2.5 p-0 animate-pulse bg-green-600"/>
              )}
            </TabsTrigger>
            <TabsTrigger value="transactions">Transactions</TabsTrigger>
            <TabsTrigger value="rental-history">Rental history</TabsTrigger>
          </TabsList>

          <div className="h-px mt-4 mb-8 -mx-4 bg-border"></div>

          <TabsContent value="rental">
            {isNotNullish(data.activeRental) && (
              <Section caption="Active rental">
                <ActiveRental rental={data.activeRental}/>
              </Section>
            )}

            {isNullish(data.activeRental) && (
              <Section caption="Available vehicles">
                <StartRental vehicles={data.availableVehicles}/>
              </Section>
            )}
          </TabsContent>

          <TabsContent value="transactions">
            <Section caption="Transactions" right={<CreatePaymentButton/>}>
              <TransactionsTable transactions={data.transactions}/>
            </Section>
          </TabsContent>

          <TabsContent value="rental-history">
            <Section caption="Rental history">
              <PastRentalsTable rentals={data.pastRentals}/>
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
