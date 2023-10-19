<template>
  <section class="main">
    <div class="container main-container left-main size-320">
      <div class="left-container">
        <div class="main-content no-padding no-bg topics-wrapper">
          <!-- <div class="topics-nav"><topics-nav :nodes="nodes" /></div> -->
          <div class="topics-main">
            <div style="display: flex">
              <div class="news">
                <a href="/article/3">
                  <img style="width: 100%; height: 280px" src="./2.png" />
                </a>
                <div style="margin-left: 20px">
                  <a href="/user/2">liam</a>
                  <p class="text">以太坊的上海升级</p>
                </div>
              </div>
              <div class="news">
                <a href="/article/3">
                  <img style="width: 100%; height: 280px" src="./3.png" />
                </a>
                <div style="margin-left: 20px">
                  <a href="/user/2">liam</a>
                  <p class="text">Axie</p>
                </div>
              </div>
            </div>
            <div
              style="
                width: 150px;
                height: 80px;
                display: flex;
                justify-content: center;
                align-items: center;
              "
            >
              <p style="margin-top: 20px; font-size: 30px">最新文章</p>
            </div>
            <sticky-topics :node-id="0" />
            <load-more
              v-slot="{ results }"
              :init-data="articlesPage"
              url="/api/article/articles"
            >
              <article-list :articles="results" />
            </load-more>
          </div>
        </div>
      </div>
      <div class="right-container">
        <News :articlest="articlest" />
        <!-- <check-in /> -->
        <site-notice />
        <!-- <score-rank :score-rank="scoreRank" /> -->
        <friend-links :links="links" />
      </div>
    </div>
  </section>
</template>
<script>
export default {
  async asyncData({ $axios, store }) {
    store.commit('env/setCurrentNodeId', 0) // 设置当前所在node
    try {
      const [articlesPage] = await Promise.all([
        $axios.get('/api/article/articles'),
      ])
      const articlest = articlesPage
      const [nodes, topicsPage, scoreRank, links] = await Promise.all([
        $axios.get('/api/topic/node_navs'),
        $axios.get('/api/topic/topics'),
        $axios.get('/api/user/score/rank'),
        $axios.get('/api/link/toplinks'),
      ])
      return { nodes, topicsPage, scoreRank, links, articlesPage, articlest }
    } catch (e) {
      console.error(e)
    }
  },
  data() {},
  head() {
    return {
      title: this.$siteTitle(),
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

<style lang="scss" scoped>
.news {
  display: inline-block;
  margin-left: 10px;
  margin-right: 10px;
  width: 50%;
  background-color: #2c2c41;
}
.text {
  height: 80px;
  font-size: 30px;
}
</style>
