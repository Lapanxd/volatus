<script lang="ts" setup>
import {invoke} from "@tauri-apps/api/core";
import {ref} from "vue";
import {RouteNames, router} from "../main.ts";

const username = ref("");
const password = ref("");

const onLogin = async () => {
  await invoke("login", {username: username.value, password: password.value});
  await invoke("get_me");
  void router.push(RouteNames.Home)

}
</script>

<template>
  <p>Login Page</p>
  <form>
    <input v-model="username" placeholder="Username" type="text"/>
    <input v-model="password" placeholder="Password" type="password"/>
    <button type="button" @click="onLogin">Login</button>
  </form>

</template>

<style scoped>
form {
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
}

form input,
form button {
  width: 200px;
}
</style>