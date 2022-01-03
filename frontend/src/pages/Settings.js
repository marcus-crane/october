import React, { useState } from 'react';
import Navbar from "../Components/Navbar";

export default function Settings() {
  const [settings, setSettings] = useState({})

  return (
    <div className="bg-gray-100 dark:bg-gray-800 ">
      <Navbar />
      <div className="min-h-screen flex items-center justify-center pb-24 px-24 grid grid-cols-2 gap-14">
        <div className="space-y-2">
          <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
            Settings
          </h2>
        </div>
      </div>
    </div>
  )
}
