import type { Metadata } from "next";
import { ReactNode, Suspense } from "react";
import "./global.css";
import { Providers } from "~/components/Providers";
import { Toaster } from "~/components/ui/sonner";


export const metadata: Metadata = {
  title: "EBike",
  description: "EBike application",
};

interface Props {
  children: ReactNode;
}

export default async function RootLayout(props: Props) {
  return (
    <html lang="en">
    <body>
      <Providers>
        <Suspense>
          <div className="flex flex-col flex-1 py-8">
            {props.children}
          </div>
        </Suspense>
      </Providers>
      <Toaster/>
    </body>
    </html>
  );
}
