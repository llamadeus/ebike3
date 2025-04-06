import { Vec2d } from "~/gql/graphql";


interface Props {
  /**
   * The vehicle to display the position for.
   */
  position: Vec2d;
}

export function PositionBadge(props: Props) {
  return (
    <div className="flex items-center">
      <div className="rounded-full bg-foreground/10 px-2">
        {Math.round(props.position.x * 100) / 100}
        {"; "}
        {Math.round(props.position.y * 100) / 100}
      </div>
    </div>
  );
}
