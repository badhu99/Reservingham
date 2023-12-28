import React, { useEffect, useState } from "react";
import { CanvasElement } from "../../interfaces/shapes";
import EditorDetailsInformation from "./editor-details-information-settings";
import EditorOverviewInformation from "./editor-overview-information";
import EditorDetailsReservationsSettings from "./editor-details-reservations";

interface IEditorSidebarProps {
        Elements: CanvasElement[];
        UpdateElements: (updatedElements: CanvasElement[]) => void;
        ElementSelected: CanvasElement | undefined;
        UpdateElement: (updatedElement: CanvasElement) => void;
}

const EditorSidebar: React.FC<IEditorSidebarProps> = ({Elements, UpdateElements, ElementSelected, UpdateElement}) => {
    const [activeTab, setActiveTab] = useState(1);

    const handleTabClick = (tabNumber: number) => {
        setActiveTab(tabNumber);
    };

    useEffect(() => {
        if (ElementSelected !== undefined) {
            setActiveTab(2);
        }
        if (ElementSelected === undefined) {
            setActiveTab(1);
        }
    }, [ElementSelected])

    const showTabDetails = () : boolean => {
        return ElementSelected !== undefined;
    }

    const showTabReservationsSettings = () : boolean => {
        return ElementSelected?.isReservable === true && ElementSelected !== undefined;
    }

    const updateReservations = () => {

        ElementSelected!.isReservable = !ElementSelected?.isReservable;
        UpdateElement(ElementSelected!);

        const updatedElements = Elements.map(element => {
            if (element === ElementSelected) {
                return ElementSelected;
            }
            return element;
        });        
        UpdateElements(updatedElements);
    }

    return (
        <>
            <div className="div-editor-sidebar-content">
                {activeTab === 1 && <EditorOverviewInformation CanvasElements={Elements} UpdateElements={UpdateElements} />}
                {activeTab === 2 && <EditorDetailsInformation Reservations={ElementSelected?.isReservable} UpdateReservations={updateReservations}/>}
                {activeTab === 3 && <EditorDetailsReservationsSettings />}
            </div>
            <div className="div-editor-tabs">
                <button className="btn-tabs" onClick={() => handleTabClick(1)}>Overview</button>
                {showTabDetails() && <button className="btn-tabs" onClick={() => handleTabClick(2)}>Details</button>}
                {showTabReservationsSettings() && <button className="btn-tabs" onClick={() => handleTabClick(3)}>Reservations</button>}
            </div>
        </>
    );
};

export default EditorSidebar;
