import React, { Fragment, useState, useEffect, useRef } from 'react';
import Navbar from "../components/Navbar";
import { toast } from "react-hot-toast"
import { CountDeviceBookmarks } from "../../wailsjs/go/backend/Kobo";
import { GetSettings, GetSelectedKobo, ForwardToReadwise } from "../../wailsjs/go/backend/Backend";
import { MarkUploadStorePromptShown } from "../../wailsjs/go/backend/Settings";
import { Dialog, Transition } from '@headlessui/react'
import { ExclamationTriangleIcon } from '@heroicons/react/24/outline'

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
                <button onClick={() => {
                    if (readwiseConfigured) {
                      console.log("Readwise configured")
                      if (uploadStorePromptSeen || uploadStoreHighlights) {
                        console.log("Syncing prompt seen")
                        syncWithReadwise()
                      } else {
                        console.log("Syncing prompt not seen")
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
      <Transition.Root show={storeUploadWarningOpen} as={Fragment}>
        <Dialog as="div" className="relative z-10" initialFocus={cancelButtonRef} onClose={setStoreUploadWarningOpen}>
          <Transition.Child
            as={Fragment}
            enter="ease-out duration-300"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
          </Transition.Child>

          <div className="fixed inset-0 z-10 overflow-y-auto">
            <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
              <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                enterTo="opacity-100 translate-y-0 sm:scale-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              >
                <Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6">
                  <div className="sm:flex sm:items-start">
                    <div className="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                      <ExclamationTriangleIcon className="h-6 w-6 text-red-600" aria-hidden="true" />
                    </div>
                    <div className="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                      <Dialog.Title as="h3" className="text-base font-semibold leading-6 text-gray-900">
                        Want to sync highlights from store-purchased books?
                      </Dialog.Title>
                      <div className="mt-2">
                        <p className="text-sm text-gray-500 mb-2">
                          You appear to have some highlights from officially purchased titles.
                        </p>
                        <p className="text-sm text-gray-500 mb-2">
                          If you would like to use October to sync highlights from the Kobo store, you can do this from the Settings page.
                        </p>
                        <p className="text-sm text-gray-500 mb-2">
                          This functionality is disabled by default in order to avoid duplicate highlights for users of the official Readwise Kobo integration.
                        </p>
                        <p className="text-sm text-gray-500 mb-2 italic">
                          This message won't be shown again once accepted.
                        </p>
                      </div>
                    </div>
                  </div>
                  <div className="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
                    <button
                      type="button"
                      className="inline-flex w-full justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto"
                      onClick={() => {
                        setStoreUploadWarningOpen(false)
                        MarkUploadStorePromptShown()
                        syncWithReadwise()
                      }}
                    >
                      Understood
                    </button>
                  </div>
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </div>
        </Dialog>
      </Transition.Root>
    </>
  )
}
