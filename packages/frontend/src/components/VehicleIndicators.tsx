import { BatteryIndicator } from "~/components/BatteryIndicator";
import { Vehicle } from "~/gql/graphql";


type VehicleType = Pick<Vehicle, "battery">;

interface Props {
  /**
   * The vehicle to display the indicators for.
   */
  vehicle: VehicleType;
}

export function VehicleIndicators(props: Props) {
  return (
    <div className="flex gap-1">
      <BatteryIndicator battery={props.vehicle.battery}/>
    </div>
  );
}
