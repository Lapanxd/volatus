<script lang="ts" setup>
import {inject, onMounted, onUnmounted, ref} from "vue";
import {Nullable} from "./core/types/nullable.ts";
import {User} from "./core/models/user.ts";
import {Store} from "@tauri-apps/plugin-store";

const store = inject<Store>("store")!;

const sseEvents = ref<string[]>([]);
let eventSource: Nullable<EventSource> = null;

const startSSE = async () => {
  // todo: use .env with API url
  const token = await store.get<string>("auth_token");

  if (!token) {
    return;
  }

  const encodedToken = encodeURIComponent(token);

  eventSource = new EventSource(`http://localhost:8080/sse/events?token=${encodedToken}`);

  eventSource.onopen = () => {
    console.log("SSE connected");
  };

  eventSource.onmessage = (e) => {
    console.log("Received SSE:", e.data);
    sseEvents.value.push(e.data);
  };

  eventSource.addEventListener("handshake_accepted", (e) => {
    console.log("Handshake accepted:", e.data);
    sseEvents.value.push(e.data);
  });

  eventSource.onerror = (err) => {
    console.error("SSE error or disconnected:", err);
    console.log("readyState:", eventSource?.readyState);
    if (eventSource?.readyState === 2) {
      eventSource = null;
    }
  };
}

const stopSSE = () => {
  eventSource?.close();
  eventSource = null;
};

onMounted(async () => {
  const u = await store.get<User>("user");

  if (u) {
    void startSSE();
  }
})

onUnmounted(() => stopSSE());
</script>

<template>
  <div class="app-container">
    <RouterView/>
  </div>

</template>

<style scoped></style>
<style>
body {
  margin: 0;
  padding: 0;
  font-family: "Inter", sans-serif;
}

* {
  margin: 0;
}

.app-container {
  height: 100vh;
  width: 100vw;
}

p {
  font-size: 0.9rem;
}
</style>