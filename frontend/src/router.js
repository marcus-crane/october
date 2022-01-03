import React from "react"
import { HashRouter, Route, Routes } from "react-router-dom"

import DeviceSelector from "./pages/DeviceSelector"
import Overview from "./pages/Overview"
import Settings from "./pages/Settings"

const Router = () => (
  <HashRouter>
    <Routes>
      <Route exact path="/" element={<DeviceSelector />} />
      <Route exact path="/overview" element={<Overview />} />
      <Route exact path="/settings" element={<Settings />} />
    </Routes>
  </HashRouter>
)

export default Router
