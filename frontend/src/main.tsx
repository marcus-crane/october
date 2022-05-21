import React from "react";
import { createRoot } from "react-dom/client";
import { HashRouter, Route, Routes } from "react-router-dom";
import { Toaster } from 'react-hot-toast';
import "./style.css";
import DeviceSelector from "./pages/DeviceSelector"
import Library from "./pages/Library"
import Settings from "./pages/Settings"

const routes = (
  <React.StrictMode>
    <HashRouter>
      <Routes>
        <Route path="/" element={<DeviceSelector />} />
        <Route path="/library" element={<Library />} />
        <Route path="/settings" element={<Settings />} />
      </Routes>
    </HashRouter>
    <Toaster />
  </React.StrictMode>
);
const container = document.getElementById("root");
const root = createRoot(container!);
root.render(routes);
