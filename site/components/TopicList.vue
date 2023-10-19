<template>
  <ul class="topic-list">
    <li v-for="topic in topics" :key="topic.topicId" class="topic-item">
      <div
        class="topic-avatar"
        :href="'/user/' + topic.user.id"
        :title="topic.user.nickname"
      >
        <avatar :user="topic.user" />
      </div>
      <div class="topic-main-content">
        <div class="topic-top">
          <div class="topic-userinfo">
            <avatar class="topic-inline-avatar" :user="topic.user" size="20" />
            <nuxt-link :to="'/user/' + topic.user.id">{{
              topic.user.nickname
            }}</nuxt-link>
            <span v-if="showSticky && topic.sticky" class="topic-sticky-icon"
              >置顶</span
            >
          </div>
          <div class="topic-time">
            发布于{{ topic.createTime | prettyDate }}
          </div>
        </div>
        <div class="topic-content" :class="{ 'topic-tweet': topic.type === 1 }">
          <template v-if="topic.type === 0">
            <h1 class="topic-title">
              <nuxt-link :to="'/topic/' + topic.topicId">{{
                topic.title
              }}</nuxt-link>
            </h1>
            <nuxt-link :to="'/topic/' + topic.topicId" class="topic-summary">{{
              topic.summary
            }}</nuxt-link>
          </template>
          <template v-if="topic.type === 1">
            <nuxt-link
              v-if="topic.content"
              :to="'/topic/' + topic.topicId"
              class="topic-summary"
              >{{ topic.content }}</nuxt-link
            >
            <ul
              v-if="topic.imageList && topic.imageList.length"
              class="topic-image-list"
            >
              <li v-for="(image, index) in topic.imageList" :key="index">
                <nuxt-link :to="'/topic/' + topic.topicId" class="image-item">
                  <img v-lazy="image.preview" />
                </nuxt-link>
              </li>
            </ul>
          </template>
        </div>
        <div class="topic-bottom">
          <div class="topic-handlers">
            <div
              class="btn"
              :class="{ liked: topic.liked }"
              @click="like(topic)"
            >
              <i class="iconfont icon-like" />{{ topic.liked ? '已赞' : '赞' }}
              <span v-if="topic.likeCount > 0">{{ topic.likeCount }}</span>
            </div>
            <p>{{ likeUsers }}123</p>
            <!-- 点赞用户列表 -->
            <div v-if="likeUsers && likeUsers.length" class="topic-like-users">
              <avatar
                v-for="likeUser in likeUsers"
                :key="likeUser.id"
                :user="likeUser"
                :round="true"
                :has-border="true"
                size="24"
              />
              <span class="like-count">{{ topic.likeCount }}</span>
            </div>
            <div class="btn" @click="toTopicDetail(topic.topicId)">
              <i class="iconfont icon-comment" />评论
              <span v-if="topic.commentCount > 0">{{
                topic.commentCount
              }}</span>
            </div>
            <div class="btn" @click="toTopicDetail(topic.topicId)">
              <i class="iconfont icon-read" />浏览
              <span v-if="topic.viewCount > 0">{{ topic.viewCount }}</span>
            </div>
          </div>
          <div class="topic-tags">
            <nuxt-link
              v-if="topic.node"
              class="topic-tag"
              :to="'/topics/node/' + topic.node.nodeId"
              :alt="topic.node.name"
              >{{ topic.node.name }}</nuxt-link
            >
          </div>
        </div>
      </div>
    </li>
  </ul>
</template>

<script>
export default {
  props: {
    topics: {
      type: Array,
      default() {
        return []
      },
      required: false,
    },
    showAvatar: {
      type: Boolean,
      default: true,
    },
    showSticky: {
      type: Boolean,
      default: false,
    },
  },
  async asyncData({ $axios, params, error }) {
    let topic
    try {
      topic = await $axios.get('/api/topic/' + params.id)
    } catch (e) {
      error({
        message: e.message,
      })
      return
    }

    const [liked, commentsPage, likeUsers] = await Promise.all([
      $axios.get('/api/like/liked', {
        params: {
          entityType: 'topic',
          entityId: params.id,
        },
      }),
      $axios.get('/api/comment/comments', {
        params: {
          entityType: 'topic',
          entityId: params.id,
        },
      }),
      $axios.get('/api/topic/recentlikes/4'),
    ])

    return {
      topic,
      commentsPage,
      liked: liked.liked,
      likeUsers,
    }
  },
  data() {
    return {
      commentsPage: null,
      liked: null,
      likeUsers: null,
    }
  },
  methods: {
    async asyncData({ $axios, params }) {
      const [liked, commentsPage, likeUsers] = await Promise.all([
        $axios.get('/api/like/liked', {
          params: {
            entityType: 'topic',
            entityId: 4,
          },
        }),
        $axios.get('/api/comment/comments', {
          params: {
            entityType: 'topic',
            entityId: 4,
          },
        }),
        $axios.get('/api/topic/recentlikes/' + 4),
      ])
      console.log('>>>>>>>>>>>>>>')
      return {
        commentsPage,
        liked: liked.liked,
        likeUsers,
      }
    },
    async like(topic) {
      try {
        await this.$axios.post('/api/topic/like/' + topic.topicId)
        topic.liked = true
        topic.likeCount++
        this.$message.success('点赞成功')
      } catch (e) {
        if (e.errorCode === 1) {
          this.$msgSignIn()
        } else {
          this.$message.error(e.message || e)
        }
      }
    },
    toTopicDetail(topicId) {
      this.$linkTo(`/topic/${topicId}`)
    },
  },
}
</script>

<style lang="scss" scoped></style>
