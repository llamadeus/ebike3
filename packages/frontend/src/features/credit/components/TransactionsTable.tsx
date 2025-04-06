"use client";

import dayjs from "dayjs";
import { ReactNode, useMemo } from "react";
import { FixHydration } from "~/components/FixHydration";
import { Badge } from "~/components/ui/badge";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "~/components/ui/table";
import { Tooltip, TooltipContent, TooltipTrigger } from "~/components/ui/tooltip";
import { Expense, Maybe, Payment, PaymentStatus, Rental } from "~/gql/graphql";
import { pluralize } from "~/lib/plural";
import { cn } from "~/lib/utils";
import { NonNullish, WithTypename } from "~/types";
import { formatCurrency } from "~/utils/currency";


type PaymentType = Pick<WithTypename<Payment>, "__typename" | "id" | "amount" | "status" | "createdAt">;
type ExpenseType = Pick<WithTypename<Expense>, "__typename" | "id" | "amount" | "createdAt" | "rentalId">;
type TransactionType = PaymentType | ExpenseType;

interface Props {
  /**
   * The transactions to display.
   */
  transactions: TransactionType[];
}

interface TableRowData {
  id: string;
  amount: number;
  timestamp: string;
}

interface TableRowGroupData {
  id: string;
  type: "Payment" | "Expense";
  amount: number;
  details: ReactNode;
  timestamp: string;
  children: TableRowData[];
}

export function TransactionsTable(props: Props) {
  const sorted = useMemo(() => (
    props.transactions.sort((a, b) => dayjs(b.createdAt).unix() - dayjs(a.createdAt).unix())
  ), [props.transactions]);
  const grouped = useMemo(() => {
    const groups = new Map<string, TableRowGroupData>();

    for (const transaction of sorted) {
      if (transaction.__typename === "Payment") {
        let details: ReactNode = null;

        switch (transaction.status) {
        case PaymentStatus.Pending:
          details = (
            <Tooltip>
              <TooltipTrigger asChild><Badge variant="secondary">Pending</Badge></TooltipTrigger>
              <TooltipContent>Please hold on as we&apos;re checking your payment.</TooltipContent>
            </Tooltip>
          );
          break;
        case PaymentStatus.Confirmed:
          details = (
            <Tooltip>
              <TooltipTrigger asChild><Badge variant="success">Confirmed</Badge></TooltipTrigger>
              <TooltipContent>This payment has been confirmed.</TooltipContent>
            </Tooltip>
          );
          break;
        case PaymentStatus.Rejected:
          details = (
            <Tooltip>
              <TooltipTrigger asChild><Badge variant="destructive">Rejected</Badge></TooltipTrigger>
              <TooltipContent>
                This payment has been rejected. Please contact our support team.
              </TooltipContent>
            </Tooltip>
          );
          break;
        }

        groups.set(transaction.id, {
          id: transaction.id,
          type: "Payment",
          amount: transaction.amount,
          details: details,
          timestamp: transaction.createdAt,
          children: [],
        });
      }
      else if (transaction.__typename === "Expense") {
        const existing = groups.get(transaction.rentalId) ?? {
          id: transaction.id,
          type: "Expense",
          amount: 0,
          details: (
            <Badge variant="outline" className="gap-1.5">
              <span>Rental</span>
              <code>#{transaction.rentalId}</code>
            </Badge>
          ),
          timestamp: transaction.createdAt,
          children: [],
        };

        groups.set(transaction.rentalId, {
          ...existing,
          amount: existing.amount + transaction.amount,
          children: [
            ...existing.children,
            {
              id: transaction.id,
              amount: transaction.amount,
              timestamp: transaction.createdAt,
            },
          ],
        });
      }
    }

    return [...groups.values()];
  }, [sorted]);

  return (
    <Table>
      <TableCaption>
        {sorted.length > 0 && "A list of all transactions"}
        {sorted.length === 0 && "Nothing to see here"}
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="max-w-24">#</TableHead>
          <TableHead>Type</TableHead>
          <TableHead>Amount</TableHead>
          <TableHead>Details</TableHead>
          <TableHead className="w-36">Date</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {grouped.map((group) => (
          <TableRow key={group.id}>
            <TableCell className="font-medium">{group.id}</TableCell>
            <TableCell>
              {group.type === "Payment" && "Recharge"}
              {group.type === "Expense" && `Expense (${pluralize("item", group.children.length)})`}
            </TableCell>
            <TableCell className={cn(group.type === "Expense" && "text-destructive")}>
              {formatCurrency(group.amount)}
            </TableCell>
            <TableCell>
              {group.details}
            </TableCell>
            <TableCell>
              <FixHydration>
                {dayjs(group.timestamp).format("DD.MM.YYYY HH:mm")}
              </FixHydration>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );

  return (
    <Table>
      <TableCaption>
        {sorted.length > 0 && "A list of all transactions"}
        {sorted.length === 0 && "Nothing to see here"}
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="max-w-24">#</TableHead>
          <TableHead>Type</TableHead>
          <TableHead>Amount</TableHead>
          <TableHead>Details</TableHead>
          <TableHead className="w-36">Date</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {sorted.map((transaction) => (
          <TableRow key={transaction.id}>
            <TableCell className="font-medium">{transaction.id}</TableCell>
            <TableCell>
              {transaction.__typename === "Payment" && "Recharge"}
              {transaction.__typename === "Expense" && "Expense"}
            </TableCell>
            <TableCell className={cn(transaction.__typename === "Expense" && "text-destructive")}>
              {formatCurrency(transaction.amount)}
            </TableCell>
            <TableCell>
              {transaction.__typename === "Payment" && (
                <>
                  {transaction.status === "PENDING" && (
                    <Tooltip>
                      <TooltipTrigger asChild><Badge variant="secondary">Pending</Badge></TooltipTrigger>
                      <TooltipContent>Please hold on as we&apos;re checking your payment.</TooltipContent>
                    </Tooltip>
                  )}
                  {transaction.status === "CONFIRMED" && (
                    <Tooltip>
                      <TooltipTrigger asChild><Badge variant="success">Confirmed</Badge></TooltipTrigger>
                      <TooltipContent>This payment has been confirmed.</TooltipContent>
                    </Tooltip>
                  )}
                  {transaction.status === "REJECTED" && (
                    <Tooltip>
                      <TooltipTrigger asChild><Badge variant="destructive">Rejected</Badge></TooltipTrigger>
                      <TooltipContent>
                        This payment has been rejected. Please contact our support team.
                      </TooltipContent>
                    </Tooltip>
                  )}
                </>
              )}
              {transaction.__typename === "Expense" && (
                <>
                  <Badge variant="outline" className="gap-1.5">
                    <span>Rental</span>
                    <code>#{transaction.rentalId}</code>
                  </Badge>
                </>
              )}
            </TableCell>
            <TableCell>
              <FixHydration>
                {dayjs(transaction.createdAt).format("DD.MM.YYYY HH:mm")}
              </FixHydration>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
