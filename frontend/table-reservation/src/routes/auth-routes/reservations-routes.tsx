import { Route } from "react-router-dom";
import Reservations from "../../pages/reservations/reservations";
import { Roles } from "../../utils/enums/roles";
import AuthGuard from "../../utils/helpers/guards/auth-guard";

export default [
    <Route
    path="/reservations"
    element={
      <AuthGuard allowedRoles={[Roles.User]} element={<Reservations />} />
    }
  />
]