import React from "react"
import { HashRouter, Route, Routes } from "react-router-dom"
import {Slide, toast, ToastContainer} from "react-toastify";

import DeviceSelector from "./pages/DeviceSelector"
import Overview from "./pages/Overview"
import Settings from "./pages/Settings"

const contextClass = {
  success: "bg-blue-600",
  error: "bg-red-600",
  info: "bg-gray-600",
  warning: "bg-orange-400",
  default: "bg-indigo-600",
  dark: "bg-white-600 font-gray-300",
};

const Router = () => (
  <HashRouter>
    <Routes>
      <Route exact path="/" element={<DeviceSelector />} />
      <Route exact path="/overview" element={<Overview />} />
      <Route exact path="/settings" element={<Settings />} />
    </Routes>
    <ToastContainer
      autoClose={3000}
      position={toast.POSITION.BOTTOM_CENTER}
      limit={2}
      transition={Slide}
    />
  </HashRouter>
)

export default Router
