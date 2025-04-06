import { Tabs, TabsContent, TabsList, TabsTrigger } from "~/components/ui/tabs";
import { RegisterAdminCard } from "~/features/auth/components/RegisterAdminCard";
import { RegisterCustomerCard } from "~/features/auth/components/RegisterCustomerCard";


export default function Register() {
  return (
    <Tabs defaultValue="customer" className="flex flex-col w-full">
      <TabsList>
        <TabsTrigger value="customer" className="flex-1">Customer</TabsTrigger>
        <TabsTrigger value="admin" className="flex-1">Admin</TabsTrigger>
      </TabsList>
      <TabsContent value="customer">
        <RegisterCustomerCard/>
      </TabsContent>

      <TabsContent value="admin">
        <RegisterAdminCard/>
      </TabsContent>
    </Tabs>
  );
}
