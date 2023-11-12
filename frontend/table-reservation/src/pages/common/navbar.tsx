import { Outlet, Link, useNavigate } from "react-router-dom";
import useToken from "../../hooks/useToken";
import { Roles } from "../../utils/enums/roles";
import { AllowedAccessBool } from "../../utils/helpers/guards/check-access";
import { ReactComponent as SwitchIcon} from './../../assets/switch-icon.svg'
import { ReactComponent as LogoutIcon} from './../../assets/logout-icon.svg'
import { useState } from "react";
import { Modal } from "./modal";

export default function Navbar() {
  const [openModal, setOpenModal] = useState(false);
  let navigate = useNavigate();

  const { clearToken, tokenData } = useToken();

  const deleteToken = () => {
    clearToken();
    navigate("./../auth");
  };

  const toggleOpenModal = () => {
    setOpenModal(prev => !prev)
  }

  return (
    <div className="container">
      <Modal handleClose={ toggleOpenModal } show={openModal}>
        <p>Here comes popup</p>
      </Modal>
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
                <Link to="/users">Users</Link>
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
          <SwitchIcon onClick={toggleOpenModal}/>
          <text>{tokenData.Username}</text>
          <LogoutIcon onClick={deleteToken}/>
        </div>
      </nav>
      <div className="main-content">
        <Outlet />
      </div>
    </div>
  );
}
