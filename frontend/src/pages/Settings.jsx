import React, { useState, useEffect } from "react";
import { toast } from "react-hot-toast";
import { BrowserOpenURL } from "../../wailsjs/runtime";
import { GetSettings } from "../../wailsjs/go/backend/Backend";
import {
  SaveToken,
  SaveCoverUploading,
} from "../../wailsjs/go/backend/Settings";
import { CheckTokenValidity } from "../../wailsjs/go/backend/Readwise";

export default function Settings() {
  const [loaded, setLoadState] = useState(false);
  const [token, setToken] = useState("");
  const [coversUploading, setCoversUploading] = useState(false);
  const [tokenInput, setTokenInput] = useState("");

  useEffect(() => {
    GetSettings().then((settings) => {
      setLoadState(true);
      setToken(settings.readwise_token);
      setTokenInput(settings.readwise_token);
      setCoversUploading(settings.upload_covers);
    });
  }, [loaded]);

  function saveToken() {
    setToken(tokenInput);
    SaveToken(tokenInput);
    toast.success("Your changes have been saved");
  }

  function saveCoverUploads() {
    setCoversUploading(!coversUploading);
    SaveCoverUploading(!coversUploading);
    toast.success("Your changes have been saved");
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

  return (
    <main className="w-full">
      <div className="bg-gray-100 dark:bg-gray-800 ">
        <div className="min-h-screen flex pt-12 justify-center pb-24 px-24 gap-14">
          <div className="space-y-4">
            <h1 class="text-3xl">Settings</h1>
            <div className="bg-white shadow sm:rounded-lg">
              <div className="px-4 py-5 sm:p-6">
                <h3 className="text-lg leading-6 font-medium text-gray-900">
                  Set your Readwise access token
                </h3>
                <div className="mt-2 max-w-xl text-sm text-gray-500">
                  <p>
                    You can find your access token at{" "}
                    <button
                      className="text-gray-600 underline"
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
                      className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
                      placeholder="Your access token goes here"
                      value={tokenInput}
                    />
                  </div>
                  <div className="w-full mt-4 sm:flex flex-row">
                    <button
                      onClick={checkTokenValid}
                      type="submit"
                      className="mt-3 w-full inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm font-medium rounded-md text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 sm:mt-0 sm:text-sm"
                    >
                      Validate
                    </button>
                    <button
                      onClick={saveToken}
                      type="submit"
                      className="mt-3 w-full inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:text-sm"
                    >
                      Save
                    </button>
                  </div>
                </form>
              </div>
            </div>
            <div className="shadow overflow-hidden sm:rounded-md">
              <div className="px-4 py-5 bg-white space-y-6 sm:p-6">
                <fieldset>
                  <legend className="text-base font-medium text-gray-900">
                    Kobo metadata upload
                  </legend>
                  <div className="mt-4 space-y-4">
                    <div className="flex items-start">
                      <div className="flex items-center h-5">
                        <input
                          // TODO: This probably causes the render method to seize up
                          onInput={(e) =>
                            saveCoverUploads(!e.currentTarget.checked)
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
                          className="font-medium text-gray-700"
                        >
                          Upload covers
                        </label>
                        <p className="text-gray-500">
                          This will slow down the upload process a bit. It also
                          requires you to have configured Calibre correctly!
                        </p>
                      </div>
                    </div>
                  </div>
                </fieldset>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
}
