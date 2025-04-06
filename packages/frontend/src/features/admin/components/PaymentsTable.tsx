"use client";

import dayjs from "dayjs";
import { Check, X } from "lucide-react";
import { useCallback, useMemo, useState } from "react";
import { toast } from "sonner";
import { useMutation } from "urql";
import { FixHydration } from "~/components/FixHydration";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "~/components/ui/alert-dialog";
import { Badge } from "~/components/ui/badge";
import { Button } from "~/components/ui/button";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "~/components/ui/table";
import { Tooltip, TooltipContent, TooltipTrigger } from "~/components/ui/tooltip";
import { graphql } from "~/gql";
import { Maybe, Payment } from "~/gql/graphql";
import { NonNullish } from "~/types";
import { formatCurrency } from "~/utils/currency";
import { errorMessage } from "~/utils/error";
import { isNotNullish } from "~/utils/value";


type PaymentType = Pick<Payment, "id" | "amount" | "status" | "createdAt"> & {
  customer?: Maybe<Pick<NonNullish<NonNullish<Payment["customer"]>>, "id" | "name">>;
}

interface Props {
  /**
   * The payments to display.
   */
  payments: PaymentType[];
}

const confirmPaymentDocument = graphql(`
  mutation ConfirmPayment($id: ID!) {
    confirmPayment(id: $id) {
      id
    }
  }
`);

const rejectPaymentDocument = graphql(`
  mutation RejectPayment($id: ID!) {
    rejectPayment(id: $id) {
      id
    }
  }
`);

export function PaymentsTable(props: Props) {
  const sorted = useMemo(() => (
    props.payments.sort((a, b) => dayjs(b.createdAt).unix() - dayjs(a.createdAt).unix())
  ), [props.payments]);
  const [confirmPaymentId, setConfirmPaymentId] = useState<string | null>(null);
  const [rejectPaymentId, setRejectPaymentId] = useState<string | null>(null);
  const [{ fetching: fetchingConfirm }, confirmPayment] = useMutation(confirmPaymentDocument);
  const [{ fetching: fetchingReject }, rejectPayment] = useMutation(rejectPaymentDocument);
  const fetching = fetchingConfirm || fetchingReject;
  const handleConfirm = useCallback(async (id: string) => {
    try {
      const { error } = await confirmPayment({ id });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      toast.success("Payment confirmed successfully");
      setConfirmPaymentId(null);
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  }, [confirmPayment]);
  const handleReject = useCallback(async (id: string) => {
    try {
      const { error } = await rejectPayment({ id });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      toast.success("Payment rejected");
      setRejectPaymentId(null);
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  }, [rejectPayment]);

  return (
    <Table>
      <TableCaption>
        {sorted.length > 0 && "A list of all payments"}
        {sorted.length === 0 && "No payments so far"}
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="max-w-24">#</TableHead>
          <TableHead>Customer</TableHead>
          <TableHead>Amount</TableHead>
          <TableHead>Status</TableHead>
          <TableHead className="w-36">Date</TableHead>
          <TableHead></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {sorted.map((payment) => (
          <TableRow key={payment.id}>
            <TableCell className="font-medium">{payment.id}</TableCell>
            <TableCell>{payment.customer?.name ?? "unknown customer"}</TableCell>
            <TableCell>{formatCurrency(payment.amount)}</TableCell>
            <TableCell>
              {payment.status === "PENDING" && (
                <Badge variant="secondary">Pending</Badge>
              )}
              {payment.status === "CONFIRMED" && (
                <Badge variant="success">Confirmed</Badge>
              )}
              {payment.status === "REJECTED" && (
                <Badge variant="destructive">Rejected</Badge>
              )}
            </TableCell>
            <TableCell>
              <FixHydration>
                {dayjs(payment.createdAt).format("DD.MM.YYYY HH:mm")}
              </FixHydration>
            </TableCell>
            <TableCell>
              {payment.status !== "CONFIRMED" && (
                <div className="flex gap-1 justify-end">
                  {payment.status !== "REJECTED" && (
                    <Tooltip>
                      <TooltipTrigger asChild>
                        <Button
                          variant="destructive"
                          size="xs"
                          loading={fetching}
                          icon={<X/>}
                          onClick={() => setConfirmPaymentId(payment.id)}
                        />
                      </TooltipTrigger>
                      <TooltipContent>Reject this payment</TooltipContent>
                    </Tooltip>
                  )}

                  <Tooltip>
                    <TooltipTrigger asChild>
                      <Button
                        variant="success"
                        size="xs"
                        loading={fetching}
                        icon={<Check/>}
                        onClick={() => setConfirmPaymentId(payment.id)}
                      />
                    </TooltipTrigger>
                    <TooltipContent>Confirm this payment</TooltipContent>
                  </Tooltip>
                </div>
              )}
            </TableCell>
          </TableRow>
        ))}

        <AlertDialog open={rejectPaymentId !== null}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Are you sure you want to reject this payment?</AlertDialogTitle>
              <AlertDialogDescription>
                You can still confirm this payment later.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel onClick={() => setConfirmPaymentId(null)}>Cancel</AlertDialogCancel>
              <AlertDialogAction
                variant="destructive"
                onClick={() => isNotNullish(confirmPaymentId) && handleReject(confirmPaymentId)}
              >
                Reject
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>

        <AlertDialog open={confirmPaymentId !== null}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Are you sure you want to confirm this payment?</AlertDialogTitle>
              <AlertDialogDescription>
                This action cannot be undone.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel onClick={() => setConfirmPaymentId(null)}>Cancel</AlertDialogCancel>
              <AlertDialogAction onClick={() => isNotNullish(confirmPaymentId) && handleConfirm(confirmPaymentId)}>
                Confirm
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </TableBody>
    </Table>
  );
}
