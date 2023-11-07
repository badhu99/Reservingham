import "./App.scss";
import {
  BrowserRouter,
  Navigate,
  Route,
  Routes,
} from "react-router-dom";
import Login from "./pages/authentication/login";
import BackboneAuthentication from "./pages/authentication/backbone-auth";
import Registration from "./pages/authentication/registration";
import BackboneContent from "./pages/common/backbone-content";
import Dashboard from "./pages/dashboard/dashboard";
import Editor from "./pages/editor/editor";
import NotFound from "./pages/common/not-found";
import AuthGuard from "./utils/helpers/guards/auth-guard";
import { Roles } from "./utils/enums/roles";
import Reservations from "./pages/reservations/reservations";
import UnAuthGuard from "./utils/helpers/guards/un-auth-guard";
import Manager from "./pages/manager/manager";
import Admin from "./pages/admin/admin";

const AuthRoutes = [
  <Route
    path=""
    element={
      <AuthGuard
        allowedRoles={[Roles.Admin, Roles.Editor, Roles.User]}
        element={<BackboneContent />}
      />
    }
  >
    <Route path="" element={<Navigate to="/dashboard" />} />
    <Route
      path="/dashboard"
      element={
        <AuthGuard allowedRoles={[Roles.User]} element={<Dashboard />} />
      }
    />
    <Route
      path="/reservations"
      element={
        <AuthGuard allowedRoles={[Roles.User]} element={<Reservations />} />
      }
    />
    <Route
      path="/editor"
      element={
        <AuthGuard allowedRoles={[Roles.Editor]} element={<Editor />} />
      }
    />
        <Route
      path="/manager"
      element={
        <AuthGuard allowedRoles={[Roles.Manager]} element={<Manager />} />
      }
    />
    <Route
      path="/admin"
      element={
        <AuthGuard allowedRoles={[Roles.Admin]} element={<Admin />} />
      }
    />
  </Route>,
];

const UnAuthRoutes = [
  <Route path="auth" element={<UnAuthGuard element={<BackboneAuthentication />} />}>
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
];

const PublicRoutes = [
  <Route path="not-found" element={<NotFound />} />,
  <Route path="*" element={<Navigate to="/not-found" />} />
]

const App = () =>  {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          {AuthRoutes}
          {UnAuthRoutes}
          {PublicRoutes}
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
