import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import Navbar from "../components/Navbar"
import logo from '../logo.png'
import {BrowserOpenURL} from "../../wailsjs/runtime";
import {
  FormatSystemDetails,
  GetPlainSystemDetails,
  GetSettings,
  NavigateExplorerToLogLocation
} from "../../wailsjs/go/backend/Backend";
import {SaveCoverUploading, SaveStoreHighlights, SaveToken} from "../../wailsjs/go/backend/Settings";
import {toast} from "react-hot-toast";
import {CheckTokenValidity} from "../../wailsjs/go/backend/Readwise";

export default function Onboarding() {
  const [loaded, setLoadState] = useState(false);
  const [onboardingComplete, setOnboardingComplete] = useState(false)
  const [token, setToken] = useState("")
  const [coversUploading, setCoversUploading] = useState(false);
  const [storeHighlights, setStoreHighlights] = useState(false); // default on as users run into this issue more than not but give ample warning
  const [tokenInput, setTokenInput] = useState("");
  const [systemDetails, setSystemDetails] = useState(
    "Fetching system details..."
  );

  useEffect(() => {
    GetSettings().then((settings) => {
      setLoadState(true);
      if (settings.readwise_token !== "") {
        setOnboardingComplete(true)
      }
      setToken(settings.readwise_token);
      setTokenInput(settings.readwise_token);
      setCoversUploading(settings.upload_covers);
      setStoreHighlights(settings.upload_store_highlights)
    });
    GetPlainSystemDetails().then((details) => setSystemDetails(details));
  }, [loaded]);

  function saveAllSettings() {
    if (tokenInput === "") {
      toast.error("Please enter your Readwise token")
      return
    }
    SaveToken(tokenInput);
    SaveCoverUploading(coversUploading);
    SaveStoreHighlights(storeHighlights);
    navigate("/selector")
  }

  function checkTokenValid() {
    toast.promise(CheckTokenValidity(tokenInput), {
      loading: "Contacting Readwise...",
      success: () => "Your API token is valid!",
      error: (err) => {
        if (err === "401 Unauthorized") {
          return "Readwise rejected your token";
        }
        return err;
      },
    });
  }

  const navigate = useNavigate()

  if (onboardingComplete) {
    navigate("/selector")
  }
  return (
    <div className="min-h-screen bg-gray-100 dark:bg-gray-800 flex flex-col">
      <Navbar />
      <div className="flex-grow items-center justify-center pb-24 px-24 grid grid-cols-2 gap-14">
        <div className="space-y-2">
          <img
            className="mx-auto h-36 w-auto logo-animation"
            src={logo}
            alt="The October logo, which is a cartoon octopus reading a book."
          />
          <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:dark:text-gray-300">
            First time setup with October
          </h2>
          <p className="mt-0 text-center text-sm text-gray-600 dark:text-gray-400">
            This should only take a minute of your time
          </p>
        </div>
        <div className="space-y-4">
          <div className="bg-white dark:bg-slate-700 shadow sm:rounded-lg">
            <div className="px-4 py-5 sm:p-6">
              <h3 className="text-lg leading-6 font-medium text-gray-900 dark:text-gray-300">
                Set your Readwise access token
              </h3>
              <div className="mt-2 max-w-xl text-sm text-gray-500 dark:text-gray-400">
                <p>
                  You can find your access token at{" "}
                  <button
                    className="text-gray-600 dark:text-gray-400 underline"
                    onClick={() =>
                      BrowserOpenURL("https://readwise.io/access_token")
                    }
                  >
                    https://readwise.io/access_token
                  </button>
                </p>
              </div>
              <form
                onSubmit={(e) => e.preventDefault()}
                className="sm:flex flex-col"
              >
                <div className="w-full mt-4 sm:flex sm:items-center">
                  <input
                    onChange={(e) => setTokenInput(e.target.value)}
                    type="text"
                    name="token"
                    id="token"
                    className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 dark:bg-gray-200 focus:bg-white rounded-md"
                    placeholder="Your access token goes here"
                    value={tokenInput}
                  />
                </div>
                <div className="w-full mt-4 sm:flex flex-row">
                  <button
                    onClick={checkTokenValid}
                    type="submit"
                    className="mt-3 w-full inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 sm:mt-0 sm:text-sm"
                  >
                    Validate
                  </button>
                </div>
              </form>
            </div>
          </div>
          <div className="bg-white dark:bg-slate-700 shadow sm:rounded-lg">
            <div className="shadow overflow-hidden sm:rounded-md">
              <div className="px-4 py-5 bg-white dark:bg-slate-700 space-y-6 sm:p-6">
                <fieldset>
                  <legend className="text-base font-medium text-gray-900 dark:text-gray-300">
                    Highlight Settings
                  </legend>
                  <div className="mt-4 space-y-4">
                    <div className="flex items-start">
                      <div className="flex items-center h-5">
                        <input
                          // TODO: This probably causes the render method to seize up
                          onInput={(e) => {
                            setStoreHighlights(!e.currentTarget.checked)
                          }}
                          checked={storeHighlights}
                          id="storeBought"
                          name="storeBought"
                          type="checkbox"
                          className="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
                        />
                      </div>
                      <div className="ml-3 text-sm">
                        <label
                          htmlFor="storeBought"
                          className="font-medium text-gray-700 dark:text-gray-300"
                        >
                          Upload highlights from store-bought books
                        </label>
                        <p className="text-red-500 dark:text-red-400">
                          WARNING: If you are using the{" "}
                          <button
                            className="text-gray-600 dark:text-gray-400 underline"
                            onClick={() =>
                              BrowserOpenURL(
                                "https://help.readwise.io/article/135-how-do-i-import-highlights-from-kobo"
                              )
                            }
                          >
                            official Readwise integration
                          </button>{" "}
                          to sync books from the Kobo store, you should <strong>disable</strong> this option or risk having duplicate highlights!
                        </p>
                      </div>
                    </div>
                  </div>
                  <div className="mt-4 space-y-4">
                    <div className="flex items-start">
                      <div className="flex items-center h-5">
                        <input
                          // TODO: This probably causes the render method to seize up
                          onInput={(e) =>
                            setCoversUploading(!e.currentTarget.checked)
                          }
                          checked={coversUploading}
                          id="comments"
                          name="comments"
                          type="checkbox"
                          className="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
                        />
                      </div>
                      <div className="ml-3 text-sm">
                        <label
                          htmlFor="comments"
                          className="font-medium text-gray-700 dark:text-gray-300"
                        >
                          Upload covers
                        </label>
                        <p className="text-gray-500 dark:text-gray-400">
                          This will slow down the upload process a bit. It also
                          requires you to have <button className="text-gray-600 dark:text-gray-400 underline" onClick={() => BrowserOpenURL("https://october.utf9k.net/prerequisites#configuring-calibre-to-sync-high-quality-covers")}>configured Calibre</button> to get the most benefit.
                        </p>
                      </div>
                    </div>
                  </div>
                </fieldset>
              </div>
            </div>
          </div>
          <div className="bg-white dark:bg-slate-700 shadow sm:rounded-lg">
            <div className="shadow overflow-hidden sm:rounded-md">
              <div className="px-4 py-5 bg-white dark:bg-slate-700 space-y-6 sm:p-6">
                <fieldset>
                  <legend className="text-base font-medium text-gray-900 dark:text-gray-300">
                    All done?
                  </legend>
                  <div className="space-y-4">
                    <div className="flex items-start">
                      <div className="w-full mt-4 sm:flex flex-row">
                        <button
                          onClick={() =>
                            FormatSystemDetails().then((details) =>
                              BrowserOpenURL(
                                `https://github.com/marcus-crane/october/issues/new?body=${encodeURI(
                                  "I have an issue with...\n\n---\n\n" + details
                                )}`
                              )
                            )
                          }
                          type="submit"
                          className="mt-3 w-full inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:mt-0 sm:text-sm"
                        >
                          Having trouble?
                        </button>
                        <button
                          onClick={saveAllSettings}
                          type="submit"
                          className="mt-3 w-full inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm font-medium rounded-md text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 sm:mt-0 sm:ml-3 sm:text-sm"
                        >
                          Complete setup
                        </button>
                      </div>
                    </div>
                  </div>
                </fieldset>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
