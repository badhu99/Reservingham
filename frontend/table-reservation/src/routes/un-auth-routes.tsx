import { Navigate, Route } from "react-router-dom";
import BackboneAuthentication from "../pages/authentication/backbone-auth";
import Login from "../pages/authentication/login";
import Registration from "../pages/authentication/registration";
import UnAuthGuard from "../utils/helpers/guards/un-auth-guard";

const UnAuthRoutes = [
    <Route
    path="auth"
    element={<UnAuthGuard element={<BackboneAuthentication />} />}
  >
    <Route path="" element={<Navigate to="/auth/signin" />} />
    <Route
      key="Login"
      path="signin"
      element={<UnAuthGuard element={<Login />} />}
    />
    <Route
      key="Register"
      path="signup"
      element={<UnAuthGuard element={<Registration />} />}
    />
  </Route>,
]

export default UnAuthRoutes;
