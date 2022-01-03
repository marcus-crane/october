import React, { useState, useEffect } from 'react';

export default function Overview(props) {
  console.log(props)
  const [readwiseConfigured, setReadwiseConfigured] = useState(true)
  const [selectedKobo, setSelectedKobo] = useState({})

  useEffect(() => {
    window.go.main.KoboService.GetSelectedKobo()
      .then(kobo => {
        console.log(kobo)
        setSelectedKobo(kobo)
      })
  }, [selectedKobo.mnt_path])

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-800 py-12 px-24 grid grid-cols-2 gap-14">
      <div className="space-y-2">
        <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
          {selectedKobo.name}
        </h2>
        <p className="mt-0 text-center text-sm text-gray-600 dark:text-gray-400">
          {selectedKobo.storage} GB · {selectedKobo.display_ppi} PPI
        </p>
      </div>
    </div>
  )
}
