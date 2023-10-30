import { Navigate } from "react-router-dom";
import { Roles } from "../../enums/roles";
import useToken from "../../../hooks/useToken";
import { useEffect, useState } from "react";

interface IGuard {
  allowedRoles: Roles[];
  element: JSX.Element;
}

export default function AuthGuard({ allowedRoles, element }: IGuard) {
  const [allowed, setAllowed] = useState(0);
  const { tokenData } = useToken();
  useEffect(() => {
    check();
  }, [tokenData, allowedRoles, element]);

  const check = () => {
    const val = allowedRoles.find(
      (role) => tokenData.HasValue && tokenData.Roles.includes(role)
    );
    if (val !== undefined) setAllowed(1);
    else setAllowed(2);
  };

  if (allowed === 0) return <></>
  else if (allowed === 1) return element;

  return tokenData.HasValue ? 
    <Navigate to="/not-found" replace /> : 
    <Navigate to="/auth/signin" replace />
}