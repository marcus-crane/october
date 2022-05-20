import React from "react";
import { NavLink, useLocation } from "react-router-dom";
import { SearchIcon } from "@heroicons/react/solid";
import { BellIcon, CogIcon } from "@heroicons/react/outline";
import { BookmarkIcon, ChipIcon } from "@heroicons/react/solid";
import logo from '../logo.png'

export default function Navbar() {
  const location = useLocation();
  return (
    <div className="flex justify-between items-center px-4 py-6 sm:px-6 space-x-10">
      <div>
        <a href="#" className="flex">
          <span className="sr-only">Workflow</span>
          <img
            className="h-8 w-auto sm:h-10 logo-animation"
            src={logo}
            alt=""
          />
        </a>
      </div>
      <div className="-mr-2 -my-2">
        <div className="flex items-center ml-12">
          <a
            href="#"
            className="text-base font-medium text-gray-500 hover:text-gray-900"
          >
            Pick another device
          </a>
          <a
            href="#"
            className="ml-8 inline-flex items-center justify-center px-4 py-2 border border-transparent rounded-md shadow-sm text-base font-medium text-white bg-indigo-600 hover:bg-indigo-700"
          >
            Settings
          </a>
        </div>
      </div>
    </div>
  );
}
