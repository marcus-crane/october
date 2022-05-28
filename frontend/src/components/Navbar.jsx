import React from "react";
import { NavLink, useLocation } from "react-router-dom"
import { BookmarkIcon, ChipIcon, CogIcon } from '@heroicons/react/solid'

export default function Navbar() {
  const location = useLocation()
  return (
    <header className="flex p-3">
      <div className="w-full text-left">
        {location.pathname === "/overview" && <NavLink to="/"><ChipIcon className="h-5 w-5 inline-block" /> Pick a different device</NavLink>}
        {location.pathname === "/settings" && <NavLink to="/overview"><BookmarkIcon className="h-5 w-5 inline-block" /> Return to device overview</NavLink>}
      </div>
      <div className="w-full text-right">
        {location.pathname === "/overview" && <NavLink to="/settings">Settings <CogIcon className="h-5 w-5 inline-block" /></NavLink>}
      </div>
    </header>
  )
}
