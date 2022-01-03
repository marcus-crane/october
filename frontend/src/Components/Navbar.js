import React from "react";
import { NavLink, useLocation } from "react-router-dom"

export default function Navbar() {
  const location = useLocation()
  return (
    <header className="flex p-3">
      <div className="w-full text-left">
        {location.pathname === "/overview" && <NavLink to="/">← Pick a different device</NavLink>}
        {location.pathname === "/settings" && <NavLink to="/overview">← Return to device overview</NavLink>}
      </div>
      <div className="w-full text-right">
        {location.pathname === "/overview" && <NavLink to="/settings">Settings →</NavLink>}
      </div>
    </header>
  )
}
