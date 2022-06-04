import React, { useState, useEffect } from "react";
import Navbar from "../components/Navbar";
import { toast } from "react-hot-toast";
import { CountDeviceBookmarks, CountDeviceBooks } from "../../wailsjs/go/backend/Kobo";
import {
  GetSelectedKobo,
  ForwardToReadwise,
} from "../../wailsjs/go/backend/Backend";
import {
  DatabaseIcon,
  BookmarkIcon,
  DesktopComputerIcon,
  FolderOpenIcon
} from "@heroicons/react/solid";
import { CheckReadwiseConfig } from "../../wailsjs/go/backend/Settings";

export default function Overview(props) {
  const [readwiseConfigured, setReadwiseConfigured] = useState(false);
  const [selectedKobo, setSelectedKobo] = useState({});
  const [highlightCount, setHighlightCount] = useState(0);
  const [bookCount, setBookCount] = useState(0);

  useEffect(() => {
    GetSelectedKobo()
      .then((kobo) => setSelectedKobo(kobo))
      .catch((err) => toast.error(err));
  }, [selectedKobo.mnt_path]);

  useEffect(() => {
    CountDeviceBookmarks()
      .then((bookmarkCount) => setHighlightCount(bookmarkCount))
      .catch((err) => toast.error(err));
  }, [highlightCount]);

  useEffect(() => {
    CountDeviceBooks()
      .then((bookCount) => setBookCount(bookCount))
      .catch((err) => toast.error(err));
  }, [bookCount]);

  useEffect(() => {
    CheckReadwiseConfig()
      .then((result) => setReadwiseConfigured(result))
      .catch((err) => toast.error(err));
  });

  function promptReadwise() {
    toast(
      "In order to upload to Readwise, you need to configure your token on the Settings page!"
    );
  }

  function syncWithReadwise() {
    const toastId = toast.loading("Preparing your highlights...");
    ForwardToReadwise()
      .then((res) => {
        if (typeof res == "number") {
          toast.success(
            `Successfully forwarded ${res} highlights to Readwise`,
            { id: toastId }
          );
        } else {
          toast.error(
            `There was a problem sending your highlights: ${res.message}`,
            { id: toastId }
          );
        }
      })
      .catch((err) => {
        if (err.includes("401")) {
          toast.error(
            "Received 401 Unauthorised from Readwise. Is your access token correct?",
            { id: toastId }
          );
        } else if (err.includes("failed to upload covers")) {
          toast.error(err, { id: toastId });
        } else {
          toast.error(`There was a problem sending your highlights: ${err}`, {
            id: toastId,
          });
        }
      });
  }

  return (
    <main className="w-full">
      <div className="bg-white shadow">
        <div className="px-4 sm:px-6 lg:max-w-6xl lg:mx-auto lg:px-8">
          <div className="py-6 md:flex md:items-center md:justify-between lg:border-t lg:border-gray-200">
            <div className="flex-1 min-w-0">
              {/* Profile */}
              <div className="flex items-center">
                <img
                  className="hidden h-16 w-16 rounded-full sm:block"
                  src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2.6&w=256&h=256&q=80"
                  alt=""
                />
                <div>
                  <div className="flex items-center">
                    <img
                      className="h-16 w-16 rounded-full sm:hidden"
                      src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2.6&w=256&h=256&q=80"
                      alt=""
                    />
                    <h1 className="ml-3 text-2xl font-bold leading-7 text-gray-900 sm:leading-9 sm:truncate">
                      {selectedKobo.name}
                    </h1>
                  </div>
                  <dl className="mt-6 flex flex-col sm:ml-3 sm:mt-1 sm:flex-row sm:flex-wrap">
                    {selectedKobo.name !== "Local Database" && (
                      <>
                        <dt className="sr-only">Storage</dt>
                        <dd className="flex items-center text-sm text-gray-500 font-medium capitalize sm:mr-6">
                          <DatabaseIcon
                            className="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                            aria-hidden="true"
                          />
                          {selectedKobo.storage} GB
                        </dd>
                        <dt className="sr-only">Display Resolution</dt>
                        <dd className="mt-3 flex items-center text-sm text-gray-500 font-medium sm:mr-6 sm:mt-0 capitalize">
                          <DesktopComputerIcon
                            className="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                            aria-hidden="true"
                          />
                          {selectedKobo.display_ppi} PPI
                        </dd>
                      </>
                    )}
                    <dt className="sr-only">Books</dt>
                    <dd className="mt-3 flex items-center text-sm text-gray-500 font-medium sm:mr-6 sm:mt-0 capitalize">
                      <BookmarkIcon
                        className="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                        aria-hidden="true"
                      />
                      {bookCount} books
                    </dd>
                    <dt className="sr-only">Highlights</dt>
                    <dd className="mt-3 flex items-center text-sm text-gray-500 font-medium sm:mr-6 sm:mt-0 capitalize">
                      <BookmarkIcon
                        className="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                        aria-hidden="true"
                      />
                      {highlightCount} highlights
                    </dd>
                    <dt className="sr-only">Mount Path</dt>
                    <dd className="mt-3 flex items-center text-sm text-gray-500 font-medium sm:mr-6 sm:mt-0 capitalize">
                      <FolderOpenIcon
                        className="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                        aria-hidden="true"
                      />
                      {selectedKobo.mnt_path}
                    </dd>
                  </dl>
                </div>
              </div>
            </div>
            <div className="mt-6 flex space-x-3 md:mt-0 md:ml-4">
              <button
                type="button"
                className="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500"
              >
                View highlights
              </button>
              <button
                onClick={readwiseConfigured ? syncWithReadwise : promptReadwise}
                type="button"
                className="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-cyan-600 hover:bg-cyan-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500"
              >
                Sync all highlights
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
}
