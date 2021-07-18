<template>
  <div className="min-h-screen p-16 bg-gray-100 dark:bg-gray-800">
    <h1 class="text-3xl font-bold">{{ device.Name }}</h1>
    <p>Mounted at {{ device.MntPath }}</p>
    <br />
    <p>{{ highlightCount }} highlights on device</p>
    <br />
    <button class="p-3 px-4 rounded-md border text-white font-bold text-lg bg-purple-600" @click="getRandomHighlight">Fetch a new highlight</button>
    <br />
    <br />
    <blockquote class="max-w-4xl text-lg sm:text-2xl font-medium sm:leading-10 space-y-6 mb-6">
      {{ highlight.Text }}
      <p class="pt-8 text-md font-md">— {{ highlight.VolumeID }}</p>
    </blockquote>

  </div>
</template>

<script>
export default {
  name: "Overview",
  data() {
    return {
      device: {},
      highlightCount: 0,
      highlight: {}
    }
  },
  methods: {
    fetchDeviceDetails() {
      window.backend
        .getBasicKoboDetails()
        .then(res => this.device = res)
        .catch(err => console.log(err))
    },
    fetchHighlightCount() {
      window.backend
        .getHighlightCount()
        .then(res => this.highlightCount = res)
        .catch(err => console.log(err))
    },
    getRandomHighlight() {
      window.backend
        .getMostRecentHighlight()
        .then(res => {
          this.highlight = res
          console.log(res)
        })
        .catch(err => console.log(err))
    }
  },
  mounted: function() {
    this.$nextTick(this.fetchDeviceDetails)
    this.$nextTick(this.fetchHighlightCount)
    this.$nextTick(this.getRandomHighlight)
  }
};
</script>