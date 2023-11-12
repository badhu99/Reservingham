import { Route } from "react-router-dom";
import Admin from "../../pages/admin/admin";
import { Roles } from "../../utils/enums/roles";
import AuthGuard from "../../utils/helpers/guards/auth-guard";

export default [
    <Route
    path="/admin"
    element={<AuthGuard allowedRoles={[Roles.Admin]} element={<Admin />} />}
  />
]