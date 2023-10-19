<template>
  <div class="widget">
    <div class="widget-header">最新快讯</div>
    <div class="widget-content checkin">
      <div class="checkedin">
        <div class="gold-icon-box">
          <li v-for="(article, index) in articlest.results" :key="index">
            {{ article.user.nickname }} :{{ article.summary }}
            <p>-----------------------------------</p>
          </li>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  // props: {
  //   checkInRank: {
  //     type: Array,
  //     default() {
  //       return null
  //     },
  //   },
  // },
  props: {
    articlest: {
      type: Array,
      default() {
        return []
      },
      required: false,
    },
  },
  data() {
    return {
      checkIn: null,
      checkInRank: null,
    }
  },
  computed: {
    user() {
      return this.$store.state.user.current
    },
    isLogin() {
      return this.$store.state.user.current != null
    },
  },
  mounted() {
    this.getCheckIn()
    this.loadCheckInRank()
  },
  methods: {
    async getCheckIn() {
      try {
        this.checkIn = await this.$axios.get('/api/checkin/checkin')
      } catch (e) {
        console.log(e)
      }
    },
    async doCheckIn() {
      if (!this.isLogin) {
        this.$toSignin()
      }
      try {
        await this.$axios.post('/api/checkin/checkin')
        this.$message.success('签到成功')
        await this.getCheckIn()
        await this.loadCheckInRank()
      } catch (e) {
        console.error(e)
      }
    },
    async loadCheckInRank() {
      try {
        this.checkInRank = await this.$axios.get('/api/checkin/rank')
      } catch (e) {
        console.error(e)
      }
    },
  },
}
</script>

<style lang="scss" scoped>
.checkin {
  .checkedin {
    display: flex;
    justify-content: space-between;
    align-items: center;
    .gold-icon-box {
      width: 80%;
      span {
        color: var(--text-color4);
        font-size: 0.875rem;
        line-height: 1.25rem;
        display: flex;
        svg {
          width: 1.5rem;
          height: 1.5rem;
        }
      }
    }
    .gold-info-box {
      width: 25%;
      .gold-info {
        color: var(--text-color4);
        font-size: 0.875rem;
        line-height: 1.25rem;
        display: flex;
      }
    }
    .checkedin-btn-box {
      display: flex;
      flex-direction: column;
      width: 50%;
      .checkedin-btn {
        display: flex;
        outline: none;
        align-items: center;
        justify-content: center;
        padding: 6px 0.5rem;
        font-size: 0.75rem;
        line-height: 1rem;
        font-weight: 500;
        white-space: nowrap;
        border-radius: 0.25rem;
        background-color: var(--bg-color);
        border: 1px solid var(--border-color2);
        color: var(--text-color3);
        &:hover {
          background-color: var(--bg-color2);
          border-color: var(--border-color2);
        }
        .checkedin-btn-icon {
          margin-right: 0.5rem;
          display: flex;
          svg {
            width: 1rem;
            height: 1rem;
          }
        }
      }
    }
  }
  .checkedin-tips-box {
    display: flex;
    flex-direction: row;
    justify-content: center;
    text-align: center;
    align-items: center;
    color: var(--text-color3);
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
    margin-top: 0.5rem;
    margin-bottom: 0.5rem;
    border-radius: 0.25rem;
    background: var(--bg-color2);
    .checkedin-tips-icon {
      margin-right: 0.5rem;
      display: flex;
      svg {
        width: 1.25rem;
        height: 1.25rem;
      }
    }
    .checkedin-tips-info {
      font-size: 0.875rem;
      line-height: 1.25rem;
      display: flex;
      .checkedin-tips-day {
        color: #00a1d6; // TODO
      }
    }
  }
  .rank {
    border-top: 1px solid var(--border-color);
    margin-top: 10px;
    padding-top: 10px;
    .rank-title {
      font-size: 14px;
      font-weight: 600;
    }
    li {
      display: flex;
      list-style: none;
      margin: 8px 0;
      font-size: 13px;
      position: relative;
      background: var(--bg-color);
      padding: 0.5rem;
      border-radius: 0.25rem;
      cursor: pointer;

      &:hover {
        background: var(--bg-color2);
      }

      &:not(:last-child) {
        border-bottom: 1px solid var(--border-color);
      }

      .rank-user-avatar {
        min-width: 30px;
        margin-top: 0.5rem;
      }

      .rank-user-info {
        width: 100%;
        margin-left: 10px;
        line-height: 1.4;
        font-size: 12px;
        a {
          color: var(--text-color2);
          font-weight: 700;
          &:hover {
            color: var(--text-color1);
            text-decoration: none;
          }
        }
        p {
          margin-top: 0.5rem;
        }
      }
    }
  }
}
</style>
