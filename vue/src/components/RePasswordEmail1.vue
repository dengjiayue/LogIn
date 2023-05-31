<template>
    <div>
      <h1>验证码改密1</h1>
      <VerifyPic @verificationImage="handleVerificationImage"></VerifyPic>
      <form @submit="submitForm">
        <label for="piccode">验证码：</label>
        <input type="text" id="piccode" v-model="piccode" pattern="[0-9]{4}" required>
        <br>
        <button type="submit">提交</button>
      </form>
      <button @click="goBack">返回</button>
      <button @click="goToRePasswordPsw">使用密码改密</button>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  import VerifyPic from './VerifyPic.vue';

  
  export default {
    components: {
      VerifyPic
    },
    data() {
      return {
        piccode: '',
        picid: ''
      };
    },
    methods: {
      goBack() {
        this.$router.push('/HomeV');
      },
      goToRePasswordPsw() {
        this.$router.push('/RePasswordPsw');
      },
      handleVerificationImage(picid) {
      this.picid = picid;
    },
    validatePiccode(piccode) {
        return piccode.length === 4 && /^\d+$/.test(piccode);
      },
      async submitForm(event) {
        event.preventDefault(); // 阻止表单提交的默认行为
        if (!this.validatePiccode(this.piccode)) {
          alert('验证码必须为4位数字');
          return;
        }
        try {
          const idcode = this.$cookies.get('idcode');
          const requestData = {
            picid: this.picid, // 请替换为您自己的数据
            piccode: this.piccode.split('').map(Number),
            idcode: idcode
          };
          const response = await axios.post('http://localhost:8888/home/resetpaswbyemail1', requestData);
          if (response.data.code === 200) {
            this.$router.push('/RePasswordEmail2');
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
  