import React, { useState, useEffect } from "react";

import { HeartIcon } from "@heroicons/react/outline";
import {
  PencilIcon,
  ViewGridIcon as ViewGridIconSolid,
  ViewListIcon,
} from "@heroicons/react/solid";

import { ListDeviceContentWithCovers } from '../../wailsjs/go/backend/Backend'

function classNames(...classes) {
  return classes.filter(Boolean).join(" ");
}

export default function Library() {
  const [books, setBooks] = useState([])
  const [selectedBook, setSelectedBook] = useState({})

  useEffect(() => {
    ListDeviceContentWithCovers()
      .then((content) => setBooks(content))
      .catch((err) => toast.error(err));
  }, [books.length]);

  console.log(books)

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
              className="grid grid-cols-2 gap-x-4 gap-y-8 sm:grid-cols-3 sm:gap-x-6 md:grid-cols-3 lg:grid-cols-3 xl:grid-cols-4 xl:gap-x-8"
            >
              {books.map((book) => (
                <li key={book.title} className="relative">
                  <div
                    className={classNames(
                      selectedBook.title === book.title
                        ? "ring-2 ring-offset-2 ring-indigo-500"
                        : "focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-offset-gray-100 focus-within:ring-indigo-500",
                      "group block w-full aspect-w-10 aspect-h-7 rounded-lg bg-gray-100 overflow-hidden"
                    )}
                  >
                    <img
                      src={book.cover_bytes || "https://via.placeholder.com/186x283.png"}
                      alt=""
                      className={classNames(
                        selectedBook.title === book.title ? "" : "group-hover:opacity-75",
                        "object-cover pointer-events-none"
                      )}
                    />
                    <button
                      type="button"
                      onClick={() => setSelectedBook(book)}
                      className="absolute inset-0 focus:outline-none"
                    >
                      <span className="sr-only">
                        View details for {book.title}
                      </span>
                    </button>
                  </div>
                  <p className="mt-2 block text-sm font-medium text-gray-900 truncate pointer-events-none">
                    {book.title}
                  </p>
                  <p className="block text-sm font-medium text-gray-500 pointer-events-none">
                    {book.attribution}
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
            <div className="block w-full aspect-w-5 aspect-h-7 rounded-lg overflow-hidden">
              <img src={selectedBook.cover_bytes || "https://via.placeholder.com/319x412.png"} alt="" className="object-cover" />
            </div>
            <div className="mt-4 flex items-start justify-between">
              <div>
                <h2 className="text-md font-medium text-gray-900">
                  <span className="sr-only">Details for </span>
                  {selectedBook.title}
                </h2>
                <p className="text-sm font-medium text-gray-500">
                  {selectedBook.attribution}
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
          {/* <div>
            <h3 className="font-medium text-gray-900">Information</h3>
            <dl className="mt-2 border-t border-b border-gray-200 divide-y divide-gray-200">
              {Object.keys(selectedBook.name).map((key) => (
                <div
                  key={key}
                  className="py-3 flex justify-between text-sm font-medium"
                >
                  <dt className="text-gray-500">{key}</dt>
                  <dd className="text-gray-900">
                    {selectedBook.name}
                  </dd>
                </div>
              ))}
            </dl>
          </div> */}
          <div>
            <h3 className="font-medium text-gray-900">Description</h3>
            <div className="mt-2 flex items-center justify-between">
              <p className="text-sm text-gray-500 italic" dangerouslySetInnerHTML={{__html: selectedBook.description}} />
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
