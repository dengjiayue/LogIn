<template>
  <div>
    <h2>注册2</h2>
    <p>hallo:{{ email }}</p>
  </div>
  <div>
  <form @submit="submitForm">
      <label for="name">Name:</label>
      <input type="text" id="name" v-model="name" required>
      <br>
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="password" required>
      <br>
      <label for="verification">Verification Code:</label>
      <input type="text" id="verification" v-model="verification" pattern="[0-9]{6}" required>
      <button @click="ReVerify">重发验证码</button>
      <br>
      <button type="submit">Submit</button>
    </form>
  </div>
  <div>
    <button @click="goToLogin">登录</button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      name: '',
      password: '',
      email: '',
      verification: '',
    };
  },
  created() {
    this.email = this.$route.query.email || '';
  },
  methods: {
    ReVerify() {
        this.$router.push('/signup1');
      },
    goToLogin() {
        this.$router.push('/LogInPsw');
      },
    async submitForm(event) {
      event.preventDefault();
      if (!this.validateName(this.name)) {
        console.error('Name must be less than 20 characters');
        return;
      }
      if (!this.validatePassword(this.password)) {
        console.error('Password must be less than 20 characters');
        return;
      }
      if (!this.validateVerification(this.verification)) {
        console.error('Verification code must be 6 digits');
        return;
      }
      try {
        const requestData = {
          name: this.name,
          password: this.password,
          emailverificationcode: {
            email: this.email,
            verification: parseInt(this.verification),
          },
        };
        const response = await axios.post('http://localhost:8888/signup2', requestData);
        if (response.data.code === 200) {
          console.log('Signup successful');
        } else {
          console.error(response.data.msg);
        }
      } catch (error) {
        console.error('Signup failed', error);
      }
    },
    validateName(name) {
      return name.length > 0 && name.length <= 20;
    },
    validatePassword(password) {
      return password.length > 0 && password.length <= 20;
    },
    validateVerification(verification) {
      return verification.length === 6 && /^\d+$/.test(verification);
    },
  },
};
</script>
