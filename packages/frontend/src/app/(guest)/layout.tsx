import { redirect } from "next/navigation";
import { ReactNode } from "react";
import { queryAuth } from "~/features/auth/server";
import { isNotNullish } from "~/utils/value";


interface Props {
  children: ReactNode;
}

export default async function GuestLayout(props: Props) {
  const { data } = await queryAuth();

  if (isNotNullish(data?.auth)) {
    return redirect("/");
  }

  return (
    <div className="flex flex-col w-full max-w-96 mx-auto">
      {props.children}
    </div>
  );
}
