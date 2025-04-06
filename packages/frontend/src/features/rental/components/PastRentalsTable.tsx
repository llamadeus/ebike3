import dayjs from "dayjs";
import { useMemo } from "react";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "~/components/ui/table";
import { Rental } from "~/gql/graphql";
import { formatCurrency } from "~/utils/currency";
import { formatDifference } from "~/utils/dayjs";


type RentalType = Pick<Rental, "id" | "start" | "end" | "cost">;

interface Props {
  /**
   * The rentals to display.
   */
  rentals: RentalType[];
}

export function PastRentalsTable(props: Props) {
  const sorted = useMemo(() => (
    props.rentals.sort((a, b) => dayjs(b.start).unix() - dayjs(a.start).unix())
  ), [props.rentals]);

  return (
    <Table>
      <TableCaption>
        {sorted.length > 0 && "Your rental history"}
        {sorted.length === 0 && "No past rentals"}
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Date</TableHead>
          <TableHead>Duration</TableHead>
          <TableHead>Cost</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {sorted.map((rental) => {
          return (
            <TableRow key={rental.id}>
              <TableCell>{dayjs(rental.start).format("DD.MM.YYYY")}</TableCell>
              <TableCell>{formatDifference(dayjs(rental.start).toDate(), dayjs(rental.end).toDate())}</TableCell>
              <TableCell className="text-right">{formatCurrency(rental.cost ?? 0)}</TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}
