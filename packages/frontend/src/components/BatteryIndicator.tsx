"use client";

import { Battery, BatteryFull, BatteryLow, BatteryMedium, BatteryWarning } from "lucide-react";
import { useEffect, useMemo, useState } from "react";
import { Tooltip, TooltipContent, TooltipTrigger } from "~/components/ui/tooltip";


interface Props {
  /**
   * The battery level of the vehicle.
   */
  battery: number;
}

export function BatteryIndicator(props: Props) {
  const [warning, setWarning] = useState(false);
  const icon = useMemo(() => {
    if (props.battery > 0.65) {
      return <BatteryFull/>;
    }

    if (props.battery > 0.35) {
      return <BatteryMedium/>;
    }

    if (props.battery > 0.05) {
      return <BatteryLow/>;
    }

    if (warning) {
      return <BatteryWarning/>;
    }

    return <Battery/>;
  }, [props.battery, warning]);

  useEffect(() => {
    if (props.battery > 0.05) {
      return;
    }

    const interval = setInterval(() => {
      setWarning(prev => ! prev);
    }, 1000);

    return () => clearInterval(interval);
  }, [props.battery]);

  return (
    <Tooltip>
      <TooltipTrigger>{icon}</TooltipTrigger>
      <TooltipContent>
        Battery: {Math.round(props.battery * 100)}%
      </TooltipContent>
    </Tooltip>
  );
}
