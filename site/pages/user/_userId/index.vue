<template>
  <section class="main">
    <div class="container">
      <user-profile :user="user" />

      <div class="container main-container right-main size-320">
        <user-center-sidebar :user="user" />
        <div class="right-container">
          <div class="tabs-warp">
            <div class="tabs">
              <ul>
                <li :class="{ 'is-active': activeTab === 'topics' }">
                  <nuxt-link :to="'/user/' + user.id">
                    <span class="icon is-small">
                      <i class="iconfont icon-topic" aria-hidden="true" />
                    </span>
                    <span>话题</span>
                  </nuxt-link>
                </li>
                <li :class="{ 'is-active': activeTab === 'articles' }">
                  <nuxt-link :to="'/user/' + user.id + '/articles'">
                    <span class="icon is-small">
                      <i class="iconfont icon-article" aria-hidden="true" />
                    </span>
                    <span>文章</span>
                  </nuxt-link>
                </li>
              </ul>
            </div>

            <div>
              <div
                v-if="
                  topicsPage && topicsPage.results && topicsPage.results.length
                "
              >
                <load-more
                  v-if="topicsPage"
                  v-slot="{ results }"
                  :init-data="topicsPage"
                  :url="'/api/topic/user/topics?userId=' + user.id"
                >
                  <topic-list :topics="results" :show-avatar="false" />
                </load-more>
              </div>
              <div v-else class="notification is-primary">暂无话题</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  async asyncData({ $axios, params, query, error }) {
    let user
    try {
      user = await $axios.get('/api/user/' + params.userId)
    } catch (err) {
      error({
        statusCode: 404,
        message: err.message || '系统错误',
      })
      return
    }
    const topicsPage = await $axios.get('/api/topic/user/topics', {
      params: { userId: params.userId },
    })
    return {
      activeTab: 'topics',
      user,
      topicsPage,
    }
  },
  data() {
    return {}
  },
  head() {
    return {
      title: this.$siteTitle(this.user.nickname + ' - 话题'),
    }
  },
  computed: {
    currentUser() {
      return this.$store.state.user.current
    },
    isOwner() {
      const current = this.$store.state.user.current
      return this.user && current && this.user.id === current.id
    },
  },
}
</script>

<style lang="scss" scoped>
.tabs-warp {
  background-color: var(--bg-color);
  padding: 0 10px 10px;

  .tabs {
    margin-bottom: 5px;
  }

  .more {
    text-align: right;
  }
}
</style>
