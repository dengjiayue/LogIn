<template>
    <div>
      <h1>密码改密</h1>
      <form @submit="changePassword">
        <label for="password">当前密码:</label>
        <input type="password" id="password" v-model="password" required>
        <br>
        <label for="newpassword">新密码:</label>
        <input type="password" id="newpassword" v-model="newPassword" required>
        <br>
        <button type="submit">改密</button>
      </form>
      <button @click="goBack">返回</button>
      <button @click="goToForgotPassword">忘记密码</button>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        password: '',
        newPassword: ''
      };
    },
    methods: {
      goBack() {
        this.$router.push('/HomeV');
      },
      goToForgotPassword() {
        this.$router.push('/RePasswordEmail1');
      },
      async changePassword(event) {
        event.preventDefault(); // 阻止表单提交的默认行为
  
        try {
          const idcode = this.$cookies.get('idcode');
          const requestData = {
            password: this.password,
            newpassword: this.newPassword,
            idcode: idcode
          };
          const response = await axios.post('http://localhost:8888/home/resetpaswbypasw', requestData);
          if (response.data.code === 200) {
            alert('改密成功');
          }else if (response.data.code===420) {
            alert(response.data.msg);
            this.$router.push('/');
          } else {
            alert(response.data.msg);
          }
        } catch (error) {
          console.error('改密失败', error);
          if (error.response && error.response.status === 420) {
          alert('非法访问');
          this.$router.push('/');
        }
        }
      }
    }
  };
  </script>
  