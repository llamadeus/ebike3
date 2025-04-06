import { redirect } from "next/navigation";
import { ReactNode } from "react";
import { queryAuth } from "~/features/auth/server";
import { isNullish } from "~/utils/value";


interface Props {
  children: ReactNode;
}

export default async function AuthenticatedLayout(props: Props) {
  const { data } = await queryAuth();

  if (isNullish(data?.auth)) {
    return redirect("/login");
  }

  return (
    <div className="flex flex-col flex-1 w-full max-w-2xl mx-auto">
      {props.children}
    </div>
  );
}
