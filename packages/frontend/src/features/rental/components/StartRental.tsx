import { ArrowBigRightDash } from "lucide-react";
import { useCallback } from "react";
import { toast } from "sonner";
import { useMutation } from "urql";
import { Button } from "~/components/ui/button";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "~/components/ui/table";
import { Tooltip, TooltipContent, TooltipTrigger } from "~/components/ui/tooltip";
import { VehicleIndicators } from "~/components/VehicleIndicators";
import { VehiclePositionBadge } from "~/components/VehiclePositionBadge";
import { graphql } from "~/gql";
import { Vehicle } from "~/gql/graphql";
import { errorMessage } from "~/utils/error";
import { isNotNullish } from "~/utils/value";


type VehicleType = Pick<Vehicle, "id" | "battery" | "position">;

interface Props {
  /**
   * The vehicles to display.
   */
  vehicles: VehicleType[];
}

const startRentalDocument = graphql(`
  mutation StartRental($vehicleId: ID!) {
    startRental(vehicleId: $vehicleId) {
      id
    }
  }
`);

export function StartRental(props: Props) {
  const [{ fetching }, startRental] = useMutation(startRentalDocument);
  const handleStart = useCallback(async (vehicleId: string) => {
    try {
      const { error } = await startRental({ vehicleId });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      toast.success("Rental started");
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  }, [startRental]);

  return (
    <Table>
      <TableCaption>All available vehicles</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="max-w-24">#</TableHead>
          <TableHead>Position</TableHead>
          <TableHead>Details</TableHead>
          <TableHead></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {props.vehicles.map((vehicle) => (
          <TableRow key={vehicle.id}>
            <TableCell className="font-medium">{vehicle.id}</TableCell>
            <TableCell>
              <VehiclePositionBadge vehicle={vehicle}/>
            </TableCell>
            <TableCell>
              <VehicleIndicators vehicle={vehicle}/>
            </TableCell>
            <TableCell className="text-right">
              <Tooltip>
                <TooltipTrigger asChild>
                  <Button
                    size="xs"
                    loading={fetching}
                    icon={<ArrowBigRightDash/>}
                    onClick={() => handleStart(vehicle.id)}
                  />
                </TooltipTrigger>
                <TooltipContent>Rent this vehicle</TooltipContent>
              </Tooltip>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
