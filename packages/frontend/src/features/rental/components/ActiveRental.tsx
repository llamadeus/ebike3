import dayjs from "dayjs";
import { Bike } from "lucide-react";
import { useCallback } from "react";
import { toast } from "sonner";
import { useMutation } from "urql";
import { FixHydration } from "~/components/FixHydration";
import { Button } from "~/components/ui/button";
import { graphql } from "~/gql";
import { Maybe, Rental } from "~/gql/graphql";
import { NonNullish } from "~/types";
import { formatCurrency } from "~/utils/currency";
import { formatDifference } from "~/utils/dayjs";
import { errorMessage } from "~/utils/error";
import { isNotNullish } from "~/utils/value";


type RentalType = Pick<Rental, "id" | "start" | "cost">;

interface Props {
  /**
   * The active rental.
   */
  rental: RentalType;
}

const stopRentalDocument = graphql(`
  mutation StopRental($id: ID!) {
    stopRental(id: $id) {
      id
    }
  }
`);

export function ActiveRental(props: Props) {
  const [{ fetching }, stopRental] = useMutation(stopRentalDocument);
  const handleStop = useCallback(async () => {
    try {
      const { error } = await stopRental({ id: props.rental.id });
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      toast.success("Rental stopped");
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  }, [props.rental.id, stopRental]);

  return (
    <div className="flex gap-6">
      <Bike size="6rem"/>

      <div className="flex flex-col flex-1 justify-center gap-2">
        <FixHydration>
          <div>
            <div><b>Duration:</b> {formatDifference(dayjs(props.rental.start).toDate(), new Date())}</div>
            <div><b>Cost:</b> {formatCurrency(props.rental.cost ?? 0)}</div>
          </div>

          <div>
            <Button size="sm" loading={fetching} onClick={handleStop}>Stop rental</Button>
          </div>
        </FixHydration>
      </div>
    </div>
  );
}
