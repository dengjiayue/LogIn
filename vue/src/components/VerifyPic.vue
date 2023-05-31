<template>
  <div>
    <button @click="getVerificationImage">获取验证码图片</button>
    <div v-if="verificationImage">
      <img :src="decodeBase64(verificationImage.png)" alt="验证码图片">
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      verificationImage: null,
      picid: ''
    };
  },
  methods: {
    async getVerificationImage() {
    try {
      const response = await axios.get('http://localhost:8888/verifypic');
      if (response.data.code === 200) {
        this.verificationImage = response.data.data;
        this.picid = response.data.data.picid;
        // 触发自定义事件，将 picid 传递给父组件
        this.$emit('verificationImage', this.picid);
      } else {
        console.error(response.data.msg);
      }
    } catch (error) {
      console.error('获取验证码图片失败', error);
    }
  },
    decodeBase64(base64Data) {
      return `data:image/png;base64, ${base64Data}`;
    }
  }
};
</script>
