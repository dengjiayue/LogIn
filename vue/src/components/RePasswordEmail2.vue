<template>
    <div>
      <h1>验证码改密2</h1>
      <form @submit="submitForm">
        <label for="newpassword">新密码:</label>
        <input type="password" id="newpassword" v-model="newPassword" required>
        <br>
        <label for="verification">验证码:</label>
        <input type="text" id="verification" v-model="verification" required>
      <button @click="ReVerify">重发验证码</button>
        <br>
        <button type="submit">提交</button>
      </form>
      <button @click="goBack">返回</button>
      <button @click="goToRePasswordPsw">使用密码改密</button>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        newPassword: '',
        verification: ''
      };
    },
    methods: {
      ReVerify() {
        this.$router.push('/RePasswordEmail1');
      },
      goBack() {
        this.$router.push('/HomeV');
      },
      goToRePasswordPsw() {
        this.$router.push('/RePasswordPsw');
      },
      async submitForm(event) {
        event.preventDefault(); // 阻止表单提交的默认行为
  
        try {
          const idcode = this.$cookies.get('idcode');
          const requestData = {
            newpassword: this.newPassword,
            idcode: idcode,
            verification: parseInt(this.verification)
          };
          const response = await axios.post('http://localhost:8888/home/resetpaswbyemail2', requestData);
          if (response.data.code === 200) {
            alert('改密成功');
          }else if (response.data.code===420) {
            alert(response.data.msg);
            this.$router.push('/');
          } else {
            alert(response.data.msg);
          }
        } catch (error) {
          console.error('提交失败', error);
          if (error.response && error.response.status === 420) {
          alert('非法访问');
          this.$router.push('/');
        }
        }
      }
    }
  };
  </script>
  