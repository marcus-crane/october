import React, { useState, useEffect } from 'react';
import Navbar from "../Components/Navbar";

export default function Overview(props) {
  const [readwiseConfigured, setReadwiseConfigured] = useState(true)
  const [selectedKobo, setSelectedKobo] = useState({})
  const [highlightCount, setHighlightCount] = useState(0)

  useEffect(() => {
    window.go.main.KoboService.GetSelectedKobo()
      .then(kobo => setSelectedKobo(kobo))
      .catch(err => console.log(err))
  }, [selectedKobo.mnt_path])

  useEffect(() => {
    window.go.main.KoboService.CountDeviceBookmarks()
      .then(bookmarkCount => setHighlightCount(bookmarkCount))
      .catch(err => console.log(err))
  }, [highlightCount])

  function syncWithReadwise() {
    console.log("hello")
  }

  function exportDatabase() {
    console.log("ah")
  }

  return (
    <div className="bg-gray-100 dark:bg-gray-800 ">
      <Navbar />
      <div className="min-h-screen flex items-center justify-center pb-24 px-24 grid grid-cols-2 gap-14">
        <div className="space-y-2">
          <p className="text-center text-xs text-gray-600 dark:text-gray-400">Currently connected</p>
          <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
            {selectedKobo.name}
          </h2>
          <p className="mt-0 text-center text-sm text-gray-600 dark:text-gray-400">
            {selectedKobo.storage} GB · {selectedKobo.display_ppi} PPI
          </p>
        </div>
        <div className="space-y-4 text-center">
          <h3 className="text-md font-medium">What would you like to do?</h3>
          <ul>
            <li>
              <a onClick={syncWithReadwise} className="bg-purple-200 hover:bg-purple-300 group block rounded-lg p-4 mb-2 cursor-pointer">
                <dl>
                  <div>
                    <dt className="sr-only">Title</dt>
                    <dd className="border-gray leading-6 font-medium text-black">
                      Sync your highlights with Readwise
                    </dd>
                    <dt className="sr-only">Description</dt>
                    <dd className="text-xs text-gray-600 dark:text-gray-400">
                      Your Kobo is currently home to {highlightCount} highlights
                    </dd>
                  </div>
                </dl>
              </a>
            </li>
            <li>
              <a onClick={exportDatabase} className="bg-purple-200 hover:bg-purple-300 group block rounded-lg p-4 cursor-pointer">
                <dl>
                  <div>
                    <dt className="sr-only">Title</dt>
                    <dd className="border-gray leading-6 font-medium text-black">
                      Export KoboReader.sqlite
                    </dd>
                    <dt className="sr-only">Description</dt>
                    <dd className="text-xs text-gray-600 dark:text-gray-400">
                      Create a local copy of your Kobo database
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
