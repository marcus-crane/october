import React from "react";
import { NavLink, useLocation } from "react-router-dom"
import { BookmarkIcon, CpuChipIcon, CogIcon } from '@heroicons/react/20/solid'

export default function Navbar() {
  const location = useLocation()
  return (
    <header className="flex p-3">
      <div className="w-full text-left text-gray-900 dark:text-gray-300">
        {location.pathname === "/overview" && <NavLink to="/"><CpuChipIcon className="h-5 w-5 inline-block" /> Pick a different device</NavLink>}
        {location.pathname === "/settings" && <NavLink to="/overview"><BookmarkIcon className="h-5 w-5 inline-block" /> Return to device overview</NavLink>}
      </div>
      <div className="w-full text-right text-gray-900 dark:text-gray-300">
        {location.pathname === "/overview" && <NavLink to="/settings">Settings <CogIcon className="h-5 w-5 inline-block" /></NavLink>}
      </div>
    </header>
  )
}
