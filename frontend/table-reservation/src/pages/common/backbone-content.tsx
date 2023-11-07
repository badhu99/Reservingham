import "./backbone-content.scss";

import { Outlet, Link, useNavigate } from "react-router-dom";
import useToken from "../../hooks/useToken";
import { Roles } from "../../utils/enums/roles";
import { AllowedAccessBool } from "../../utils/helpers/guards/check-access";

export default function BackboneContent() {
  let navigate = useNavigate();

  const { clearToken, tokenData } = useToken();

  const deleteToken = () => {
    clearToken();
    navigate("./../auth");
  };

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
            {AllowedAccessBool([Roles.User], tokenData) && (
              <li>
                <Link to="/reservations">Reservations</Link>
              </li>
            )}
            {AllowedAccessBool([Roles.Editor], tokenData) && (
              <li>
                <Link to="/editor">Editor</Link>
              </li>
            )}
            {AllowedAccessBool([Roles.Manager], tokenData) && (
              <li>
                <Link to="/manager">Manager</Link>
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
