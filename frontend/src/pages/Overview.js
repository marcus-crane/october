import React, { useState } from 'react';

export default function Overview() {
  const [readwiseConfigured, setReadwiseConfigured] = useState(true)
  const [selectedKobo, setSelectedKobo] = useState({})

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-800 py-12 px-24 grid grid-cols-2 gap-14">
      <div className="space-y-2">
        <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
          Overview
        </h2>
      </div>
    </div>
  )
}
