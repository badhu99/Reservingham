import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { GetUser, User } from "../../apis/user-api";
import { AxiosError } from "axios";
import { ReactComponent as XIcon } from "./../../assets/x2-icon.svg";
import { ReactComponent as AddIcon } from "./../../assets/add2-icon.svg";
import { GetRoles, Role } from "../../apis/roles-api";

interface LoginInfo {
  Username: string;
  Email: string;
}

interface Password {
  Password: string;
  PasswordRepeat: string;
}

interface PersonalInfo {
  Firstname: string;
  Lastname: string;
}

export function ManagerSinglePage() {
  const param = useParams();
  const [userDetailsData, setUserDetailsData] = useState<User>();
  const [roles, setRoles] = useState<Role[]>();
  const [loginInfo, setLoginInfo] = useState<LoginInfo>();
  const [password, setPassword] = useState<Password>({
    Password: "",
    PasswordRepeat: "",
  });
  const [personalInfo, setPersonalInfo] = useState<PersonalInfo>();

  const loginInfoChanged =
    loginInfo?.Username !== userDetailsData?.Username ||
    loginInfo?.Email !== userDetailsData?.Email;
  const passwordChanged = () => {
    let changed = password?.Password !== "" || password.PasswordRepeat !== "";
    if (password.Password !== password.PasswordRepeat) {
      changed = false;
    }

    return changed;
  };
  const personalInfoChanged =
    personalInfo?.Firstname !== userDetailsData?.Firstname ||
    personalInfo?.Lastname !== userDetailsData?.Lastname;

  useEffect(() => {
    if (param.id) {
      getUserDetailInfo(param.id);
    }
  }, []);

  const getUserDetailInfo = async (id: string) => {
    const data = await GetUser(id);
    if (!(data instanceof AxiosError)) {
      setUserDetailsData(data);
      const allRoles = await GetRoles();
      if (!(allRoles instanceof AxiosError)) {
        const available = allRoles.Items.filter(
          (item) => !data.Roles.some((ur) => ur.Id === item.Id)
        );
        setRoles(available);
      }

      setLoginInfo({ Username: data.Username, Email: data.Email });
      setPersonalInfo({ Firstname: data.Firstname, Lastname: data.Lastname });
    }
  };

  const handleAddingRoles = (role: Role) => {
    if (userDetailsData !== undefined) {
      let newRoles = userDetailsData?.Roles!;
      newRoles.push(role);
      setUserDetailsData({ ...userDetailsData, Roles: newRoles });
    }
    const newRoles = roles?.filter((item) => item.Id !== role.Id);
    setRoles(newRoles);
  };

  const handleRemovingRoles = (role: Role) => {
    if (userDetailsData !== undefined) {
      let newRoles = userDetailsData?.Roles!;
      newRoles = newRoles.filter((item) => item.Id !== role.Id);
      setUserDetailsData({ ...userDetailsData, Roles: newRoles });
    }
    roles?.push(role);
    setRoles(roles);
  };

  const handleChangeUsername = (e: React.ChangeEvent<HTMLInputElement>) => {
    setLoginInfo({ Email: loginInfo?.Email!, Username: e.target.value });
  };

  const handleChangeEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setLoginInfo({ Email: e.target.value, Username: loginInfo?.Username! });
  };

  const handleChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword({
      Password: e.target.value,
      PasswordRepeat: password?.PasswordRepeat ?? "",
    });
  };

  const handleChangeRepeatPassword = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    setPassword({
      Password: password?.Password ?? "",
      PasswordRepeat: e.target.value,
    });
  };

  const handleChangeFirstname = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPersonalInfo({
      Firstname: e.target.value,
      Lastname: personalInfo?.Lastname ?? "",
    });
  };

  const handleChangeLastname = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPersonalInfo({
      Firstname: personalInfo?.Firstname ?? "",
      Lastname: e.target.value,
    });
  };

  return (
    <div className="detail-page-grid">
      <div className="profile-pic">
        <img
          className="img-responsive user-pic"
          src="https://i.imgur.com/o8ouMW0.jpg"
          alt=""
        />
        <a href="#">
          <p className="change-image">CHANGE IMAGE</p>
        </a>
        <button type="button" className="btn btn-default change-img-btn">
          SAVE IMAGE
        </button>
      </div>

      <div className="login-info">
        <h3 className="heading">Login info</h3>
        <form>
          <label>Username:</label>
          <br />
          <input
            type="text"
            className="input-git"
            value={loginInfo?.Username}
            onChange={handleChangeUsername}
          />
          <br />
          <label>Email:</label>
          <br />
          <input
            type="text"
            className="input-git"
            value={loginInfo?.Email}
            onChange={handleChangeEmail}
          />
          <br />
          {loginInfoChanged && <button className="btn-git">Save</button>}
        </form>
      </div>

      <div className="password-info">
        <h3 className="heading">Password</h3>
        <form>
          <label>New Password:</label>
          <br />
          <input
            type="password"
            className="input-git"
            value={password?.Password}
            onChange={handleChangePassword}
          />
          <br />
          <label>New Password:</label>
          <br />
          <input
            type="password"
            className="input-git"
            value={password?.PasswordRepeat}
            onChange={handleChangeRepeatPassword}
          />
          <br />
          {passwordChanged() && <button className="btn-git">Save</button>}
        </form>
      </div>

      <div className="personal-info">
        <h3 className="heading">Personal Information</h3>
        <form>
          <label>First name:</label>
          <br />
          <input
            className="input-git"
            type="text"
            value={personalInfo?.Firstname}
            onChange={handleChangeFirstname}
          />
          <br />
          <label>Last name:</label>
          <br />
          <input
            className="input-git"
            type="text"
            value={personalInfo?.Lastname}
            onChange={handleChangeLastname}
          />
          <br />
          {personalInfoChanged && <button className="btn-git">Save</button>}
        </form>
      </div>

      <div className="permissions-info">
        <h3 className="heading">Permissions</h3>
        <div className="permission-container">
          <h5>Current</h5>
          <br />
          {userDetailsData && (
            <UserRoles
              data={userDetailsData.Roles}
              iconType={"Remove"}
              funcClick={handleRemovingRoles}
            />
          )}
          <h5>Available</h5>
          <br />
          {roles && (
            <UserRoles
              data={roles}
              iconType={"Add"}
              funcClick={handleAddingRoles}
            />
          )}
        </div>
      </div>
    </div>
  );
}

const UserRoles = ({
  data,
  iconType,
  funcClick,
}: {
  data: Role[];
  iconType: "Add" | "Remove";
  funcClick: (role: Role) => void;
}) => {
  return (
    <>
      {data.map((item) => {
        return (
          <div className="user-roles-container" key={item.Id}>
            <div className="sticker">
              <p>{item.Name}</p>
              {iconType === "Remove" ? (
                <XIcon onClick={(e) => funcClick(item)} />
              ) : (
                <AddIcon onClick={(e) => funcClick(item)} />
              )}
            </div>
          </div>
        );
      })}
    </>
  );
};
