import React, { useState, useEffect } from 'react';
import Navbar from "../components/Navbar";
import { toast } from "react-hot-toast"
import { CountDeviceBookmarks } from "../../wailsjs/go/backend/Kobo";
import { GetSelectedKobo, ForwardToReadwise } from "../../wailsjs/go/backend/Backend";
import { CheckReadwiseConfig } from "../../wailsjs/go/backend/Settings";

export default function Overview(props) {
  const [readwiseConfigured, setReadwiseConfigured] = useState(false)
  const [selectedKobo, setSelectedKobo] = useState({})
  const [highlightCount, setHighlightCount] = useState(0)

  useEffect(() => {
    GetSelectedKobo()
      .then(kobo => setSelectedKobo(kobo))
      .catch(err => toast.error(err))
  }, [selectedKobo.mnt_path])

  useEffect(() => {
    CountDeviceBookmarks()
      .then(bookmarkCount => setHighlightCount(bookmarkCount))
      .catch(err => toast.error(err))
  }, [highlightCount])

  useEffect(() => {
    CheckReadwiseConfig()
      .then(result => setReadwiseConfigured(result))
      .catch(err => toast.error(err))
  })

  function promptReadwise() {
    toast("In order to upload to Readwise, you need to configure your token on the Settings page!")
  }

  function syncWithReadwise() {
    const toastId = toast.loading("Preparing your highlights...")
    ForwardToReadwise()
      .then(res => {
        if (typeof (res) == "number") {
          toast.success(`Successfully forwarded ${res} highlights to Readwise`, { id: toastId })
        } else {
          toast.error(`There was a problem sending your highlights: ${res.message}`, { id: toastId })
        }
      })
      .catch(err => {
        if (err.includes("401")) {
          toast.error("Received 401 Unauthorised from Readwise. Is your access token correct?", { id: toastId })
        } else if (err.includes("failed to upload covers")) {
          toast.error(err, { id: toastId })
        } else {
          toast.error(`There was a problem sending your highlights: ${err}`, { id: toastId })
        }
      })
  }

  return (
    <div className="bg-gray-100 dark:bg-gray-800 ">
      <Navbar />
      <div className="min-h-screen flex items-center justify-center pb-24 px-24 grid grid-cols-2 gap-14">
        <div className="space-y-2">
          <p className="text-center text-xs text-gray-600 dark:text-gray-400">Currently connected</p>
          <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-300">
            {selectedKobo.name}
          </h2>
          {selectedKobo.storage !== 0 && selectedKobo.display_ppi !== 0 && (
            <p className="mt-0 text-center text-sm text-gray-600 dark:text-gray-400">
              {selectedKobo.storage} GB · {selectedKobo.display_ppi} PPI
            </p>
          )}
          {selectedKobo.storage === 0 && selectedKobo.display_ppi === 0 && (
            <p className="mt-0 text-center text-sm text-gray-600 dark:text-gray-400">
              Dang, that's some hardcore hacker stuff!
            </p>
          )}
        </div>
        <div className="space-y-4 text-center">
          <h3 className="text-md font-medium dark:text-gray-300">What would you like to do?</h3>
          <ul>
            <li>
              <button onClick={readwiseConfigured ? syncWithReadwise : promptReadwise} className="w-full bg-purple-200 hover:bg-purple-300 dark:bg-purple-300 group block rounded-lg p-4 mb-2 cursor-pointer">
                <dl>
                  <div>
                    <dt className="sr-only">Title</dt>
                    <dd className="border-gray leading-6 font-medium text-black">
                      Sync your highlights with Readwise
                    </dd>
                    <dt className="sr-only">Description</dt>
                    <dd className="text-xs text-gray-600">
                      Your Kobo is currently home to {highlightCount} highlights
                    </dd>
                  </div>
                </dl>
              </button>
            </li>
          </ul>
        </div>
      </div>
    </div>
  )
}
