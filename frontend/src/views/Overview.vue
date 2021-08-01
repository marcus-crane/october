<template>
  <div class="p-4 bg-gray-100 dark:bg-gray-800">
    <h1 class="text-3xl font-bold">{{ device.Name }}</h1>
    <p>Mounted at {{ device.MntPath }}</p>
    <p>{{ highlightCount }} highlights on device</p>
    <button class="p-2 rounded-md border text-white font-normal text-sm bg-purple-600" @click="navigateToSettings">Settings</button>
    <br />
    <br />
    <div class="grid grid-cols-9">
      <div class="bg-blue-100 col-span-3 overflow-auto">
        <ul class="space-y-4 p-4">
          <li v-for="item in items" :key="item.ContentID" @click="selectItem(item)">
            <div class="flex items-center">
              <div class="flex-shrink-0 h-10 w-10">
                <img class="h-10 w-10 rounded-full" src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60" alt="">
              </div>
              <div class="ml-4">
                <div class="text-sm font-medium text-gray-900">{{ item.Title }}</div>
                <div class="text-sm text-gray-500">Author goes here</div>
              </div>
            </div>
          </li>
        </ul>
      </div>
      <div class="bg-red-100 col-span-6 overflow-auto p-4">
        <h1 class="text-sm font-medium text-gray-900">{{ selectedItem.Title }}</h1>
        <p class="text-sm text-gray-500">Author goes here</p>
        <p>Last read: {{ selectedItem.DateLastRead }}</p>
        <br />
        <br />
        <progress :value="selectedItem.PercentRead" max="100"></progress> {{ selectedItem.PercentRead }}%
        <br />
        <br />
        <p v-html="selectedItem.Description"></p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Overview",
  data() {
    return {
      device: {},
      highlightCount: 0,
      items: [],
      selectedItem: {}
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
    },
    selectItem(item) {
      this.selectedItem = this.items.find(i => i.ContentID === item.ContentID)
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