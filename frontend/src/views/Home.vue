<template>
  <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-800 py-12 px-4 sm:px-6 lg:px-8">
    <div className="max-w-md w-full space-y-8">
      <img
        className="mx-auto h-36 w-auto logo-animation"
        src="../assets/logo.png"
        alt="The Octowise logo, which is a cartoon octopus reading a book."
      />
      <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
        Octowise
      </h2>
      <p className="mt-2 text-center text-sm text-gray-600 dark:text-gray-400">
        Easily access your Kobo highlights
      </p>
      <button
        @click="detectDevices"
        to="/about"
        className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 dark:bg-indigo-500 hover:bg-indigo-700 dark:bg-indigo-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Select your connected Kobo device
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: "Home",
  setup() {
    const detectDevices = () => {
      window.backend
        .detectKobo()
        .then(res => {
          if (!res.Name) {
            console.log('No Kobo detected :(')
          } else {
            console.log('Kobo detected:', res)
          }
        })
        .catch(err => console.log(err))
    }

    return {
      detectDevices
    }
  }
};
</script>