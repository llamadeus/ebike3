import { redirect } from "next/navigation";
import { Card, CardDescription, CardFooter, CardHeader, CardTitle } from "~/components/ui/card";
import { LogoutButton } from "~/features/auth/components/LogoutButton";
import { queryAuth } from "~/features/auth/server";


export default async function Home() {
  const { data } = await queryAuth();

  if (data?.auth?.role === "ADMIN") {
    return redirect("/admin");
  }

  if (data?.auth?.role === "CUSTOMER") {
    return redirect("/customer");
  }

  return (
    <div className="flex flex-col flex-1 justify-center">
      <Card className="max-w-96 mx-auto">
        <CardHeader className="flex flex-1">
          <CardTitle>Ooopsie...</CardTitle>
          <CardDescription>This should never happen, yet it did. Please logout and try again.</CardDescription>
        </CardHeader>

        <CardFooter className="flex justify-end gap-2">
          <LogoutButton variant="destructive"/>
        </CardFooter>
      </Card>
    </div>
  );
}
