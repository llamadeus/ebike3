"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { useMutation } from "urql";
import { z } from "zod";
import { Button } from "~/components/ui/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "~/components/ui/card";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "~/components/ui/form";
import { Input } from "~/components/ui/input";
import { graphql } from "~/gql";
import { errorMessage } from "~/utils/error";
import { isNotNullish } from "~/utils/value";


const schema = z.object({
  username: z.string().min(3, "Your username must be at least 3 characters long."),
  password: z.string().min(6, "Your password must be at least 6 characters long."),
});

const loginDocument = graphql(`
  mutation Login($username: String!, $password: String!) {
    login(username: $username, password: $password) {
      __typename
    }
  }
`);

export function LoginCard() {
  const [{ fetching }, login] = useMutation(loginDocument);
  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const handleLogin = form.handleSubmit(async (data) => {
    try {
      const { error } = await login({ username: data.username, password: data.password });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      window.location.href = "/";
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  });

  return (
    <Form {...form}>
      <form onSubmit={handleLogin}>

        <Card>
          <CardHeader>
            <CardTitle>Welcome to EBike</CardTitle>
            <CardDescription>Great to have you back. Login to start riding!</CardDescription>
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
            <Button variant="ghost" asChild><Link href="/register">Register</Link></Button>
            <Button type="submit" loading={fetching}>Login</Button>
          </CardFooter>
        </Card>
      </form>
    </Form>
  );
}
