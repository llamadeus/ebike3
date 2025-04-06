import { ReactNode } from "react";


interface Props {
  /**
   * The caption of the section.
   */
  caption: string;
  /**
   * An optional right-aligned element.
   */
  right?: ReactNode;

  /**
   * The content of the section.
   */
  children: ReactNode;
}

export function Section(props: Props) {
  return (
    <div className="flex flex-col gap-4">
      <div className="min-h-10 flex items-center gap-4">
        <h2 className="flex-1 text-xl">{props.caption}</h2>
        {props.right}
      </div>

      {props.children}
    </div>
  );
}
