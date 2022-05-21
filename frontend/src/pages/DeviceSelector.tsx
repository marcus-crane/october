import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/Navbar";
import { toast } from "react-toastify";
import { backend } from '../../wailsjs/go/models'
import { DetectKobos, SelectKobo, PromptForLocalDBPath } from '../../wailsjs/go/backend/Backend'

export default function DeviceSelector() {
  const navigate = useNavigate();
  const [devices, setDevices] = useState<Array<backend.Kobo>>([]);

  // We need a "flat" state to check if we should do a re-render or not.
  // Device metadata itself is not flat so React can't easily tell if state is "new" or not
  // ie; we'd have to recurse through devices to see if anything is new.
  // As a proxy, we'll use the number of connected devices although this will mean if you unplug a device
  // and plug in a new one, without closing the application, you'd have to manually click update.
  // That really shouldn't be an issue though and works good enough for now.
  useEffect(() => detectDevices(), [devices.length])

  function detectDevices() {
    DetectKobos()
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

  function selectDevice(path: string) {
    SelectKobo(path)
      .then(error => {
        if (error === null) {
          navigate("/overview")
        } else {
          console.log(error)
          toast.error("Something went wrong selecting your Kobo")
        }
      })
      .catch(err => toast.error(err))
  }

  function selectLocalDatabase() {
    PromptForLocalDBPath()
      .then(error => {
        if (error === null) {
          navigate("/overview")
        } else {
          console.log(error)
          toast.error("Something went wrong selecting your local sqlite database")
        }
      })
      .catch(err => toast.error(err))
  }

  return (
    <div className="bg-gray-100 dark:bg-gray-800 min-h-screen">
      <Navbar />
      <div className="flex flex-col items-center pt-24 px-24">
        <div className="space-y-4 text-center">
          <h1 className="text-3xl font-bold">Select your Kobo</h1>
          <button onClick={detectDevices}>Don't see your device? Click here to refresh device list.</button>
          <ul>
            {devices.map(device => {
              let description = `${device.storage} GB · ${device.display_ppi} PPI`
              if (!device.name) {
                description = "October did not recognise this Kobo but it's safe to continue"
              }
              return (
                <li key={device.mnt_path}>
                  <button onClick={() => selectDevice(device.mnt_path)} className="w-full bg-purple-200 hover:bg-purple-300 group block rounded-lg p-4 mb-2 cursor-pointer">
                    <dl>
                      <div>
                        <dt className="sr-only">Title</dt>
                        <dd className="border-gray leading-6 font-medium text-black">
                          {device.name || "Unknown Kobo"}
                        </dd>
                        <dt className="sr-only">System Specifications</dt>
                        <dd className="text-xs text-gray-600 dark:text-gray-400">
                          {description}
                        </dd>
                      </div>
                    </dl>
                  </button>
                </li>
              )
            })}
            <li>
              <button onClick={selectLocalDatabase} className="w-full bg-purple-200 hover:bg-purple-300 group block rounded-lg p-4 mb-2 cursor-pointer">
                <dl>
                  <div>
                    <dt className="sr-only">Title</dt>
                    <dd className="border-gray leading-6 font-medium text-black">
                      Load in a local Kobo database (advanced)
                    </dd>
                    <dt className="sr-only">Description</dt>
                    <dd className="text-xs text-gray-600 dark:text-gray-400">
                      Provide an instance of KoboReader.sqlite3
                    </dd>
                  </div>
                </dl>
              </button>
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}
