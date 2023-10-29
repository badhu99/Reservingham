import { jwtDecode } from "jwt-decode";
import { useState } from "react";
import { UserData } from "../classes/user_data";

const constToken = "userToken";

export default function useToken() {

  const saveToken = (userToken: string) => {
    sessionStorage.setItem(constToken, userToken);
    // setToken(userToken);
    const userData = getTokenData();
    setTokenData({...userData})
  };

  const clearToken = () => {
    sessionStorage.removeItem(constToken);
    setTokenData({...tokenData, aud: "",exp: 0, iat: 0, iss: "", sub: "", Id: "", Username: "", Roles: [], HasValue: false})
  };

  const getTokenData = () : UserData => {
    const tokenString = sessionStorage.getItem(constToken);

    var tokenData = new UserData();
    if (tokenString) {
      tokenData = jwtDecode<UserData>(tokenString!);

        if (tokenData.Roles.length > 0 && tokenData.Username && tokenData.Username){
            tokenData.HasValue = true;
            return tokenData;
        }
    }

    tokenData.HasValue = false;
    return tokenData;
  };

  const [tokenData, setTokenData] = useState(getTokenData());

  return {
    setToken: saveToken,
    clearToken,
    tokenData,
  };
}
