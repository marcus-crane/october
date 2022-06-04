import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-hot-toast";
import {
  DetectKobos,
  SelectKobo,
  PromptForLocalDBPath,
} from "../../wailsjs/go/backend/Backend";
import { PlusIcon, DatabaseIcon } from "@heroicons/react/solid";

export default function DeviceSelection() {
  const navigate = useNavigate();
  const [devices, setDevices] = useState([]);

  // We need a "flat" state to check if we should do a re-render or not.
  // Device metadata itself is not flat so React can't easily tell if state is "new" or not
  // ie; we'd have to recurse through devices to see if anything is new.
  // As a proxy, we'll use the number of connected devices although this will mean if you unplug a device
  // and plug in a new one, without closing the application, you'd have to manually click update.
  // That really shouldn't be an issue though and works good enough for now.
  useEffect(() => detectDevices(), [devices.length]);

  function detectDevices() {
    DetectKobos()
      .then((devices) => {
        console.log(devices);
        if (devices == null) {
          return;
        }
        toast.success(`${devices.length} kobos detected`);
        setDevices(devices);
      })
      .catch((err) => {
        toast.error(err);
      });
  }

  function selectDevice(path) {
    SelectKobo(path)
      .then((error) => {
        if (error === null) {
          navigate("/overview");
        } else {
          console.log(error);
        }
      })
      .catch((err) => toast.error(err));
  }

  function selectLocalDatabase() {
    PromptForLocalDBPath()
      .then((error) => {
        if (error === null) {
          navigate("/overview");
        } else {
          console.log(error);
          toast.error(
            "Something went wrong selecting your local SQLite database"
          );
        }
      })
      .catch((err) => toast.error(err));
  }
  return (
    <div className="mt-10 w-1/2 m-auto">
      <h1 class="text-3xl">Select your Kobo to continue</h1>
      <p class="text-lg"></p>
      <button class="text-xs" onClick={detectDevices}>Just plugged in your device? Click here to refresh the list below.</button>
      <h3 className="text-xs font-semibold text-gray-500 uppercase tracking-wide pt-10">
        {devices.length} Kobo eReader{devices.length !== 1 ? 's have' : ' has'} been detected so far
      </h3>
      <ul role="list" className="mt-4 grid grid-cols-1 gap-4 sm:grid-cols-2">
        {devices.map((device, deviceIdx) => {
          let description = `${device.storage} GB · ${device.display_ppi} PPI`
          if (!device.name) {
            description = "October did not recognise this Kobo but it's safe to continue"
          }
          return (
          <li key={deviceIdx}>
            <button
              type="button"
              className="group p-2 w-full flex items-center justify-between rounded-full border border-gray-300 shadow-sm space-x-3 text-left hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              <span className="min-w-0 flex-1 flex items-center space-x-3">
                <span className="block flex-shrink-0"></span>
                <span className="block min-w-0 flex-1">
                  <span className="block text-sm font-medium text-gray-900 truncate">
                    {device.name || "Unknown Kobo"}
                  </span>
                  <span className="block text-sm font-medium text-gray-500 truncate">
                    {description}
                  </span>
                </span>
              </span>
              <span className="flex-shrink-0 h-10 w-10 inline-flex items-center justify-center">
                <PlusIcon
                  className="h-5 w-5 text-gray-400 group-hover:text-gray-500"
                  aria-hidden="true"
                />
              </span>
            </button>
          </li>
        )})}
        <li key="localdb">
          <button
            type="button"
            className="group p-2 w-full flex items-center justify-between rounded-full border border-dashed border-gray-400 shadow-sm space-x-3 text-left hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            <span className="min-w-0 flex-1 flex items-center space-x-3">
              <span className="block flex-shrink-0"></span>
              <span className="block min-w-0 flex-1">
                <span className="block text-sm font-medium text-gray-900 truncate">
                  Read from a Kobo backup
                </span>
                <span className="block text-sm font-medium text-gray-500 truncate">
                  Select a copy of KoboReader.sqlite
                </span>
              </span>
            </span>
            <span className="flex-shrink-0 h-10 w-10 inline-flex items-center justify-center">
              <PlusIcon
                className="h-5 w-5 text-gray-400 group-hover:text-gray-500"
                aria-hidden="true"
              />
            </span>
          </button>
        </li>
      </ul>
    </div>
  );
}
