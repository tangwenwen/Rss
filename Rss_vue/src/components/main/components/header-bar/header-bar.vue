<template>
  <div class="header-bar">
    <custom-bread-crumb show-icon style="margin-left: 30px;" :list="breadCrumbList"></custom-bread-crumb>
    <Button @click="value1 = true"  style="margin-left: 16px;">
  点我推荐
</Button>

    <Drawer title="推荐表" :closable="false" v-model="value1" width = "500px">
<Table stripe :columns="columns" :data="gridData"></Table>
    </Drawer>

    <div class="custom-content-con">
      <slot></slot>
    </div>
  </div>
</template>
<script>
import siderTrigger from './sider-trigger'
import customBreadCrumb from './custom-bread-crumb'
import './header-bar.less'
import { mapActions } from 'vuex'
export default {
  name: 'HeaderBar',
  components: {
    siderTrigger,
    customBreadCrumb
  },
  data() {
    return {
      value1: false,
      columns: [
                    {
                        title: 'Id',
                        key: 'id'
                    },
                    {
                        title: '标题',
                        key: 'title'
                    },
                    {
                        title: '分类',
                        key: 'genres'
                    },
                    {
                        title: '评价人数',
                        key: 'nopa'
                    },
                    {
                        title: '封面',
                        key: 'img'
                    }
                ],
      gridData:[]
    };
  },

  props: {
    collapsed: Boolean
  },
  computed: {
    breadCrumbList () {
      return this.$store.state.app.breadCrumbList
    }
  },
  created(){
    this.initTable();
  },
  methods: {
    ...mapActions([
      'getUserRecommend'
    ]),
    handleCollpasedChange (state) {
      this.$emit('on-coll-change', state)
    },
    initTable(){
      this.getUserRecommend().then(res => {
          const data = res.data
          if (data.code==1){
            this.$message.error(data.error);
          }
          gridData = data
      })
    }
  }
}
</script>
