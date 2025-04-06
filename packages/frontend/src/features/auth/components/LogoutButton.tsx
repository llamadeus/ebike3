"use client";

import { Button, ButtonProps } from "~/components/ui/button";
import { useLogout } from "~/features/auth/hooks/use-logout";


export function LogoutButton(props: Omit<ButtonProps, "loading">) {
  const { children, ...other } = props;
  const { logout, fetching } = useLogout();

  return (
    <Button variant="ghost" onClick={logout} loading={fetching} {...other}>
      {children ?? "Logout"}
    </Button>
  );
}
