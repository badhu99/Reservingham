import { Route } from "react-router-dom";
import Editor from "../../pages/editor/editor";
import { Roles } from "../../utils/enums/roles";
import AuthGuard from "../../utils/helpers/guards/auth-guard";

export default [
    <Route
    path="/editor"
    element={<AuthGuard allowedRoles={[Roles.Editor]} element={<Editor />} />}
  />
]