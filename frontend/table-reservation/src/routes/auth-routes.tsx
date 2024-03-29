import { Route, Navigate } from "react-router-dom";
import Admin from "../pages/admin/admin";
import Navbar from "../pages/common/navbar";
import Dashboard from "../pages/dashboard/dashboard";
import Editor from "../pages/editor/editor";
import Reservations from "../pages/reservations/reservations";
import { Roles } from "../utils/enums/roles";
import AuthGuard from "../utils/helpers/guards/auth-guard";
import userRoutes from "./auth-routes/user-management-routes";
import dashboardRoutes from "./auth-routes/dashboard-routes";
import reservationsRoutes from "./auth-routes/reservations-routes";
import editorRoutes from "./auth-routes/editor-routes";
import adminRoutes from "./auth-routes/admin-routes";

const authRoutes =  [
  <Route
    path=""
    element={
      <AuthGuard
        allowedRoles={[Roles.Admin, Roles.Editor, Roles.User]}
        element={<Navbar />}
      />
    }
  >
    <Route path="" element={<Navigate to="/dashboard" />} />
    {dashboardRoutes}
    {reservationsRoutes}
    {editorRoutes}
    {userRoutes}
    {adminRoutes}
  </Route>,
];

export default authRoutes;