"use client";

import dayjs from "dayjs";
import { Trash } from "lucide-react";
import { useCallback, useMemo, useState } from "react";
import { toast } from "sonner";
import { useMutation } from "urql";
import { FixHydration } from "~/components/FixHydration";
import { PositionBadge } from "~/components/PositionBadge";
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
import { graphql } from "~/gql";
import { Station } from "~/gql/graphql";
import { errorMessage } from "~/utils/error";
import { isNotNullish, isNullish } from "~/utils/value";


type StationType = Pick<Station, "id" | "name" | "position" | "createdAt">;

interface Props {
  /**
   * The stations to display.
   */
  stations: StationType[];
}

const deleteStationDocument = graphql(`
  mutation deleteStation($id: ID!) {
    deleteStation(id: $id) {
      id
    }
  }
`);

export function StationsTable(props: Props) {
  const sorted = useMemo(() => (
    props.stations.sort((a, b) => a.name.localeCompare(b.name))
  ), [props.stations]);
  const [deleteStationId, setDeleteStationId] = useState<string | null>(null);
  const [{ fetching }, deleteStation] = useMutation(deleteStationDocument);
  const handleDelete = useCallback(async (id: string) => {
    try {
      const { error } = await deleteStation({ id });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      toast.success("Station deleted successfully");
      setDeleteStationId(null);
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  }, [deleteStation]);

  return (
    <Table>
      <TableCaption>
        {sorted.length > 0 && "A list of stations"}
        {sorted.length === 0 && "No stations created so far"}
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="max-w-24">#</TableHead>
          <TableHead>Name</TableHead>
          <TableHead>Position</TableHead>
          <TableHead className="w-36">Created at</TableHead>
          <TableHead></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {sorted.map((station) => (
          <TableRow key={station.id}>
            <TableCell className="font-medium">{station.id}</TableCell>
            <TableCell>{station.name}</TableCell>
            <TableCell>
              <PositionBadge position={station.position}/>
            </TableCell>
            <TableCell>
              <FixHydration>
                {dayjs(station.createdAt).format("DD.MM.YYYY HH:mm")}
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
                    onClick={() => setDeleteStationId(station.id)}
                  />
                </TooltipTrigger>
                <TooltipContent>Delete this station</TooltipContent>
              </Tooltip>
            </TableCell>
          </TableRow>
        ))}

        <AlertDialog open={deleteStationId !== null}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Are you sure you want to delete this station?</AlertDialogTitle>
              <AlertDialogDescription>
                This action cannot be undone.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel onClick={() => setDeleteStationId(null)}>Cancel</AlertDialogCancel>
              <AlertDialogAction
                onClick={() => isNotNullish(deleteStationId) && handleDelete(deleteStationId)}
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
