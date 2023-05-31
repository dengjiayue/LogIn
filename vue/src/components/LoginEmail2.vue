<template>
  <div>
      <h3>验证码登录2</h3>
    </div>
    <div>
      <form @submit="submitForm">
        <label for="verification">验证码：</label>
        <input type="number" id="verification" v-model.number="verification" max="999999" required>
        <button @click="goToLogin1">重发验证码</button>
        <br>
        <button type="submit">登录</button>
      </form>
    </div>
    <div>
    <button @click="goToSignup">注册</button>
  </div>
  <div>
    <button @click="goToLogin">使用密码登录</button>
  </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        email: '',
        verification: null
      };
    },
    created() {
      this.email = this.$route.query.email || '';
    },
    methods: {
      goToLogin1() {
        this.$router.push('/LogInEmail');
      },
      goToLogin() {
        this.$router.push('/LogInPsw');
      },
      goToSignup() {
        this.$router.push('/signup1');
      },
      async submitForm(event) {
        event.preventDefault();
        if (!this.validateVerification(this.verification)) {
          alert('验证码必须为小于6位的数字');
          return;
        }
        try {
          const requestData = {
            email: this.email,
            verification: this.verification
          };
          const response = await axios.post('http://localhost:8888/loginemail2', requestData);
          if (response.data.code === 200) {
            // 登录成功后进行页面跳转到 HomeV.vue
            this.$cookies.set('idcode', response.data.data.idcode);
            this.$router.push('/HomeV');
          } else {
            alert(response.data.msg);
          }
        } catch (error) {
          alert('提交失败', error);
        }
      },
      validateVerification(verification) {
        return verification !== null && verification.toString().length <= 6;
      }
    }
  };
  </script>
  