import dayjs from "dayjs";
import { User, UserCog } from "lucide-react";
import { FixHydration } from "~/components/FixHydration";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "~/components/ui/table";
import { Tooltip, TooltipContent, TooltipTrigger } from "~/components/ui/tooltip";
import { User as UserType } from "~/gql/graphql";
import { isNotNullish } from "~/utils/value";


interface Props {
  /**
   * The users to display.
   */
  users: Pick<UserType, "id" | "role" | "username" | "lastLogin">[];
  /**
   * The current user ID so we can highlight the current user.
   */
  userId: string;
}

export function UsersTable(props: Props) {
  const sorted = props.users.sort((a, b) => a.role.localeCompare(b.role)) ?? [];

  return (
    <Table>
      <TableCaption>
        {sorted.length > 0 && "A list of registered users"}
        {sorted.length === 0 && "No users registered"}
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="max-w-24">#</TableHead>
          <TableHead>Name</TableHead>
          <TableHead>Details</TableHead>
          <TableHead className="w-36">Last login</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {sorted.map((user) => {
          return (
            <TableRow key={user.id}>
              <TableCell className="font-medium">{user.id}</TableCell>
              <TableCell>
                {user.id === props.userId && <b>{user.username}</b>}
                {user.id !== props.userId && user.username}
              </TableCell>
              <TableCell>
                <div className="flex gap-1">
                  {user.role === "ADMIN" && (
                    <Tooltip>
                      <TooltipTrigger><UserCog/></TooltipTrigger>
                      <TooltipContent>Admin</TooltipContent>
                    </Tooltip>
                  )}
                  {user.role === "CUSTOMER" && (
                    <>
                      <Tooltip>
                        <TooltipTrigger><User/></TooltipTrigger>
                        <TooltipContent>Customer</TooltipContent>
                      </Tooltip>
                    </>
                  )}
                </div>
              </TableCell>
              <TableCell>
                <FixHydration>
                  {isNotNullish(user.lastLogin) ? dayjs(user.lastLogin).format("DD.MM.YYYY HH:mm") : "n/a"}
                </FixHydration>
              </TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}
