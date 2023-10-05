import "./backbone-content.scss"

import React from "react";
import { Outlet, Link } from "react-router-dom";

export default function BackboneContent() {
    return (
        <div className="container">
            <nav>
                <ul className="menuItems">
                    <li><Link to="/dashboard">Dashboard</Link></li>
                    <li><Link to="/editor">Editor</Link></li>
                    <li><Link to="/reservations">Reservations</Link></li>
                </ul>
            </nav>
            <div className="main-content">
                <Outlet />
            </div>
        </div>
    )
}