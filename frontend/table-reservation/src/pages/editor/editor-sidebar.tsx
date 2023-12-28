import React, { useState } from "react";
import { CanvasElement } from "../../interfaces/shapes";
import EditorDetailsInformation from "./editor-details-information";

interface IEditorSidebarProps {
    Elements: CanvasElement[]
}

const EditorSidebar: React.FC<IEditorSidebarProps> = ({Elements}) => {
  const [activeTab, setActiveTab] = useState(1);

  const handleTabClick = (tabNumber: number) => {
    setActiveTab(tabNumber);
  };

  return (
    <>
      <div className="div-editor-content">
        {activeTab === 1 && <EditorDetailsInformation canvasElements={Elements} />}
        {activeTab === 2 && <h3>Tab 2 Content</h3>}
        {activeTab === 3 && <h3>Tab 3 Content</h3>}
      </div>
      <div className="div-editor-tabs">
        <button onClick={() => handleTabClick(1)}>Tab 1</button>
        <button onClick={() => handleTabClick(2)}>Tab 2</button>
        <button onClick={() => handleTabClick(3)}>Tab 3</button>
      </div>
    </>
  );
};

export default EditorSidebar;
