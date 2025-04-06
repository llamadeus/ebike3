import { useCallback } from "react";
import { toast } from "sonner";
import { useMutation } from "urql";
import { graphql } from "~/gql";
import { errorMessage } from "~/utils/error";
import { isNotNullish } from "~/utils/value";


const logoutDocument = graphql(`
  mutation Logout {
    logout
  }
`);

/**
 * A hook that provides a logout function.
 */
export function useLogout() {
  const [{ fetching }, logout] = useMutation(logoutDocument);
  const handleLogout = useCallback(async () => {
    try {
      const { error } = await logout({});
      if (isNotNullish(error)) {
        toast.error(errorMessage(error));
        return;
      }

      window.location.href = "/";
    }
    catch (err) {
      toast.error(errorMessage(err));
    }
  }, [logout]);

  return {
    logout: handleLogout,
    fetching,
  };
}
