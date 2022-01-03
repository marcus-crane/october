import React from "react"
import { HashRouter, Route, Routes } from "react-router-dom"

import DeviceSelector from "./pages/DeviceSelector"

const Router = () => (
  <HashRouter>
    <Routes>
      <Route exact path="/" element={<DeviceSelector />} />
    </Routes>
  </HashRouter>
)

export default Router
