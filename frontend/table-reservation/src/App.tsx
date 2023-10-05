import React from 'react';
import './App.scss';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';
import Login from './pages/authentication/login';
import BackboneAuthentication from './pages/authentication/backbone-auth';
import Registration from './pages/authentication/registration';
import BackboneContent from './pages/backbone-content';
import Dashboard from './pages/dashboard/dashboard';
import Editor from './pages/editor/editor';
import Reservations from './pages/reservations/reservations';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="auth" element={<BackboneAuthentication />} >
            <Route path='' element={<Navigate to="/auth/signin" />} />
            <Route path="signin" element={<Login />} />
            <Route path="signup" element={<Registration />} />
          </Route>
          <Route path="/" element={<BackboneContent />} >
            <Route path='/' element={<Navigate to="/dashboard" />} />
            <Route path="dashboard" element={<Dashboard />} />
            <Route path="editor" element={<Editor />} />
            <Route path="reservations" element={<Reservations />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
