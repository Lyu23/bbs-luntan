<template>
  <div class="topics-main">
    <div class="top-nav">
      <a
        v-for="(game, index) in articlesGamesjson"
        :key="index"
        class="gameto"
        :href="'/topics/tag/' + tagId + '/' + game"
      >
        {{ index }}
      </a>
    </div>
    <load-more
      v-slot="{ results }"
      :init-data="articlesPage"
      :params="{ tagId: tag.tagId }"
      url="/api/article/tag/articles"
    >
      <article-list :articles="results" />
    </load-more>
  </div>
</template>

<script>
export default {
  async asyncData({ $axios, params }) {
    try {
      const [tag, articlesPage, articlesGames] = await Promise.all([
        $axios.get('/api/tag/' + params.tagId),
        $axios.get('/api/article/tag/articles', {
          params: {
            tagId: params.tagId,
          },
        }),
        $axios.get('api/article/game/' + params.tagId),
      ])
      const articlesGamesjson = JSON.parse(articlesGames.ArticleGame)
      const tagId = params.tagId
      return {
        tagId,
        tag,
        articlesPage,
        articlesGamesjson,
      }
    } catch (e) {
      console.error(e)
    }
  },
  head() {
    return {
      title: this.$siteTitle(this.tag.tagName + ' - 文章'),
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
.gameto {
  width: calc(33.33% - 10px);
  display: inline-block;
  height: 50px;
  text-align: center;
  text-decoration: none;
  background-color: #87ceeb;
  color: block;
  border-radius: 10%;
  margin: 0 5px 10px;

  text-align: center;
  line-height: 50px;
  font-family: 'Times New Roman';
  vertical-align: middle;
  font-size: 23px;
  font-weight: bold;
  text-decoration: none;
  color: #333;
}
.gameto:hover {
  background-color: #0062cc;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.5);
}

.top-nav {
  flex-wrap: wrap;
  border: 5px solid #ccc;
  display: flex;
  justify-content: center;
  align-items: center;

  padding: 50px 100px;
  background-color: #ffffe0;
  color: #333;
  text-decoration: none;
  font-weight: bold;
  border-radius: 10px;
  margin-right: 10px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
}
</style>
