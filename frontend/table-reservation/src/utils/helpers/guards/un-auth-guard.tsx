import React, { useEffect, useState } from "react";
import useToken from "../../../hooks/useToken";
import { Navigate } from "react-router-dom";

interface IUnGuard {
  element: JSX.Element;
}

const UnAuthGuard = ({ element }: IUnGuard) => {
  const [allowed, setAllowed] = useState(0);
  const { tokenData } = useToken();
  useEffect(() => {
    check()
  }, [tokenData, element]);

  const check = () => {
    if (!tokenData.HasValue){
        setAllowed(1);
    }
    else setAllowed(2)
  }

  if (allowed == 1) return element
  else if (allowed == 2) return <Navigate to="/dashboard" replace />
  return <></>;
};

export default UnAuthGuard;
