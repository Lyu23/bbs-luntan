<template>
  <div class="topics-main">
    <sticky-topics :node-id="0" />
    <load-more
      v-if="topicsPage"
      v-slot="{ results }"
      :init-data="topicsPage"
      url="/api/topic/topics"
    >
      <p>123</p>
      <topic-list :topics="results" />
    </load-more>
  </div>
</template>

<script>
export default {
  async asyncData({ $axios, store }) {
    try {
      store.commit('env/setCurrentNodeId', 0) // 设置当前所在node
      const [topicsPage] = await Promise.all([$axios.get('/api/topic/topics')])
      return { topicsPage }
    } catch (e) {
      console.error(e)
    }
  },
  head() {
    return {
      title: this.$siteTitle('话题'),
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: this.$siteDescription(),
        },
        { hid: 'keywords', name: 'keywords', content: this.$siteKeywords() },
      ],
    }
  },
}
</script>

<style lang="scss" scoped></style>
