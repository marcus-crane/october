import React from 'react'
import { createRoot } from 'react-dom/client'
import { HashRouter, Route, Routes } from "react-router-dom"
import { Toaster } from "react-hot-toast"

import Home from './pages/Home'

const routes = (
  <React.StrictMode>
    <HashRouter>
      <Routes>
        <Route path="/" element={<Home />} />
      </Routes>
    </HashRouter>
    <Toaster />
  </React.StrictMode>
)

const container = document.getElementById('root')
const root = createRoot(container!)
root.render(routes)