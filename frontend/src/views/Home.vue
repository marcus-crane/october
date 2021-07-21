<template>
  <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-800 py-12 px-24 grid grid-cols-2 gap-14">
    <div className="space-y-2">
      <img
        className="mx-auto h-36 w-auto logo-animation"
        src="../assets/logo.png"
        alt="The Octowise logo, which is a cartoon octopus reading a book."
      />
      <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
        Octowise
      </h2>
      <p className="mt-0 text-center text-sm text-gray-600 dark:text-gray-400">
        Easily access your Kobo highlights
      </p>
    </div>
    <div className="space-y-4 text-center">
      <h1 class="text-3xl font-bold">Select your Kobo</h1>
      <button @click="detectDevices" class="">Don't see your device? Click here to scan for devices.</button>
      <ul>
        <li v-for="device in devices" :key="device.MntPath" @click="selectDevice(device)">
          <a class="bg-red-200 hover:bg-red-500 hover:shadow-lg group block rounded-lg p-4">
            <dl>
              <div>
                <dt class="sr-only">Title</dt>
                <dd class="border-gray leading-6 font-medium text-black">
                  {{ device.Name }}
                </dd>
                <dt class="sr-only">System Specifications</dt>
                <dd class="text-sm font-normal">
                  {{ device.Storage }} GB · {{ device.DisplayPPI }} PPI
                </dd>
              </div>
            </dl>
          </a>
        </li>
        <p>Local path: {{ localPath }}</p>
        <li>
          <a @click="loadDatabaseFile" class="bg-red-200 hover:bg-red-500 hover:shadow-lg group block rounded-lg p-4">
            <dl>
              <div>
                <dt class="sr-only">Title</dt>
                <dd class="border-gray leading-6 font-medium text-black">
                  Pick a locally available Kobo database (KoboReader.sqlite)
                </dd>
              </div>
            </dl>
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      devices: [],
      localPath: ""
    }
  },
  methods: {
    detectDevices() {
      window.backend.Kobo
        .DetectKobo()
        .then(res => {
          console.log(res)
          this.devices = res.Kobos
        })
        .catch(err => console.log(err))
    },
    loadDatabaseFile() {
      window.backend.Database
        .SelectLocalDatabase()
        .then(res => this.localPath = res)
        .catch(err => console.log(err))
    },
    selectDevice(device) {
      window.backend.Kobo
        .SelectKobo(device.MntPath)
        .then(res => {
          if (res) {
            this.$router.push('overview')
          }
        })
        .catch(err => console.log(err))
    }
  },
  mounted: function() {
    this.$nextTick(this.detectDevices)
  }
};
</script>