import { Route } from "react-router-dom";
import {EditorListAll} from "../../pages/editor/editor-list-all";
import { Roles } from "../../utils/enums/roles";
import AuthGuard from "../../utils/helpers/guards/auth-guard";
import EditorDetails from "../../pages/editor/editor-details";


const EditorRoutes = [
  <Route path="editor">
    <Route
      path=""
      element={
        <AuthGuard allowedRoles={[Roles.Editor]} element={<EditorListAll />} />
      }
    />
    <Route
      path=":id"
      element={
        <AuthGuard
          allowedRoles={[Roles.Editor]}
          element={<EditorDetails />}
        />
      }
    />
  </Route>,
];

export default EditorRoutes;