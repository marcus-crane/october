<template>
  <div className="min-h-screen p-16 bg-gray-100 dark:bg-gray-800">
    <h1 class="text-3xl font-bold">{{ device.Name }}</h1>
    <p>Mounted at {{ device.MntPath }}</p>
    <p>{{ highlightCount }} highlights on device</p>
    <button class="p-2 rounded-md border text-white font-normal text-sm bg-purple-600" @click="navigateToSettings">Settings</button>
    <br />
    <br />
    <ul>
      <li v-for="item in items" :key="item.BookID">
        <p>{{ item.Title }}</p>
        <progress :value="item.PercentRead" max="100"></progress> {{ item.PercentRead }}%
      </li>
    </ul>

  </div>
</template>

<script>
export default {
  name: "Overview",
  data() {
    return {
      device: {},
      highlightCount: 0,
      items: {}
    }
  },
  methods: {
    navigateToSettings() {
      this.$router.push('settings')
    },
    fetchDeviceDetails() {
      window.backend.Kobo
        .GetBasicKoboDetails()
        .then(res => this.device = res)
        .catch(err => console.log(err))
    },
    fetchHighlightCount() {
      window.backend.Bookmark
        .GetHighlightCount()
        .then(res => this.highlightCount = res)
        .catch(err => console.log(err))
    },
    getAllItems() {
      window.backend.Content
        .GetAllItems()
        .then(res => {
          this.items = res
          console.log(res)
        })
    }
  },
  mounted: function() {
    this.$nextTick(this.fetchDeviceDetails)
    this.$nextTick(this.fetchHighlightCount)
    this.$nextTick(this.getRandomHighlight)
    this.$nextTick(this.getAllItems)
  }
};
</script>