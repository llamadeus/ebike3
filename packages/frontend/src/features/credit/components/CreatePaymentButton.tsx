"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { useMutation } from "urql";
import { z } from "zod";
import { Button } from "~/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "~/components/ui/dialog";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "~/components/ui/form";
import { Input } from "~/components/ui/input";
import { graphql } from "~/gql";
import { errorMessage } from "~/utils/error";
import { isNotNullish } from "~/utils/value";


const schema = z.object({
  amount: z.string().regex(/^\d+(\.\d+)?$/, "Invalid number"),
});

const createPaymentDocument = graphql(`
  mutation CreatePayment($amount: Int!) {
    createPayment(amount: $amount) {
      id
    }
  }
`);

export function CreatePaymentButton() {
  const [showModal, setShowModal] = useState(false);
  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: {
      amount: "",
    },
  });
  const [{ fetching }, createPayment] = useMutation(createPaymentDocument);

  const handleAdd = form.handleSubmit(async ({ amount }) => {
    try {
      const { error } = await createPayment({
        amount: Number(amount.replace(",", ".")) * 100,
      });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      toast.success("Payment created");
      setShowModal(false);
      form.reset();
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  });

  return (
    <>
      <Button size="sm" onClick={() => setShowModal(true)}>Recharge credit</Button>
      <Dialog open={showModal} onOpenChange={setShowModal}>
        <DialogContent onPointerDownOutside={event => event.preventDefault()}>
          <Form {...form}>
            <form onSubmit={handleAdd}>
              <DialogHeader>
                <DialogTitle>Recharge your credit balance</DialogTitle>
                <DialogDescription>
                  Send a payment to our bank account to recharge your credit balance and fill out the form below.
                </DialogDescription>
              </DialogHeader>

              <div className="py-4">
                <FormField
                  control={form.control}
                  name="amount"
                  render={({ field, fieldState }) => (
                    <div className="space-y-2">
                      <FormItem className="grid grid-cols-4 items-center gap-4">
                        <FormLabel>Amount</FormLabel>

                        <FormControl className="flex gap-2 col-span-3">
                          <Input type="number" step="0.01" {...field}/>
                        </FormControl>
                      </FormItem>

                      {isNotNullish(fieldState.error) && (
                        <div className="grid grid-cols-4 items-center gap-4">
                          <div className="col-start-2 col-span-3">
                            <FormMessage/>
                          </div>
                        </div>
                      )}
                    </div>
                  )}
                />
              </div>

              <DialogFooter>
                <DialogClose asChild>
                  <Button variant="outline">
                    Cancel
                  </Button>
                </DialogClose>
                <Button type="submit" loading={fetching}>Confirm</Button>
              </DialogFooter>
            </form>
          </Form>
        </DialogContent>
      </Dialog>
    </>
  );
}
