import { useState } from "react";
import CanvasEditor from "../../components/canvas-editor/canvas-editor";
import { CanvasElement } from "../../interfaces/shapes";
import EditorSidebar from "./editor-sidebar";

export default function EditorDetails() {
  const [elements, setElements] = useState<CanvasElement[]>([
    {
      id: "640ab176-ce44-4e1f-afae-5121578610c4",
      name: "rect1",
      x: 10,
      y: 100,
      width: 30,
      height: 30,
      fill: "#444444",
      isDragging: false,
      isReservable: false,
    },
    {
      id: "bd332da1-6070-4604-8a9b-f70d086a187b",
      name: "rect2",
      x: 80,
      y: 100,
      width: 30,
      height: 30,
      fill: "#ff550d",
      isDragging: false,
      isReservable: false,
    },
    {
      id: "9ef59004-5ea2-43f9-90a0-d0065a92fabd",
      name: "circle2",
      x: 200,
      y: 100,
      r: 10,
      fill: "#0c64e8",
      isDragging: false,
      isReservable: false,
    },
    {
      id: "23732654-6b10-40af-82e9-e3fbf43dfaa1",
      name: "circle1",
      x: 150,
      y: 100,
      r: 10,
      fill: "#800080",
      isDragging: false,
      isReservable: false,
    },
    {
      id: "039a1ec4-4cfc-47cb-894d-578b5ad3eea4",
      name: "rect3",
      x: 180,
      y: 250,
      width: 30,
      height: 30,
      fill: "#68038c",
      isDragging: false,
      isReservable: false,
    },
  ]);

  const [element, setElement] = useState<CanvasElement>();

  const updateElements = (updatedShapes: CanvasElement[]) => {
    setElements([...updatedShapes]);
  };

  const setClickedElement = (element: CanvasElement | undefined) => {
    setElement(element);
  }

  const updateSelectedElement = (updatedElement: CanvasElement) => {
    if (element) {
      const updatedElements = elements.map((el) =>
        el.id === element.id ? { ...el, ...updatedElement } : el
      );
      setElements(updatedElements);
    }
  };

  return (
    <>
      <h1>Editor details</h1>
      <div className="div-container-editor-details">
        <div className="div-editor">
          <CanvasEditor Elements={elements} UpdateElements={updateElements} SetClickedElement={setClickedElement}/>
        </div>
        <div className="div-container-editor-sidebar">
          <EditorSidebar Elements={elements} UpdateElements={updateElements} ElementSelected={element} UpdateElement={updateSelectedElement}/>
        </div>
      </div>
    </>
  );
}
