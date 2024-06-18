import React, { Fragment, useState, useEffect, useRef } from 'react';
import Navbar from "../components/Navbar";
import { toast } from "react-hot-toast"
import { CountDeviceBookmarks } from "../../wailsjs/go/backend/Kobo";
import { GetSettings, GetSelectedKobo, ForwardToReadwise } from "../../wailsjs/go/backend/Backend";

export default function Overview(props) {
  const [settingsLoaded, setSettingsLoaded] = useState(false)
  const [readwiseConfigured, setReadwiseConfigured] = useState(false)
  const [uploadStorePromptSeen, setUploadStorePromptSeen] = useState(false)
  const [selectedKobo, setSelectedKobo] = useState({})
  const [highlightCounts, setHighlightCounts] = useState({})
  const [storeUploadWarningOpen, setStoreUploadWarningOpen] = useState(false)
  const [uploadStoreHighlights, setUploadStoreHighlights] = useState(false)

  const cancelButtonRef = useRef(null)

  useEffect(() => {
    GetSelectedKobo()
      .then(kobo => setSelectedKobo(kobo))
      .catch(err => toast.error(err))
  }, [selectedKobo.mnt_path])

  useEffect(() => {
    CountDeviceBookmarks()
      .then(bookmarkCounts => setHighlightCounts(bookmarkCounts))
      .catch(err => toast.error(err))
  }, [highlightCounts.total])

  useEffect(() => {
    GetSettings().then((settings) => {
      setSettingsLoaded(true);
      setUploadStorePromptSeen(settings.upload_store_prompt_shown)
      setReadwiseConfigured(settings.readwise_token !== "")
      setUploadStoreHighlights(settings.upload_store_highlights)
    });
  }, [settingsLoaded])

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
    <>
      <div className="min-h-screen bg-gray-100 dark:bg-gray-800 flex flex-col">
        <Navbar />
        <div className="flex-grow items-center justify-center pb-24 px-24 grid grid-cols-2 gap-14">
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
                <button onClick={() => {
                    if (readwiseConfigured) {
                      if (uploadStorePromptSeen || uploadStoreHighlights) {
                        syncWithReadwise()
                      } else {
                        setStoreUploadWarningOpen(true)
                      }
                    } else {
                      promptReadwise()
                    }
                  }} className="w-full bg-purple-200 hover:bg-purple-300 dark:bg-purple-300 group block rounded-lg p-4 mb-2 cursor-pointer">
                  <dl>
                    <div>
                      <dt className="sr-only">Title</dt>
                      <dd className="border-gray leading-6 font-medium text-black">
                        Sync your highlights with Readwise
                      </dd>
                      <dt className="sr-only">Description</dt>
                      <dd className="text-xs text-gray-600">
                        Your Kobo is currently home to {highlightCounts.total} highlights
                      </dd>
                      <dd className="text-xs text-gray-600">
                        {highlightCounts.sideloaded} sideloaded  · {highlightCounts.official} official
                      </dd>
                    </div>
                  </dl>
                </button>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </>
  )
}
