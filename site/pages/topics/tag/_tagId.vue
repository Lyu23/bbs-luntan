<template>
  <div class="topics-main">
    <sticky-topics :node-id="0" />

    <div
      style="border: 5px solid #ccc; padding: 30px; background-color: #87ceeb"
    >
      <p>链活跃人数 : 100</p>
      <p>币： 12$</p>
      <p>消息数： 1000</p>
    </div>
    <div class="navbar-item">
      <create-topic-btn :article-id="articleId" :topic-id="tagId" />
    </div>
    <p>{{ articleId }}{{ tagId }}</p>
    <load-more
      v-if="topicsPage"
      v-slot="{ results }"
      :init-data="topicsPage"
      :url="'/api/topic/tag/topics' + tag.tagId"
    >
      <topic-list :topics="results" />
    </load-more>
  </div>
</template>

<script>
export default {
  async asyncData({ $axios, params, store }) {
    store.commit('env/setCurrentNodeId', 0) // 设置当前所在node
    const [tag, topicsPage] = await Promise.all([
      $axios.get('/api/tag/' + params.tagId),
      $axios.get('/api/topic/tag/topics', {
        params: {
          tagId: params.tagId,
        },
      }),
    ])
    const articleId = params.articleId
    const tagId = params.tagId
    return {
      tag,
      topicsPage,
      articleId,
      tagId,
    }
  },
  head() {
    return {
      title: this.$siteTitle(this.tag.tagName + ' - 话题'),
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
