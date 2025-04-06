"use client";

import { UrqlProvider } from "@urql/next";
import { ReactNode, useMemo } from "react";
import { ssrExchange } from "urql";
import { TooltipProvider } from "~/components/ui/tooltip";
import { makeClient } from "~/urql/client";


interface Props {
  /**
   * Content to render within the providers.
   */
  children: ReactNode;
}

export function Providers(props: Props) {
  const [client, ssr] = useMemo(() => {
    const ssr = ssrExchange({ isClient: typeof window != "undefined" });
    const client = makeClient(ssr);

    return [client, ssr];
  }, []);

  return (
    <UrqlProvider client={client} ssr={ssr}>
      <TooltipProvider>
        {props.children}
      </TooltipProvider>
    </UrqlProvider>
  );
}
