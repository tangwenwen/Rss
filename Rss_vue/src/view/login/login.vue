<style lang="less">
  @import './login.less';
</style>

<template>
  <div class="login">
    <div class="login-con">
      <Card icon="log-in" title="欢迎登录" :bordered="false">
        <div class="form-con">
          <login-form @on-success-valid="handleSubmit"></login-form>
        </div>
        <div class="form-reg">
          <el-button class="bu" type="text" @click="dialogFormVisible = true">注册</el-button>

          
        </div>
          <p class="login-tip">
            <Alert type="error"
                   v-show="flag"
                   show-icon>
              {{message}}
            </Alert>
          </p>
      </Card>

    </div>
      <el-dialog title="注册" :visible.sync="dialogFormVisible" center width="35%" @close='closeDialog' @open='openDialog'>
              <el-form :model="form">
                <el-form-item label="用户名" :label-width="formLabelWidth">
                  <el-input v-model="form.username" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="密码" :label-width="formLabelWidth">
                  <el-input v-model="form.password" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="年龄" :label-width="formLabelWidth">
                  <el-input v-model="form.age" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="性别" :label-width="formLabelWidth">
                  <el-switch
                    v-model="form.gender"
                    active-text="男"
                    inactive-text="女">
                  </el-switch>
                </el-form-item>
                <el-form-item label="职业" :label-width="formLabelWidth">
                    <el-radio-group v-model="form.occupation">
                      <el-radio :label="0">其他职业</el-radio>      
                      <el-radio :label="1">学者/教育家</el-radio>
                      <el-radio :label="2">艺术家</el-radio>
                      <el-radio :label="3">文书/管理</el-radio>
                      <el-radio :label="4">大学生/研究生</el-radio>
                      <el-radio :label="5">客户服务</el-radio>
                      <el-radio :label="6">医生/保健</el-radio>
                      <el-radio :label="7">行政/管理</el-radio>
                      <el-radio :label="8">农民</el-radio>
                      <el-radio :label="9">家庭主妇</el-radio>
                      <el-radio :label="10">低龄学生</el-radio>
                      <el-radio :label="11">律师</el-radio>
                      <el-radio :label="12">程序员</el-radio>
                      <el-radio :label="13">退休</el-radio>
                      <el-radio :label="14">销售/营销</el-radio>
                      <el-radio :label="15">科学家</el-radio>
                      <el-radio :label="16">个体</el-radio>
                      <el-radio :label="17">技术员/工程师</el-radio>
                      <el-radio :label="18">商人/工匠</el-radio>
                      <el-radio :label="19">失业</el-radio>
                      <el-radio :label="20">作家</el-radio>
                    </el-radio-group> 
                </el-form-item>
              </el-form>
              <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="dialogFormVisible = false; submitReg(form);" >确 定</el-button>
              </div>
            </el-dialog>

  </div>
</template>


<script>
import LoginForm from '_c/login-form'
import { mapActions } from 'vuex'
export default {
  components: {
    LoginForm
  },
  data(){
    return{
      occupation: 0,
      gender :false,
      message :'',
      flag:false,
        dialogFormVisible: false,
        form: {
          name: '',
          region: '',
          date1: '',
          date2: '',
          delivery: false,
          type: [],
          resource: '',
          desc: ''
        },
        formLabelWidth: '130px'     

    }
  },
  methods: {
    ...mapActions([
      'handleLogin',
      'getUserInfo',
      'handleAddUser'
    ]),
    handleSubmit ({ userName, password }) {

      this.handleLogin({ userName, password }).then(res => {
        const data = res
        if (data.data.code==0){
          this.getUserInfo().then(res => {
            this.$router.push({
              name: this.$config.homeName
          })
        })
        }else{
          this.message = data.data.error// 获取后始返回消息显示
          this.flag = true
        } 
      })
    },
    submitReg(form){
      var userName = form.username
      var passWord = form.password
      var age = form.age
      var gender = ""
      if (form.age){
        gender = "F"
      }else{
        gender = "M"
      }
      var occupation = form.occupation+""
      this.handleAddUser({ userName, passWord, age, gender, occupation }).then(res => {
          const data = res.data
          if (data.code==1){
            this.$message.error(data.error);
          }
      })
    },
    closeDialog(){
      this.form.username = '';
      this.form.age = '';
      this.form.password = '';
      this.form.gender = false
      this.form.occupation = 0

    },
    openDialog(){
      this.form.gender = false
      this.form.occupation = 0
    },

  }
}
</script>
