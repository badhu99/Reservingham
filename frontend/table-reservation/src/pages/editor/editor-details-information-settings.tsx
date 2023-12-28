import React, { useState } from 'react';

interface IEditorDetailsInformationProps {
  Reservations: boolean | undefined;
  UpdateReservations: () => void;
}

const EditorDetailsInformation: React.FC<IEditorDetailsInformationProps> = ({Reservations, UpdateReservations}) => {
  return (
    <div className="div-editor-sidebar-details">
      <h1>Details</h1>
      <label>Name</label>
      <input className="input-name" type="text" value={"test"} />
      <input
        className="input-coordinate input-number"
        type="number"
        value={12}
      />
      <input
        className="input-coordinate input-number"
        type="number"
        value={12}
      />
      <label>Width</label>
      <input className="input-width" type="number" value={12} />
      <label>Height</label>
      <input className="input-height input-number" type="number" value={12} />
      <label>Color</label>
      <input className="input-color" type="text" value={12} />
      <input type="checkbox" checked={Reservations} onChange={UpdateReservations} />
    </div>
  );
};

export default EditorDetailsInformation;