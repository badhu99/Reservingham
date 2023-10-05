import "./login.scss";

import React from "react";
import { Link } from "react-router-dom";

export default function Login() {
    return (
        <div className="content">
            <div className="div-other">
            </div>
            <div className="login-form">

                <form>
                    <div className="header"><h1>Login</h1></div>
                    <div className="form-content">
                        <label>Username</label>
                        <input type="text" />
                        <label>Password</label>
                        <input type="password" />
                        <div className="div-rememberme">
                        <p>Forgot your password?</p>
                        </div>
                    </div>
                    <div className="footer">                        
                        <button>Login</button>
                        <p>Don't have an account? <Link to="../signup">Sign up</Link></p>
                    </div>
                </form>
            </div>

        </div>
    )
}