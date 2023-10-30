import { jwtDecode } from "jwt-decode";
import { useState } from "react";
import { UserData } from "../classes/user_data";

const constToken = "userToken";

export default function useToken() {

  const saveToken = (userToken: string) => {
    sessionStorage.setItem(constToken, userToken);
    const userData = getTokenData();
    setTokenData(prevUserData => ({...prevUserData, ...userData}))
    
  };

  const clearToken = () => {
    sessionStorage.removeItem(constToken);
    saveToken("");
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
