import { BrowserRouter, Routes } from "react-router-dom"
import AuthRoutes from "./auth-routes"
import UnAuthRoutes from "./un-auth-routes"
import PublicRoutes from "./public-routes"

const AppRouting = () => {
    return(
        <BrowserRouter>
        <Routes>
          {AuthRoutes}
          {UnAuthRoutes}
          {PublicRoutes}
        </Routes>
      </BrowserRouter>
    )
}

export default AppRouting;