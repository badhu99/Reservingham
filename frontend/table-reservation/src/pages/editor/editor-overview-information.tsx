import React, { useRef } from "react";
import { CanvasElement } from "../../interfaces/shapes";

interface IEditorOverviewInformationProps {
  CanvasElements: CanvasElement[];
  UpdateElements: (updatedElements: CanvasElement[]) => void;
}

const EditorOverviewInformation: React.FC<IEditorOverviewInformationProps> = ({
  CanvasElements,
  UpdateElements,
}) => {
  //   const [list, setList] = useState(CanvasElements);
  const draggedItem = useRef<CanvasElement | null>(null);
  const draggedIdx = useRef<number | null>(null);

  const onDragStart = (e: React.DragEvent<HTMLLIElement>, index: number) => {
    draggedItem.current = CanvasElements[index];
    draggedIdx.current = index;
    e.dataTransfer.effectAllowed = "move";
    e.dataTransfer.setData("text/html", e.currentTarget.outerHTML);
  };

  const onDragOver = (index: number) => {
    const draggedOverItem = CanvasElements[index];

    // if the item is dragged over itself, ignore
    if (draggedItem.current === draggedOverItem) {
      return;
    }

    // filter out the currently dragged item
    let items = CanvasElements.filter((item) => item !== draggedItem.current);

    // add the dragged item after the dragged over item
    items.splice(index, 0, draggedItem.current as CanvasElement);

    UpdateElements(items);
  };

  const onDragEnd = () => {
    draggedItem.current = null;
    draggedIdx.current = null;
  };

  return (
    <>
      <div className="div-editor-overview-header">
        <h1>Overview</h1>
      </div>
      <ul className="div-editor-overview-container">
        {CanvasElements.map((item, index) => (
          <li
            key={index}
            onDragOver={() => onDragOver(index)}
            onDragStart={(e) => onDragStart(e, index)}
            onDragEnd={onDragEnd}
            draggable
          >
            Name: {item.name}, x: {item.x}, y: {item.y}
          </li>
        ))}
      </ul>
    </>
  );
};

export default EditorOverviewInformation;
