import React, { useEffect } from 'react';
import { Shape } from '../../interfaces/shapes';


interface IEditorDetailsInformationProps {
    shapes: Shape[];
  }
const EditorDetailsInformation: React.FC<IEditorDetailsInformationProps> = ({shapes}) => {

    const displayShapes = () => {
        console.log("shapes", shapes);
    }
    return (
        <div>
            {shapes.map((shape) => (
                <p>Name: {shape.name}, x: {shape.x}, y: {shape.y}</p>
            ))}
            <button onClick={displayShapes}>Update</button>
        </div>
    );
};

export default EditorDetailsInformation;
