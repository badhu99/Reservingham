import "./backbone-content.scss";

import { useEffect } from "react";
import { Outlet, Link, useNavigate } from "react-router-dom";
import useToken from "../hooks/useToken";
import { Roles } from "../utils/enums/roles";
import { AllowedAccessBool } from "../utils/helpers/allowed-access";

export default function BackboneContent() {
  let navigate = useNavigate();

  const { clearToken, tokenData } = useToken();

  const deleteToken = () => {
    clearToken();
  };

  useEffect(() => {
    AllowedAccessBool([Roles.Editor], tokenData);
    if (!tokenData.HasValue) {
      navigate("./../auth");
    }
  }, [tokenData]);

  return (
    <div className="container">
      <nav>
        <div className="left-side">
          <ul className="menuItems">
            {AllowedAccessBool([Roles.User], tokenData) && (
              <li>
                <Link to="/dashboard">Dashboard</Link>
              </li>
            )}
            {AllowedAccessBool([Roles.Editor], tokenData) && (
              <li>
                <Link to="/editor">Editor</Link>
              </li>
            )}
            {AllowedAccessBool([Roles.User], tokenData) && (
              <li>
                <Link to="/reservations">Reservations</Link>
              </li>
            )}
            {AllowedAccessBool([Roles.Admin], tokenData) && (
              <li>
                <Link to="/admin">Admin</Link>
              </li>
            )}
          </ul>
        </div>
        <div className="right-side">
          <text>{tokenData.Username}</text>
          <a onClick={deleteToken}>Logout</a>
        </div>
      </nav>
      <div className="main-content">
        <Outlet />
      </div>
    </div>
  );
}