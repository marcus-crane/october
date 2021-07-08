import React from "react"
import { HashRouter, Route, Switch } from "react-router-dom"

import DeviceSelector from "./pages/DeviceSelector"

const Router = (props) => (
  <HashRouter>
    <Route exact path="/" component={DeviceSelector} />
  </HashRouter>
)

export default Router
