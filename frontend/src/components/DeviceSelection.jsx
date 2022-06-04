import React from "react";

export default function DeviceSelection() {
  return (
    <div className="mt-10">
      <h3 className="text-xs font-semibold text-gray-500 uppercase tracking-wide">
        Detected Kobo devices
      </h3>
      <ul role="list" className="mt-4 grid grid-cols-1 gap-4 sm:grid-cols-2">
        {devices.map((person, personIdx) => (
          <li key={personIdx}>
            <button
              type="button"
              className="group p-2 w-full flex items-center justify-between rounded-full border border-gray-300 shadow-sm space-x-3 text-left hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              <span className="min-w-0 flex-1 flex items-center space-x-3">
                <span className="block flex-shrink-0">
                  <img
                    className="h-10 w-10 rounded-full"
                    src={person.imageUrl}
                    alt=""
                  />
                </span>
                <span className="block min-w-0 flex-1">
                  <span className="block text-sm font-medium text-gray-900 truncate">
                    {person.name}
                  </span>
                  <span className="block text-sm font-medium text-gray-500 truncate">
                    {person.role}
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
        ))}
      </ul>
    </div>
  );
}
