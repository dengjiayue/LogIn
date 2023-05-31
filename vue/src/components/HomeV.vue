<template>
  <div>
    <h1>Welcome to HomeV</h1>
    <button @click="confirmLogout">登出</button>
    <button @click="changePassword">改密</button>
    <button @click="getOperations">查询</button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  methods: {
    async confirmLogout() {
      const confirmLogout = confirm('是否确认登出？');
      if (confirmLogout) {
        try {
          const idcode = this.$cookies.get('idcode');
          const requestData = { idcode };
          const response = await axios.post('http://localhost:8888/signout', requestData);
          if (response.data.code === 200) {
            this.$router.push('/');
          }else if (response.data.code === 420) {
            alert(response.data.msg);
            this.$router.push('/');
          } else {
            alert(response.data.msg);
          }
        } catch (error) {
          console.error('登出失败', error);
          alert('操作失败');
          this.$router.push('/');
        }
      }
    },
    changePassword() {
      this.$router.push('/RePasswordPsw');
    },
    async getOperations() {
      try {
        const idcode = this.$cookies.get('idcode');
        const requestData = { idcode };
        const response = await axios.post('http://localhost:8888/getoperations', requestData);
        if (response.data.code === 200) {
          const operations = response.data.data.operations;
          // 在此处理返回的操作数据，可以存储到组件的数据中并进行展示
          console.log(operations);
        }else if (response.data.code === 420) {
            alert(response.data.msg);
            this.$router.push('/');
          } else {
          alert(response.data.msg);
        }
      } catch (error) {
        console.error('查询操作失败', error);
        if (error.response && error.response.status === 420) {
          alert('非法访问');
          this.$router.push('/');
        }
      }
    }
  }
};
</script>
