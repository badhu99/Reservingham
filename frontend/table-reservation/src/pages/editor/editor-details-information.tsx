import React, { useEffect } from 'react';
import { CanvasElement } from '../../interfaces/shapes';


interface IEditorDetailsInformationProps {
    canvasElements: CanvasElement[];
  }
const EditorDetailsInformation: React.FC<IEditorDetailsInformationProps> = ({canvasElements}) => {

    return (
        <div>
            {canvasElements.map((shape) => (
                <p>Name: {shape.name}, x: {shape.x}, y: {shape.y}</p>
            ))}
        </div>
    );
};

export default EditorDetailsInformation;
