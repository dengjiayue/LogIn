<template>
    <button @click="login">GitHub授权登录</button>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    methods: {
      async login() {
        try {
          const authorizationUrl = 'https://github.com/login/oauth/authorize';
          const clientId = 'YOUR_CLIENT_ID';
          const redirectUri = 'YOUR_REDIRECT_URI';
  
          // 发起授权请求，获取授权码
          const { code } = await this.getCode(authorizationUrl, clientId, redirectUri);
  
          // 使用授权码获取访问令牌
          const accessToken = await this.getAccessToken(code, clientId, redirectUri);
  
          // 发送访问令牌到后端进行处理
          await this.sendAccessToken(accessToken);
  
          // 处理成功后的逻辑
          console.log('访问令牌已发送到后端进行处理');
        } catch (error) {
          console.error('登录失败', error);
        }
      },
  
      async getCode(authorizationUrl, clientId, redirectUri) {
        const response = await axios.get(authorizationUrl, {
          params: {
            client_id: clientId,
            redirect_uri: redirectUri,
            scope: 'user',
          },
        });
  
        const code = response.data.code;
        return code;
      },
  
      async getAccessToken(code, clientId, redirectUri) {
        const tokenUrl = 'https://github.com/login/oauth/access_token';
        const response = await axios.post(tokenUrl, null, {
          params: {
            client_id: clientId,
            client_secret: 'YOUR_CLIENT_SECRET',
            code: code,
            redirect_uri: redirectUri,
          },
          headers: {
            Accept: 'application/json',
          },
        });
  
        const accessToken = response.data.access_token;
        return accessToken;
      },
  
      async sendAccessToken(accessToken) {
        const backendUrl = 'https://localhost:8888/loginGitHub';
        const data = {
          access_token: accessToken,
        };
  
        await axios.post(backendUrl, data);
      },
    },
  };
  </script>
  