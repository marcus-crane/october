import { Fragment, useState, useEffect } from "react";
import { NavLink, useLocation } from "react-router-dom"
import { GetSelectedKobo } from "../../wailsjs/go/backend/Backend";
import { Disclosure, Dialog, Menu, Transition } from "@headlessui/react";
import { ChevronDownIcon, SearchIcon } from "@heroicons/react/solid";
import {
  BookOpenIcon,
  BookmarkIcon,
  RefreshIcon,
  CogIcon,
  MapIcon,
  StatusOnlineIcon,
} from "@heroicons/react/outline";
import logo from '../logo.png'

const sidebarNavigation = [
  { name: "Load", href: "/", icon: MapIcon, koboSelectionRequired: false },
  { name: "Sync", href: "/overview", icon: RefreshIcon, koboSelectionRequired: true },
  { name: "Books", href: "#", icon: BookOpenIcon, koboSelectionRequired: true },
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
  if (item.koboSelectionRequired && !selectedKobo.mnt_path) {
    classes.push("cursor-not-allowed")
  }
  return classes.join(" ")
}

export default function Layout(props) {
  const [selectedKobo, setSelectedKobo] = useState({})
  useEffect(() => {
    GetSelectedKobo()
      .then(kobo => setSelectedKobo(kobo))
      .catch(err => console.log('No Kobo found'))
  }, [selectedKobo.mnt_path])
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
              className="flex items-center justify-center h-16 w-20 bg-indigo-500 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-600"
            >
              <img
                className="h-10 w-auto"
                src={logo}
                alt="Workflow"
              />
            </a>
          </div>

          {/* Desktop nav area */}
          <div className="min-w-0 flex-1 flex items-center justify-between w-full">
            <Disclosure as="nav" className="flex-shrink-0 bg-indigo-600 w-full">
              {({ open }) => (
                <>
                  <div className="max-w-7xl mx-auto px-2 sm:px-4 lg:px-8">
                    <div className="relative flex items-center justify-between h-16">
                      {/* Logo section */}
                      <div className="flex items-center px-2 lg:px-0 xl:w-64">
                        <div className="flex-shrink-0"></div>
                      </div>

                      {/* Search section */}
                      {/* <div className="flex-1 flex justify-center lg:justify-end">
                        <div className="w-full px-2 lg:px-6">
                          <label htmlFor="search" className="sr-only">
                            Search books
                          </label>
                          <div className="relative text-indigo-200 focus-within:text-gray-400">
                            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                              <SearchIcon
                                className="h-5 w-5"
                                aria-hidden="true"
                              />
                            </div>
                            <input
                              id="search"
                              name="search"
                              className="block w-full pl-10 pr-3 py-2 border border-transparent rounded-md leading-5 bg-indigo-400 bg-opacity-25 text-indigo-100 placeholder-indigo-200 focus:outline-none focus:bg-white focus:ring-0 focus:placeholder-gray-400 focus:text-gray-900 sm:text-sm"
                              placeholder="Search books"
                              type="search"
                            />
                          </div>
                        </div>
                      </div> */}
                      {/* Links section */}
                      <div className="block lg:w-80">
                        <div className="flex items-center justify-end">
                          <div className="flex">
                            {/* <a
                              href="#"
                              className="px-3 py-2 rounded-md text-sm font-medium text-indigo-200 hover:text-white"
                            >
                              Documentation
                            </a> */}
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </>
              )}
            </Disclosure>
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
                  to={item.koboSelectionRequired && !selectedKobo.mnt_path ? '#' : item.href}
                  disabled={item.koboSelectionRequired && !selectedKobo.mnt_path}
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
