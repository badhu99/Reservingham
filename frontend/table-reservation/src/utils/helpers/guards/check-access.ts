import { UserData } from "../../../classes/user_data";
import { Roles } from "../../enums/roles";

export function AllowedAccessBool(
    allowedRoles: Roles[],
    userData: UserData
  ): boolean {
    const val = allowedRoles.find(
      (role) => userData.HasValue && userData.Roles.includes(role)
    );
    if (val !== undefined) return true;
    return false;
  }