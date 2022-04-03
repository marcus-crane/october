import React, { useState, useEffect } from 'react';
import Navbar from "../components/Navbar";
import { toast } from "react-toastify";

export default function Settings() {
  const [token, setToken] = useState("")
  const [coversUploading, setCoversUploading] = useState(false)
  const [tokenInput, setTokenInput] = useState("")

  useEffect(() => {
    getReadwiseToken()
  }, [token])

  useEffect(() => {
    getCoverUploadStatus()
  }, [coversUploading])

  function getReadwiseToken() {
    window.go.main.KoboService.GetReadwiseToken()
      .then(token => {
        setToken(token)
        setTokenInput(token)
      })
      .catch(err => toast.error(err))
  }

  function getCoverUploadStatus() {
    window.go.main.KoboService.GetCoverUploadStatus()
      .then(status => {
        setCoversUploading(status)
      })
      .catch(err => toast.error(err))
  }

  function saveReadwiseToken() {
    console.log("calling save token")
    window.go.main.KoboService.SetReadwiseToken(tokenInput)
      .then(saveIssue => {
        console.log(saveIssue)
        if (saveIssue === null) {
          setToken(token)
          toast.success("Readwise token saved successfully")
        } else {
          throw saveIssue
        }
      })
      .catch(err => toast.error(err))
  }

  function saveCoverUploadStatus(coversChecked) {
    window.go.main.KoboService.SetCoverUploadStatus(coversChecked)
      .then(saveIssue => {
        console.log(saveIssue)
        if (saveIssue === null) {
          setCoversUploading(coversChecked)
          toast.success("Cover upload status saved successfully")
        } else {
          throw saveIssue
        }
      })
      .catch(err => toast.error(err))
  }

  function checkReadwiseTokenValid() {
    const toastId = toast.loading("Contacting the Readwise API...")
    window.go.main.KoboService.CheckTokenValidity()
      .then(apiError => {
        if (apiError === null) {
          toast.update(toastId, { render: "Successfully authenticated against the Readwise API", type: "success", isLoading: false, autoClose: 2000 })
        } else {
          throw apiError
        }
      })
      .catch(err => toast.update(toastId, { render: err.message(), type: "error", isLoading: false, autoClose: 2000 }))
  }

  return (
    <div className="bg-gray-100 dark:bg-gray-800 ">
      <Navbar />
      <div className="min-h-screen flex items-center justify-center pb-24 px-24 grid grid-cols-2 gap-14">
        <div className="space-y-2">
          <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
            Settings
          </h2>
        </div>
        <div className="space-y-4">
          <div className="bg-white shadow sm:rounded-lg">
            <div className="px-4 py-5 sm:p-6">
              <h3 className="text-lg leading-6 font-medium text-gray-900">Set your Readwise access token</h3>
              <div className="mt-2 max-w-xl text-sm text-gray-500">
                <p>You can find your access token at https://readwise.io/access_token</p>
              </div>
              <form onSubmit={(e) => e.preventDefault()} className="mt-5 sm:flex sm:items-center">
                <div className="w-full sm:max-w-xs">
                  <input
                    onChange={(e) => setTokenInput(e.target.value)}
                    type="text"
                    name="token"
                    id="token"
                    className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
                    placeholder="Your access token goes here"
                    value={tokenInput}
                  />
                </div>
                <button
                  onClick={saveReadwiseToken}
                  type="submit"
                  className="mt-3 w-full inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
                >
                  Save
                </button>
              </form>
            </div>
          </div>
          <div className="bg-white shadow sm:rounded-lg">
            <div className="px-4 py-5 sm:p-6">
              <h3 className="text-lg leading-6 font-medium text-gray-900">Validate your Readwise token</h3>
              <div className="mt-2 max-w-xl text-sm text-gray-500">
                <p>Make a call to the Readwise API to check that your token is valid and correctly entered</p>
              </div>
              <form onSubmit={(e) => e.preventDefault()} className="mt-5 sm:flex sm:items-center">
                <button
                  onClick={checkReadwiseTokenValid}
                  type="submit"
                  className="mt-3 w-full inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:text-sm"
                >
                  Validate Readwise token
                </button>
              </form>
            </div>
          </div>
          <div className="shadow overflow-hidden sm:rounded-md">
            <div className="px-4 py-5 bg-white space-y-6 sm:p-6">
              <fieldset>
                <legend className="text-base font-medium text-gray-900">Kobo metadata upload</legend>
                <div className="mt-4 space-y-4">
                  <div className="flex items-start">
                    <div className="flex items-center h-5">
                      <input
                        // TODO: This probably causes the render method to seize up
                        onInput={e => saveCoverUploadStatus(!e.currentTarget.checked)}
                        checked={coversUploading}
                        id="comments"
                        name="comments"
                        type="checkbox"
                        className="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
                      />
                    </div>
                    <div className="ml-3 text-sm">
                      <label htmlFor="comments" className="font-medium text-gray-700">
                        Upload covers
                      </label>
                      <p className="text-gray-500">This will slow down the upload process a bit. It also requires you to have configured Calibre correctly!</p>
                    </div>
                  </div>
                </div>
              </fieldset>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
