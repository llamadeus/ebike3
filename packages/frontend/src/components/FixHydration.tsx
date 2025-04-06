import { ReactNode, Suspense, useEffect, useState } from "react";


interface Props {
  /**
   * The content.
   */
  children: ReactNode;
}

export function FixHydration(props: Props) {
  const [hydrated, setHydrated] = useState(false);

  useEffect(() => {
    setHydrated(true);
  }, []);

  return (
    <Suspense key={hydrated ? "hydrated" : "ssr"}>
      {props.children}
    </Suspense>
  );
}
