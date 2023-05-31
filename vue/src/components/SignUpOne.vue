<template>
    <div>
      <h3>注册1</h3>
    </div>
    <div>
      <button @click="getVerificationImage">获取验证图片</button>
      <div v-if="verificationImage">
      <img :src="decodeBase64(verificationImage.png)" alt="验证码图片">
    </div>
      <form @submit="submitForm">
        <label for="email">邮箱：</label>
        <input type="email" id="email" v-model="email" required>
        <br>
        <label for="piccode">验证码：</label>
        <input type="text" id="piccode" v-model="piccode" pattern="[0-9]{4}" required>
        <br>
        <button type="submit">提交</button>
        <button @click="goToLogin">登录</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        verificationImage: null,
        picid: '',
        email: '',
        piccode: ''
      };
    },
    methods: {
      goToLogin() {
        this.$router.push('/LogInPsw');
      },
      async getVerificationImage() {
        try {
          const response = await axios.get('http://localhost:8888/verifypic');
          if (response.data.code === 200) {
            this.verificationImage = response.data.data;
            this.picid = response.data.data.picid;
            
          }
        } catch (error) {
          console.error('获取验证图片失败', error);
        }
      },
      decodeBase64(base64Data) {
      return `data:image/png;base64, ${base64Data}`;
    },
      async submitForm(event) {
        event.preventDefault();
        if (!this.validateEmail(this.email)) {
          console.error('邮箱格式不正确');
          return;
        }
        if (!this.validatePiccode(this.piccode)) {
          console.error('验证码必须为4位数字');
          return;
        }
        try {
          const requestData = {
            picid: this.picid,
            email: this.email,
            piccode: this.piccode.split('').map(Number)
          };
          const response = await axios.post('http://localhost:8888/signup1', requestData);
          if (response.data.code === 200) {
            // 成功后进行页面跳转到
            this.$router.push({
        path: '/signup-tow',
        query: { email: this.email },
      });
          }
        } catch (error) {
          console.error('提交失败', error);
        }
      },
      validateEmail(email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
      },
      validatePiccode(piccode) {
        return piccode.length === 4 && /^\d+$/.test(piccode);
      }
    }
  };
  </script>
  