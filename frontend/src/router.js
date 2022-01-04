import React from "react"
import { HashRouter, Route, Routes } from "react-router-dom"
import {Slide, toast, ToastContainer} from "react-toastify";

import DeviceSelector from "./pages/DeviceSelector"
import Overview from "./pages/Overview"
import Settings from "./pages/Settings"

const contextClass = {
  success: "bg-green-50 text-green-700",
  error: "bg-red-50 text-red-700",
  info: "bg-gray-50 text-gray-700",
  warning: "bg-orange-50 text-orange-700",
  default: "bg-indigo-50 text-indigo-700",
  dark: "bg-white-50 text-gray-700",
};

const Router = () => (
  <HashRouter>
    <Routes>
      <Route exact path="/" element={<DeviceSelector />} />
      <Route exact path="/overview" element={<Overview />} />
      <Route exact path="/settings" element={<Settings />} />
    </Routes>
    <ToastContainer
      position={toast.POSITION.BOTTOM_CENTER}
      limit={2}
      hideProgressBar
      pauseOnHover
      transition={Slide}
      bodyClassName={() => "text-sm w-full flex"}
      toastClassName={({ type }) => contextClass[type || "default"] +
        " rounded-md p-4 flex mb-2"
      }
    />
  </HashRouter>
)

export default Router
