import React from "react";

import { HeartIcon } from "@heroicons/react/outline";
import {
  PencilIcon,
  PlusSmIcon as PlusSmIconSolid,
  ViewGridIcon as ViewGridIconSolid,
  ViewListIcon,
} from "@heroicons/react/solid";

const files = [
  {
    name: "IMG_4985.HEIC",
    size: "3.9 MB",
    source:
      "https://images.unsplash.com/photo-1582053433976-25c00369fc93?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=512&q=80",
    current: true,
  },
  // More files...
];

const currentFile = {
  name: "IMG_4985.HEIC",
  size: "3.9 MB",
  source:
    "https://images.unsplash.com/photo-1582053433976-25c00369fc93?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=512&q=80",
  information: {
    "Uploaded by": "Marie Culver",
    Created: "June 8, 2020",
    "Last modified": "June 8, 2020",
    Dimensions: "4032 x 3024",
    Resolution: "72 x 72",
  }
};

function classNames(...classes) {
  return classes.filter(Boolean).join(" ");
}

export default function Library() {
  return (
    <div className="flex-1 flex items-stretch overflow-hidden">
      <main className="flex-1 overflow-y-auto">
        <div className="pt-8 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex">
            <h1 className="flex-1 text-2xl font-bold text-gray-900">Books</h1>
            <div className="ml-6 bg-gray-100 p-0.5 rounded-lg flex items-center">
              <button
                type="button"
                className="p-1.5 rounded-md text-gray-400 hover:bg-white hover:shadow-sm focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500"
              >
                <ViewListIcon className="h-5 w-5" aria-hidden="true" />
                <span className="sr-only">Use list view</span>
              </button>
              <button
                type="button"
                className="ml-0.5 bg-white p-1.5 rounded-md shadow-sm text-gray-400 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500"
              >
                <ViewGridIconSolid className="h-5 w-5" aria-hidden="true" />
                <span className="sr-only">Use grid view</span>
              </button>
            </div>
          </div>

          {/* Gallery */}
          <section className="mt-8 pb-16" aria-labelledby="gallery-heading">
            <h2 id="gallery-heading" className="sr-only">
              Recently viewed
            </h2>
            <ul
              role="list"
              className="grid grid-cols-2 gap-x-4 gap-y-8 sm:grid-cols-3 sm:gap-x-6 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-6 xl:gap-x-8"
            >
              {files.map((file) => (
                <li key={file.name} className="relative">
                  <div
                    className={classNames(
                      file.current
                        ? "ring-2 ring-offset-2 ring-indigo-500"
                        : "focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-offset-gray-100 focus-within:ring-indigo-500",
                      "group block w-full aspect-w-10 aspect-h-7 rounded-lg bg-gray-100 overflow-hidden"
                    )}
                  >
                    <img
                      src={file.source}
                      alt=""
                      className={classNames(
                        file.current ? "" : "group-hover:opacity-75",
                        "object-cover pointer-events-none"
                      )}
                    />
                    <button
                      type="button"
                      className="absolute inset-0 focus:outline-none"
                    >
                      <span className="sr-only">
                        View details for {file.name}
                      </span>
                    </button>
                  </div>
                  <p className="mt-2 block text-sm font-medium text-gray-900 truncate pointer-events-none">
                    {file.name}
                  </p>
                  <p className="block text-sm font-medium text-gray-500 pointer-events-none">
                    {file.size}
                  </p>
                </li>
              ))}
            </ul>
          </section>
        </div>
      </main>
      <aside className="hidden w-96 bg-white p-8 border-l border-gray-200 overflow-y-auto lg:block">
        <div className="pb-16 space-y-6">
          <div>
            <div className="block w-full aspect-w-10 aspect-h-7 rounded-lg overflow-hidden">
              <img src={currentFile.source} alt="" className="object-cover" />
            </div>
            <div className="mt-4 flex items-start justify-between">
              <div>
                <h2 className="text-lg font-medium text-gray-900">
                  <span className="sr-only">Details for </span>
                  {currentFile.name}
                </h2>
                <p className="text-sm font-medium text-gray-500">
                  {currentFile.size}
                </p>
              </div>
              <button
                type="button"
                className="ml-4 bg-white rounded-full h-8 w-8 flex items-center justify-center text-gray-400 hover:bg-gray-100 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500"
              >
                <HeartIcon className="h-6 w-6" aria-hidden="true" />
                <span className="sr-only">Favorite</span>
              </button>
            </div>
          </div>
          <div>
            <h3 className="font-medium text-gray-900">Information</h3>
            <dl className="mt-2 border-t border-b border-gray-200 divide-y divide-gray-200">
              {Object.keys(currentFile.information).map((key) => (
                <div
                  key={key}
                  className="py-3 flex justify-between text-sm font-medium"
                >
                  <dt className="text-gray-500">{key}</dt>
                  <dd className="text-gray-900">
                    {currentFile.information[key]}
                  </dd>
                </div>
              ))}
            </dl>
          </div>
          <div>
            <h3 className="font-medium text-gray-900">Description</h3>
            <div className="mt-2 flex items-center justify-between">
              <p className="text-sm text-gray-500 italic">
                Add a description to this image.
              </p>
              <button
                type="button"
                className="bg-white rounded-full h-8 w-8 flex items-center justify-center text-gray-400 hover:bg-gray-100 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500"
              >
                <PencilIcon className="h-5 w-5" aria-hidden="true" />
                <span className="sr-only">Add description</span>
              </button>
            </div>
          </div>
          <div className="flex">
            <button
              type="button"
              className="flex-1 bg-indigo-600 py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Download
            </button>
            <button
              type="button"
              className="flex-1 ml-3 bg-white py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Delete
            </button>
          </div>
        </div>
      </aside>
    </div>
  );
}
