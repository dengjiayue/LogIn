<template>
    <div>
      <h3>验证码登录1</h3>
    </div>
    <div>
        <VerifyPic @verificationImage="handleVerificationImage"></VerifyPic>
      <form @submit="submitForm">
        <label for="piccode">验证码：</label>
        <input type="text" id="piccode" v-model="piccode" pattern="[0-9]{4}" required>
        <br>
        <label for="email">邮箱：</label>
        <input type="email" id="email" v-model="email" required>
        <br>
        <button type="submit">提交</button>
      </form>
    </div>
    <div>
    <button @click="goToSignup">注册</button>
  </div>
  <div>
    <GitHubLoginButton />
    <button @click="goToLogin">使用密码登录</button>
  </div>
  </template>
  
  <script>
  import VerifyPic from './VerifyPic.vue';
  import axios from 'axios';
  import GitHubLoginButton from './components/GitHubLoginButton.vue';
  
  export default {
    components: {
      VerifyPic,
      GitHubLoginButton,
    },
    data() {
      return {
        picid: '',
        piccode: '',
        email: ''
      };
    },
    methods: {
      goToLogin() {
        this.$router.push('/LogInPsw');
      },
      goToSignup() {
        this.$router.push('/signup1');
      },
    handleVerificationImage(picid) {
      this.picid = picid;
    },
      async submitForm(event) {
        event.preventDefault();
        if (!this.validatePiccode(this.piccode)) {
          alert('验证码必须为4位数字');
          return;
        }
        if (!this.validateEmail(this.email)) {
          alert('邮箱格式不正确');
          return;
        }
        try {
          const requestData = {
            picid: this.picid,
            piccode: this.piccode.split('').map(Number),
            email: this.email
          };
          const response = await axios.post('http://localhost:8888/loginemail1', requestData);
          if (response.data.code === 200) {
            // 登录成功后进行页面跳转到 LoginEmail2.vue
            this.$router.push({
        path: '/LogInEmail2',
        query: { email: this.email },
      });
          } else {
            alert(response.data.msg);
          }
        } catch (error) {
          alert('提交失败', error);
        }
      },
      validatePiccode(piccode) {
        return piccode.length === 4 && /^\d+$/.test(piccode);
      },
      validateEmail(email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
      }
    }
  };
  </script>
  