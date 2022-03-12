import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import Navbar from "../components/Navbar"
import logo from '../logo.png'
import { toast } from "react-toastify";

export default function DeviceSelector() {
  const navigate = useNavigate()
  const [devices, setDevices] = useState([])

  // We need a "flat" state to check if we should do a re-render or not.
  // Device metadata itself is not flat so React can't easily tell if state is "new" or not
  // ie; we'd have to recurse through devices to see if anything is new.
  // As a proxy, we'll use the number of connected devices although this will mean if you unplug a device
  // and plug in a new one, without closing the application, you'd have to manually click update.
  // That really shouldn't be an issue though and works good enough for now.
  useEffect(() => detectDevices(), [devices.length])

  function detectDevices() {
    window.go.main.KoboService.DetectKobos()
      .then(devices => {
        console.log(devices)
        if (devices == null) {
          toast.info("No devices were found")
          return
        }
        setDevices(devices)
      })
      .catch(err => {
        toast.error(err)
      })
  }

  function selectDevice(path) {
    console.log(path)
    window.go.main.KoboService.SelectKobo(path)
      .then(success => {
        if (success === true) {
          navigate("/overview")
        } else {
          toast.error("Something went wrong selecting your Kobo")
        }
      })
      .catch(err => toast.error(err))
  }
  return (
    <div className="bg-gray-100 dark:bg-gray-800 ">
      <Navbar />
      <div className="min-h-screen flex items-center justify-center pb-24 px-24 grid grid-cols-2 gap-14">
        <div className="space-y-2">
          <img
            className="mx-auto h-36 w-auto logo-animation"
            src={logo}
            alt="The October logo, which is a cartoon octopus reading a book."
          />
          <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
            October
          </h2>
          <p className="mt-0 text-center text-sm text-gray-600 dark:text-gray-400">
            Easily access your Kobo highlights
          </p>
        </div>
        <div className="space-y-4 text-center">
          <h1 className="text-3xl font-bold">Select your Kobo</h1>
          <button onClick={detectDevices}>Don't see your device? Click here to refresh device list.</button>
          <ul>
            {devices.map(device => (
              <li key={device.mnt_path}>
                <button onClick={() => selectDevice(device.mnt_path)} className="w-full bg-purple-200 hover:bg-purple-300 group block rounded-lg p-4 mb-2 cursor-pointer">
                  <dl>
                    <div>
                      <dt className="sr-only">Title</dt>
                      <dd className="border-gray leading-6 font-medium text-black">
                        {device.name}
                      </dd>
                      <dt className="sr-only">System Specifications</dt>
                      <dd className="text-xs text-gray-600 dark:text-gray-400">
                        {device.storage} GB · {device.display_ppi} PPI
                      </dd>
                    </div>
                  </dl>
                </button>
              </li>
            ))}
            <li>
              <a onClick={selectLocalDatabase} className="bg-purple-200 hover:bg-purple-300 group block rounded-lg p-4 cursor-pointer">
                <dl>
                  <div>
                    <dt className="sr-only">Title</dt>
                    <dd className="border-gray leading-6 font-medium text-black">
                      Pick a local Kobo database
                    </dd>
                    <dt className="sr-only">Description</dt>
                    <dd className="text-xs text-gray-600 dark:text-gray-400">
                      Provide your own instance of KoboReader.sqlite3
                    </dd>
                  </div>
                </dl>
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  )
}
