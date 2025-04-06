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
  name: z.string().min(1, "The name must be at least 1 characters long."),
  position: z.object({
    x: z.string(),
    y: z.string(),
  })
    .refine(data => /^-?\d+(\.\d+)?$/.test(data.x), "Invalid x coordinate")
    .refine(data => /^-?\d+(\.\d+)?$/.test(data.y), "Invalid y coordinate"),
});

const createStationDocument = graphql(`
  mutation CreateStation($input: CreateStationInput!) {
    createStation(input: $input) {
      id
    }
  }
`);

export function CreateStationButton() {
  const [showModal, setShowModal] = useState(false);
  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: {
      name: "",
      position: {
        x: "",
        y: "",
      },
    },
  });
  const [{ fetching }, createStation] = useMutation(createStationDocument);

  const handleAdd = form.handleSubmit(async (values) => {
    try {
      const { error } = await createStation({
        input: {
          name: values.name,
          position: {
            x: Number(values.position.x),
            y: Number(values.position.y),
          },
        },
      });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      toast.success("Station created");
      setShowModal(false);
      form.reset();
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  });

  return (
    <>
      <Button size="sm" onClick={() => setShowModal(true)}>Create station</Button>
      <Dialog open={showModal} onOpenChange={setShowModal}>
        <DialogContent onPointerDownOutside={event => event.preventDefault()}>
          <Form {...form}>
            <form onSubmit={handleAdd}>
              <DialogHeader>
                <DialogTitle>Create a new station</DialogTitle>
                <DialogDescription>
                  Create a new station for your vehicles.
                </DialogDescription>
              </DialogHeader>

              <div className="py-4">
                <FormField
                  control={form.control}
                  name="name"
                  render={({ field, fieldState }) => (
                    <div className="space-y-2">
                      <FormItem className="grid grid-cols-4 items-center gap-4">
                        <FormLabel>Name</FormLabel>
                        <FormControl className="flex flex-col gap-2 col-span-3">
                          <Input {...field} />
                        </FormControl>
                      </FormItem>

                      {isNotNullish(fieldState.error) && (
                        <div className="grid grid-cols-4 items-center gap-4">
                          <FormMessage className="col-start-2 col-span-3"/>
                        </div>
                      )}
                    </div>
                  )}
                />

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
