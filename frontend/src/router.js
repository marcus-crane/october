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

const tailwindButton = (extraStyles) => (
  <div>
    <button type="button" className={extraStyles + " min-h-full align-middle block rounded-md p-1.5 focus:outline-none focus:ring-2 focus:ring-offset-2"}>
      <span className="sr-only">Dismiss</span>
      <svg className="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
        <path fillRule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
      </svg>
    </button>
  </div>
)

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
      closeButton={({ type }) => {
        const extraStyles = contextClass[type || "default"]
        return tailwindButton(extraStyles)
      }}
      hideProgressBar
      autoClose={3000}
      pauseOnHover
      transition={Slide}
      bodyClassName={() => "text-sm w-screen flex p-2"}
      toastClassName={({ type }) => contextClass[type || "default"] +
        " rounded-md p-4 flex mb-2"
      }
    />
  </HashRouter>
)

export default Router
