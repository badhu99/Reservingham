import { useEffect, useState } from "react";
import "./login.scss";

import { Link, useNavigate } from "react-router-dom";
import axios, { AxiosError } from "axios";
import { loginUser } from "../../apis/authentication-api";
import {
  LoginRequestModel,
  LoginResponseModel,
} from "../../classes/login_model";
import { type } from "os";
import useToken from "../../hooks/useToken";

export default function Login() {
  let navigate = useNavigate();

  const [errorMessage, setErrorMessage] = useState("")
  const [user, setUser] = useState(new LoginRequestModel());
  const {tokenData, setToken} = useToken();

  useEffect(() => {
    if (tokenData.HasValue){
        navigate("./../../dashboard")
    }
  }, [tokenData])

  const handleUserPassword = (event: React.FormEvent<HTMLInputElement>) => {
    console.log(event)
    setUser({ ...user, Password: event.currentTarget.value });
  };

  const handleUserUsername = (event: React.FormEvent<HTMLInputElement>) => {
    setUser({ ...user, Username: event.currentTarget.value });
  };


  const login = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const response = await loginUser(user);
    if (response instanceof AxiosError) {
      const error = response.response?.data as string;
      setErrorMessage(error);
      setUser({ ...user, Username: "", Password: "" });
      return;
    }

    setToken(response.AccessToken);

    navigate("./../../dashboard")
  };

  return (
    <div className="content">
      <div className="div-other"></div>
      <div className="login-form">
        <form onSubmit={login}>
          <div className="header">
            <h1>Login</h1>
          </div>
          <div className="form-content">
            <label>Username *</label>
            <input
              type="text"
              value={user.Username}
              onChange={handleUserUsername}
            />
            <label>Password *</label>
            <input
              type="password"
              value={user.Password}
              onChange={handleUserPassword}
            />
            <div className="div-rememberme">
              <p>Forgot your password?</p>
            </div>
          </div>
          <div className="footer">
            <p className="error-message"> {errorMessage}</p>
            <button type="submit">
              Login
            </button>
            <p>
              Don't have an account? <Link to="../signup">Sign up</Link>
            </p>
          </div>
        </form>
      </div>
    </div>
  );
}
