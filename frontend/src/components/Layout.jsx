import { useState, useEffect } from "react";
import { NavLink, useLocation } from "react-router-dom"
import { GetSelectedKobo } from "../../wailsjs/go/backend/Backend";
import { Disclosure } from "@headlessui/react";
import {
  BookOpenIcon,
  RefreshIcon,
  CogIcon,
  MapIcon,
  StatusOnlineIcon,
} from "@heroicons/react/outline";
import logo from '../logo.png'

const sidebarNavigation = [
  { name: "Load", href: "/", icon: MapIcon, koboSelectionRequired: false },
  { name: "Sync", href: "/overview", icon: RefreshIcon, koboSelectionRequired: true },
  { name: "Library", href: "/library", icon: BookOpenIcon, koboSelectionRequired: true },
  { name: "Queue", href: "#", icon: StatusOnlineIcon, koboSelectionRequired: true },
  { name: "Settings", href: "/settings", icon: CogIcon, koboSelectionRequired: false }
];

function classNames(...classes) {
  return classes.filter(Boolean).join(" ");
}

function determineExtraClasses(pathname, item, selectedKobo) {
  let classes = []
  if (pathname === item.href) {
    classes.push("bg-gray-900 text-white")
  } else {
    classes.push("text-gray-400 hover:bg-gray-700")
  }
  if (item.koboSelectionRequired && !selectedKobo.name) {
    classes.push("cursor-not-allowed")
  }
  return classes.join(" ")
}

export default function Layout(props) {
  const [selectedKobo, setSelectedKobo] = useState({})
  console.log('Layout render')
  useEffect(() => {
    GetSelectedKobo()
      .then(kobo => setSelectedKobo(kobo))
      .catch(err => console.log('No Kobo found'))
  }, [selectedKobo.name])
  const location = useLocation()
  return (
    <>
      <div className="h-full flex flex-col">
        {/* Top nav*/}
        <header className="flex-shrink-0 relative h-16 flex items-center">
          {/* Logo area */}
          <div className="block inset-y-0 left-0 flex-shrink-0">
            <a
              href="#"
              className="flex items-center justify-center h-16 w-20 bg-gray-800 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-600"
            >
              <img
                className="h-10 w-auto"
                src={logo}
                alt="Workflow"
              />
            </a>
          </div>
        </header>

        {/* Bottom section */}
        <div className="min-h-0 flex-1 flex overflow-hidden">
          {/* Narrow sidebar*/}
          <nav
            aria-label="Sidebar"
            className="block flex-shrink-0 bg-gray-800 overflow-y-auto"
          >
            <div className="relative w-20 flex flex-col p-3 space-y-3">
              {sidebarNavigation.map((item) => (
                <NavLink
                  key={item.name}
                  to={item.koboSelectionRequired && !selectedKobo.name ? '#' : item.href}
                  disabled={item.koboSelectionRequired && !selectedKobo.name}
                  className={classNames(
                    determineExtraClasses(location.pathname, item, selectedKobo),
                    "flex-shrink-0 inline-flex flex-col items-center justify-center h-14 w-14 rounded-lg text-xs font-medium"
                  )}
                >
                  <item.icon className="pb-1 h-6 w-6" aria-hidden="true" />
                  <span>{item.name}</span>
                </NavLink>
              ))}
            </div>
          </nav>

          {props.children}
        </div>
      </div>
    </>
  );
}
