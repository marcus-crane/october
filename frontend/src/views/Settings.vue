<template>
  <div className="min-h-screen p-16 bg-gray-100 dark:bg-gray-800">
    <h1 class="text-3xl font-bold">Settings</h1>
    <button class="p-3 px-4 rounded-md border text-white font-bold text-lg bg-purple-600" @click="navigateToOverview">Overview</button>
  </div>
</template>

<script>
export default {
  name: "Overview",
  methods: {
    navigateToOverview() {
      this.$router.push('overview')
    }
  }
};
</script>