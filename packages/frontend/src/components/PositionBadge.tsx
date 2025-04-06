import { Vec2d } from "~/gql/graphql";
import { cn } from "~/lib/utils";


interface Props {
  /**
   * The vehicle to display the position for.
   */
  position: Vec2d;

  /**
   * The class name to apply to the badge.
   */
  className?: string;
}

export function PositionBadge(props: Props) {
  return (
    <div className={cn("flex items-center", props.className)}>
      <div className="rounded-full bg-foreground/10 px-2">
        {Math.round(props.position.x * 100) / 100}
        {"; "}
        {Math.round(props.position.y * 100) / 100}
      </div>
    </div>
  );
}
