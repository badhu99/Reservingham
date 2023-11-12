import { Route } from "react-router-dom";
import Dashboard from "../../pages/dashboard/dashboard";
import { Roles } from "../../utils/enums/roles";
import AuthGuard from "../../utils/helpers/guards/auth-guard";

export default [
    <Route
    path="/dashboard"
    element={
      <AuthGuard allowedRoles={[Roles.User]} element={<Dashboard />} />
    }
  />
]