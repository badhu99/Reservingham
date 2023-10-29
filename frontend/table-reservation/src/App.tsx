import React, { useEffect } from "react";
import "./App.scss";
import {
  BrowserRouter,
  Navigate,
  PathRouteProps,
  Route,
  Routes,
} from "react-router-dom";
import Login from "./pages/authentication/login";
import BackboneAuthentication from "./pages/authentication/backbone-auth";
import Registration from "./pages/authentication/registration";
import BackboneContent from "./pages/backbone-content";
import Dashboard from "./pages/dashboard/dashboard";
import Editor from "./pages/editor/editor";
import Reservations from "./pages/reservations/reservations";
import NotFound from "./pages/common/not-found";
import AllowedAccess from "./utils/helpers/allowed-access";
import { Roles } from "./utils/enums/roles";
import useToken from "./hooks/useToken";


function App() {

  const {tokenData} = useToken()
  useEffect(() => {
    
  }, [tokenData])

  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="auth" element={<BackboneAuthentication />}>
            <Route path="" element={<Navigate to="/auth/signin" />} />
            <Route path="signin" element={<Login />} />
            <Route path="signup" element={<Registration />} />
          </Route>
          <Route path="/" element={AllowedAccess( [Roles.User, Roles.Admin, Roles.Editor], <BackboneContent /> )}>
            <Route path="/" element={<Navigate to="/dashboard" />} />
            <Route path="reservations" element={AllowedAccess( [Roles.User], <Reservations />)} />
            <Route path="dashboard" element={AllowedAccess( [Roles.User], <Dashboard />)} />
            <Route path="editor" element={AllowedAccess( [Roles.Editor], <Editor />)} />
            <Route path="admin" element={AllowedAccess( [Roles.Admin], <Editor />)} />
          </Route>
          <Route path="not-found" element={<NotFound />} />
          <Route path="*" element={<Navigate to="/not-found" />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
