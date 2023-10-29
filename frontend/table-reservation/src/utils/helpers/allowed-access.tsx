import { Navigate } from "react-router-dom";
import { Roles } from "../enums/roles"
import useToken from "../../hooks/useToken"
import { UserData } from "../../classes/user_data";

export default function AllowedAccess(allowedRoles: Roles[], element: JSX.Element) {

    const {tokenData} = useToken()

    return allowedRoles.find((role) => tokenData.HasValue && tokenData.Roles.includes(role)) ? 
    ( element ) : tokenData.HasValue ? 
        (<Navigate to="/not-found" replace  />) : 
        (<Navigate to="/auth/signin"  />);
}

export function AllowedAccessBool(
    allowedRoles: Roles[],
    userData: UserData
  ): boolean {
    const val = allowedRoles.find(
      (role) => userData.HasValue && userData.Roles.includes(role)
    );
    if (val !== undefined)
      return true;
    return false;
  }