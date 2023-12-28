import { Route } from "react-router-dom";
import Manager from "../../pages/manager/manager";
import { ManagerSinglePage } from "../../pages/manager/manager-single-page";
import { Roles } from "../../utils/enums/roles";
import AuthGuard from "../../utils/helpers/guards/auth-guard";

const UserManagementRouters = [
  <Route path="users">
    <Route
      path=""
      element={
        <AuthGuard allowedRoles={[Roles.Manager]} element={<Manager />} />
      }
    />
    <Route
      path=":id"
      element={
        <AuthGuard
          allowedRoles={[Roles.Manager]}
          element={<ManagerSinglePage />}
        />
      }
    />
  </Route>,
];
export default UserManagementRouters;
