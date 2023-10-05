import React from "react";
import "./backbone-auth.scss";
import { Outlet } from "react-router-dom";

export default function BackboneAuthentication() {
    return (
        <div className="container">
            <div className="container-holder">
                <Outlet/>
            </div>
        </div>
    )
}