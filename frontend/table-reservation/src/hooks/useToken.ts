import { jwtDecode } from "jwt-decode";
import { useState } from "react";
import { UserData } from "../classes/user_data";

const constToken = "userToken";

export default function useToken() {

  const saveToken = (userToken: string) => {
    sessionStorage.setItem(constToken, userToken);
    const userData = getTokenData();
    setTokenData(prevUserData => ({...prevUserData, ...userData}))
    setTokenString(userToken)

    
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

  const getTokenString = ():string => {
    const token = sessionStorage.getItem(constToken);
    if (token === null) return "";
    return token;
  }

  const [tokenData, setTokenData] = useState(getTokenData());
  const [tokenString, setTokenString] = useState(getTokenString())

  return {
    setToken: saveToken,
    clearToken,
    tokenData,
    tokenString,
    getTokenString
  };
}
