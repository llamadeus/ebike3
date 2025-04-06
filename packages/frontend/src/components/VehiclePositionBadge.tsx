import { Vehicle } from "~/gql/graphql";


type VehicleType = Pick<Vehicle, "position">;

interface Props {
  /**
   * The vehicle to display the position for.
   */
  vehicle: VehicleType;
}

export function VehiclePositionBadge(props: Props) {
  return (
    <div className="flex items-center">
      <div className="rounded-full bg-foreground/10 px-2">
        {Math.round(props.vehicle.position.x * 100) / 100}
        {"; "}
        {Math.round(props.vehicle.position.y * 100) / 100}
      </div>
    </div>
  );
}
