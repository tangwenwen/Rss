<template>
  <div class ="pagination-item">
    <div class = "aa">
      <el-row :gutter="20">
        <div v-for="(value,index) in info.slice(0,pagesize)">
          <el-col :span="6">                  
            <el-card :body-style="{ padding: '0px' }">
              <img src="http://127.0.0.1:8081/src/assets/images/touxiang.png" class="image">
              <div style="padding: 14px;font-size: 13px;line-height: 1.2;">
                <p>名称：{{value.title}}</p>
                <p>分类：{{value.genres}}</p>
                <p>评价人数：{{value.nopa}}</p>
              </div>
              <div class="block">
                <Rate show-text  @on-change="submitRating($event,value.id)"/>
              </div>
            </el-card>
          <!-- <h1>{{value.id}} ,{{value.img}},{{value.title}},{{value.genres}},{{value.Nopa}}</h1> -->

          </el-col>
        </div>
      </el-row>
      <el-pagination
      layout="prev, pager, next"
      :page-size="pagesize"    
       @current-change="current_change" 
       :current-page.sync="currentPage" 
       :pager-count="5" 
       :total="listTotal" style="text-align:center">  
     </el-pagination> 
   </div>
  </div>
</template>
 
 
<script>
import { mapActions } from 'vuex'

export default {
    name: 'PaginationItem',
     data(){
         return{
             info:[],
             listTotal : 0,
             pagesize:20,//每页多少数据
             currentPage:1  //默认当前页为第一页
         }
     },
     created(){
      this.initMovieList();
     },
     methods:{
       ...mapActions([
          'getAllmovie',
          'commitRating'
          ]),
         current_change(currentPage){  //改变当前页
          var pageSize = this.pagesize
          var pageNo = this.currentPage
          this.currentPage = currentPage
          this.getAllmovie({ pageNo, pageSize }).then(res => {
              const data = res.data
              this.info = data.movieList
              this.listTotal = data.pagination.total


          })
         },
         submitRating(item,movieId) {
          console.log(this)
          var rating = item
          this.commitRating({ movieId, rating }).then(res => {
              const data = res
              if (data.code==1){
                this.$message.error(data.error);
              }else{
                this.$message({
                  message: '评价成功',
                  type: 'success'
                });
              }
          })
          },
          initMovieList:function(){
            var pageNo = this.currentPage
            var pageSize = this.pagesize
              this.getAllmovie({ pageNo, pageSize }).then(res => {
              const data = res.data
              this.info = data.movieList
              this.listTotal = data.pagination.total
              this.currentPage = 1
          })
          }
    }
}
</script>
  
<style lang="less">
@import './pagination-item.less';
</style>


