<template>
  <div>
      <h3>密码登录</h3>
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
        <label for="password">密码：</label>
        <input type="password" id="password" v-model="password" required>
        <br>
        <label for="piccode">验证码：</label>
        <input type="text" id="piccode" v-model="piccode" pattern="[0-9]{4}" required>
        <br>
        <button type="submit">提交</button>
      </form>
    </div>
    <div>
    <button @click="goToLoginE">忘记密码</button>
  </div>
    <div>
    <button @click="goToSignup">注册</button>
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
        password: '',
        piccode: ''
      };
    },
    methods: {
      goToLoginE() {
        this.$router.push('/LogInEmail');
      },
      goToSignup() {
        this.$router.push('/signup1');
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
        if (!this.validatePassword(this.password)) {
          console.error('密码长度不能超过20个字符');
          return;
        }
        if (!this.validatePiccode(this.piccode)) {
          console.error('验证码必须为4位数字');
          return;
        }
        try {
          const requestData = {
            picid: this.picid,
            piccode: this.piccode.split('').map(Number),
            email: this.email,
            password: this.password
          };
          const response = await axios.post('http://localhost:8888/loginpsw', requestData);
          if (response.data.code === 200) {
            // 登录成功后进行页面跳转到 HomeV.vue
            this.$cookies.set('idcode', response.data.data.idcode);
            this.$router.push('/HomeV');
          } else {
            console.error(response.data.msg);
          }
        } catch (error) {
          console.error('提交失败', error);
        }
      },
      validateEmail(email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
      },
      validatePassword(password) {
        return password.length <= 20;
      },
      validatePiccode(piccode) {
        return piccode.length === 4 && /^\d+$/.test(piccode);
      }
    }
  };
  </script>
  