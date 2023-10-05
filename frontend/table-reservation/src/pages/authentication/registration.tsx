import "./registration.scss";
import { Link } from "react-router-dom";

import React from "react";

export default function Registration() {
    return (
        <div className="content">
            <div className="registration-form">
                <form>
                    <div className="header"><h1>Sign up</h1></div>
                    <div className="form-content">
                        <label>Username</label>
                        <input type="text" />
                        <label>Email</label>
                        <input type="text" />
                        <label>Password</label>
                        <input type="password" />
                        <label>Repeat password</label>
                        <input type="password" />
                    </div>
                    <div className="footer">
                        <button>Sign up</button>
                        <p>Already have an account? <Link to="../signin">Sign in</Link></p>
                    </div>
                </form>
            </div>
            <div className="div-other">

            </div>
        </div>
    )
}