import React from "react";
import { createRoot } from "react-dom/client";

import { HashRouter, Route, Routes } from "react-router-dom";

import "./style.css";
import DeviceSelector from "./pages/DeviceSelector"
import Settings from "./pages/Settings"

const routes = (
  <React.StrictMode>
    <HashRouter>
      <Routes>
        <Route path="/" element={<DeviceSelector />} />
        <Route path="/settings" element={<Settings />} />
      </Routes>
    </HashRouter>
    
  </React.StrictMode>
);
const container = document.getElementById("root");
const root = createRoot(container!);
root.render(routes);
