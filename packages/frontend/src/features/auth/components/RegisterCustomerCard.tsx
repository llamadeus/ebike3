"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { useMutation } from "urql";
import { z } from "zod";
import { Button } from "~/components/ui/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "~/components/ui/card";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "~/components/ui/form";
import { Input } from "~/components/ui/input";
import { graphql } from "~/gql";
import { errorMessage } from "~/utils/error";
import { isNotNullish } from "~/utils/value";


const schema = z.object({
  username: z.string().min(3, "Your username must be at least 3 characters long."),
  password: z.string().min(6, "Your password must be at least 6 characters long."),
});

const registerCustomerDocument = graphql(`
  mutation RegisterCustomer($username: String!, $password: String!) {
    registerCustomer(username: $username, password: $password) {
      id
    }
  }
`);

export function RegisterCustomerCard() {
  const router = useRouter();
  const [{ fetching }, registerCustomer] = useMutation(registerCustomerDocument);
  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const handleRegister = form.handleSubmit(async (data) => {
    try {
      const { error } = await registerCustomer({ username: data.username, password: data.password });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      toast.success("Customer registered successfully");
      window.location.href = "/";
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  });

  return (
    <Form {...form}>
      <form onSubmit={handleRegister}>

        <Card>
          <CardHeader>
            <CardTitle>EBike / Customer</CardTitle>
            <CardDescription>We are happy about every new customer. Register now and start riding!</CardDescription>
          </CardHeader>
          <CardContent className="flex flex-col gap-4">
            <FormField
              control={form.control}
              name="username"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Username</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormDescription>You can&apos;t change your username after registration.</FormDescription>
                  <FormMessage/>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input type="password" {...field} />
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />
          </CardContent>
          <CardFooter className="flex justify-end gap-2">
            <Button variant="ghost" asChild><Link href="/login">Login</Link></Button>
            <Button type="submit" loading={fetching}>Register</Button>
          </CardFooter>
        </Card>
      </form>
    </Form>
  );
}
