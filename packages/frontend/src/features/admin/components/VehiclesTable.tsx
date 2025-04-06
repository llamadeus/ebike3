"use client";

import dayjs from "dayjs";
import { Trash } from "lucide-react";
import { useCallback, useMemo, useState } from "react";
import { toast } from "sonner";
import { useMutation } from "urql";
import { FixHydration } from "~/components/FixHydration";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "~/components/ui/alert-dialog";
import { Badge } from "~/components/ui/badge";
import { Button } from "~/components/ui/button";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "~/components/ui/table";
import { Tooltip, TooltipContent, TooltipTrigger } from "~/components/ui/tooltip";
import { VehicleIndicators } from "~/components/VehicleIndicators";
import { PositionBadge } from "~/components/PositionBadge";
import { graphql } from "~/gql";
import { Maybe, Vehicle } from "~/gql/graphql";
import { NonNullish } from "~/types";
import { errorMessage } from "~/utils/error";
import { isNotNullish, isNullish } from "~/utils/value";


type VehicleType = Pick<Vehicle, "id" | "position" | "battery" | "createdAt"> & {
  activeRental?: Maybe<Pick<NonNullish<Vehicle["activeRental"]>, "id" | "start" | "customerId">>;
};

interface Props {
  /**
   * The vehicles to display.
   */
  vehicles: VehicleType[];
}

const deleteVehicleDocument = graphql(`
  mutation deleteVehicle($id: ID!) {
    deleteVehicle(id: $id) {
      id
    }
  }
`);

export function VehiclesTable(props: Props) {
  const sorted = useMemo(() => (
    props.vehicles.sort((a, b) => {
      if (isNotNullish(a.activeRental)) {
        return -1;
      }
      if (isNotNullish(b.activeRental)) {
        return 1;
      }
      return 0;
    })
  ), [props.vehicles]);
  const [deleteVehicleId, setDeleteVehicleId] = useState<string | null>(null);
  const [{ fetching }, deleteVehicle] = useMutation(deleteVehicleDocument);
  const handleDelete = useCallback(async (id: string) => {
    try {
      const { error } = await deleteVehicle({ id });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      toast.success("Vehicle deleted successfully");
      setDeleteVehicleId(null);
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  }, [deleteVehicle]);

  return (
    <Table>
      <TableCaption>
        {sorted.length > 0 && "A list of vehicles"}
        {sorted.length === 0 && "No vehicles created so far"}
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="max-w-24">#</TableHead>
          <TableHead>Position</TableHead>
          <TableHead>Details</TableHead>
          <TableHead>Active rental</TableHead>
          <TableHead className="w-36">Created at</TableHead>
          <TableHead></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {sorted.map((vehicle) => (
          <TableRow key={vehicle.id}>
            <TableCell className="font-medium">{vehicle.id}</TableCell>
            <TableCell>
              <PositionBadge position={vehicle.position}/>
            </TableCell>
            <TableCell>
              <VehicleIndicators vehicle={vehicle}/>
            </TableCell>
            <TableCell>
              {isNotNullish(vehicle.activeRental?.customerId) && (
                <Tooltip>
                  <TooltipTrigger>
                    <Badge variant="default">{vehicle.activeRental.customerId}</Badge>
                  </TooltipTrigger>
                  <TooltipContent>
                    <FixHydration>
                      Rental started on
                      {" "}
                      {dayjs(vehicle.activeRental.start).format("DD.MM.YYYY")}
                      {" "}
                      at
                      {" "}
                      {dayjs(vehicle.activeRental.start).format("HH:mm")}
                    </FixHydration>
                  </TooltipContent>
                </Tooltip>
              )}
              {isNullish(vehicle.activeRental) && <>&ndash;</>}
            </TableCell>
            <TableCell>
              <FixHydration>
                {dayjs(vehicle.createdAt).format("DD.MM.YYYY HH:mm")}
              </FixHydration>
            </TableCell>
            <TableCell className="text-right">
              <Tooltip>
                <TooltipTrigger asChild>
                  <Button
                    variant="ghost"
                    size="xs"
                    loading={fetching}
                    icon={<Trash/>}
                    onClick={() => setDeleteVehicleId(vehicle.id)}
                  />
                </TooltipTrigger>
                <TooltipContent>Delete this vehicle</TooltipContent>
              </Tooltip>
            </TableCell>
          </TableRow>
        ))}

        <AlertDialog open={deleteVehicleId !== null}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Are you sure you want to delete this vehicle?</AlertDialogTitle>
              <AlertDialogDescription>
                This action cannot be undone.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel onClick={() => setDeleteVehicleId(null)}>Cancel</AlertDialogCancel>
              <AlertDialogAction
                onClick={() => isNotNullish(deleteVehicleId) && handleDelete(deleteVehicleId)}
                variant="destructive"
              >
                Delete
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </TableBody>
    </Table>
  );
}
