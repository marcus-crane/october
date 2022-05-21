import { useState, useEffect } from "react";
import Navbar from "../components/Navbar";
import { toast } from "react-toastify";
import { backend } from "../../wailsjs/go/models"
import { GetSettings } from "../../wailsjs/go/backend/Backend";

export default function Settings() {
  const [loaded, setLoadState] = useState(false)
  const [token, setToken] = useState("");
  const [coversUploading, setCoversUploading] = useState(
    false
  );
  const [tokenInput, setTokenInput] = useState("");

  useEffect(() => {
    GetSettings()
      .then(settings => {
        setLoadState(true)
        setToken(settings.readwise_token)
        setTokenInput(settings.readwise_token)
        setCoversUploading(settings.upload_covers)
      })
  }, [loaded])

  //   function saveReadwiseToken() {
  //     console.log("calling save token")
  //     window.go.main.KoboService.SetReadwiseToken(tokenInput)
  //       .then(saveIssue => {
  //         console.log(saveIssue)
  //         if (saveIssue === null) {
  //           setToken(token)
  //           toast.success("Readwise token saved successfully")
  //         } else {
  //           throw saveIssue
  //         }
  //       })
  //       .catch(err => toast.error(err))
  //   }

  return (
    <div className="bg-gray-100 dark:bg-gray-800 min-h-screen">
      <Navbar />
      <div className="flex flex-col items-center pt-24 px-24">
        <div className="space-y-4">
          <div className="bg-white shadow sm:rounded-lg">
            <div className="px-4 py-5 sm:p-6">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                Set your Readwise access token
              </h3>
              <div className="mt-2 max-w-xl text-sm text-gray-500">
                <p>
                  You can find your access token at
                  https://readwise.io/access_token
                </p>
              </div>
              <form
                onSubmit={(e) => e.preventDefault()}
                className="mt-5 sm:flex sm:items-center"
              >
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
                  //   onClick={saveReadwiseToken}
                  type="submit"
                  className="mt-3 w-full inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
                >
                  Save
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
