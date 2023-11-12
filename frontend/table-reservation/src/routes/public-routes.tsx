import { Route, Navigate } from "react-router-dom";
import NotFound from "../pages/common/not-found";

const publicRoutes =  [
        <Route path="not-found" element={<NotFound />} />,
        <Route path="*" element={<Navigate to="/not-found" />} />,
]

export default publicRoutes;