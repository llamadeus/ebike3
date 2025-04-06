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
import { VehicleType } from "~/gql/graphql";
import { errorMessage } from "~/utils/error";
import { isNotNullish } from "~/utils/value";


const schema = z.object({
  type: z.enum(["BIKE", "EBIKE", "ABIKE"]),
  position: z.object({
    x: z.string(),
    y: z.string(),
  })
    .refine(data => /^-?\d+(\.\d+)?$/.test(data.x), "Invalid x coordinate")
    .refine(data => /^-?\d+(\.\d+)?$/.test(data.y), "Invalid y coordinate"),
});

const createVehicleDocument = graphql(`
  mutation CreateVehicle($input: CreateVehicleInput!) {
    createVehicle(input: $input) {
      id
    }
  }
`);

export function CreateVehicleButton() {
  const [showModal, setShowModal] = useState(false);
  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: {
      type: VehicleType.Bike,
      position: {
        x: "",
        y: "",
      },
    },
  });
  const [{ fetching }, createVehicle] = useMutation(createVehicleDocument);

  const handleAdd = form.handleSubmit(async ({ position }) => {
    try {
      const { error } = await createVehicle({
        input: {
          type: VehicleType.Bike,
          position: {
            x: Number(position.x),
            y: Number(position.y),
          },
        },
      });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      toast.success("Vehicle created");
      setShowModal(false);
      form.reset();
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  });

  return (
    <>
      <Button size="sm" onClick={() => setShowModal(true)}>Create vehicle</Button>
      <Dialog open={showModal} onOpenChange={setShowModal}>
        <DialogContent onPointerDownOutside={event => event.preventDefault()}>
          <Form {...form}>
            <form onSubmit={handleAdd}>
              <DialogHeader>
                <DialogTitle>Create a new vehicle</DialogTitle>
                <DialogDescription>
                  Create a new vehicle and make it available for rentals.
                </DialogDescription>
              </DialogHeader>

              <div className="py-4">
                <div className="text-red-500 bold">
                  TODO: HIER SELECT HIN, UM TYPE ZU WÄHLEN
                </div>

                <FormField
                  control={form.control}
                  name="position"
                  render={({ field, fieldState }) => (
                    <div className="space-y-2">
                      <FormItem className="grid grid-cols-4 items-center gap-4">
                        <FormLabel>Position</FormLabel>

                        <div className="flex gap-2 col-span-3">
                          <FormControl>
                            <Input
                              {...field}
                              ref={undefined}
                              value={field.value.x}
                              onChange={event => field.onChange({
                                ...field.value,
                                x: event.target.value,
                              })}
                            />
                          </FormControl>
                          <FormControl>
                            <Input
                              {...field}
                              ref={undefined}
                              value={field.value.y}
                              onChange={event => field.onChange({
                                ...field.value,
                                y: event.target.value,
                              })}
                            />
                          </FormControl>
                        </div>
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
                <Button type="submit" loading={fetching}>Create</Button>
              </DialogFooter>
            </form>
          </Form>
        </DialogContent>
      </Dialog>
    </>
  );
}
