<script lang="ts" setup>
import {invoke} from "@tauri-apps/api/core";
import {onMounted, ref, watchEffect} from "vue";
import {PendingHandshakesOutput, PendingHandshakeWithUser} from "../core/models/pending-handshakes-output.ts";
import {User} from "../core/models/user.ts";
import NewHandshake from "./NewHandshake.vue";

const peerConnections = ref<Record<string, RTCPeerConnection>>({});

const handshakes = ref<PendingHandshakesOutput[]>([]);
const handshakesWithUsers = ref<PendingHandshakeWithUser[]>([])


const getHandshakes = async () => {
  handshakes.value = await invoke("get_pending_handshakes");
}

const responseHandshake = async (session_id: string, accepted: boolean, sdp_offer: string) => {
  try {
    let sdpAnswer;
    if (accepted) {
      sdpAnswer = await createAnswer(session_id, sdp_offer);
    }

    await invoke("handshake_response", {sessionId: session_id, accepted, sdpAnswer: sdpAnswer});
    await getHandshakes();
  } catch (error) {
    console.error(error);
  }
}

const createAnswer = async (sessionId: string, sdpOffer: string) => {
  const pc = new RTCPeerConnection();

  peerConnections.value[sessionId] = pc;

  pc.ondatachannel = (event) => {
    const dataChannel = event.channel;
    dataChannel.onmessage = (e) => console.log("Received:", e.data);
    dataChannel.onopen = () => console.log("DataChannel ouvert !");
  };

  pc.onicecandidate = (event) => {
    if (event.candidate) {
      console.log("New ICE candidate:", event.candidate);
    } else {
      console.log("All ICE candidates gathered");
    }
  };

  await pc.setRemoteDescription({type: "offer", sdp: sdpOffer});

  const answer = await pc.createAnswer();
  await pc.setLocalDescription(answer);

  return answer.sdp!;
};


const getUserById = async (userId: number): Promise<User> => {
  return await invoke<User>("get_user_by_id", {id: userId});
}

watchEffect(async () => {
  handshakesWithUsers.value = await Promise.all(
      handshakes.value.map(async h => ({
        sessionId: h.session_id,
        fromUser: await getUserById(h.from_user_id),
        sdpOffer: h.sdp_offer
      }))
  )
})

onMounted(async () => {
  await getHandshakes();
});
</script>

<template>
  <div class="handshake">
    <NewHandshake/>

    <div class="header">
      <p>Pending requests</p>
      <button @click="getHandshakes">Refresh</button>
    </div>
    <div v-for="handshake in handshakesWithUsers" class="pending-handshake">
      <div class="name">
        <img :src="`https://ui-avatars.com/api/?name=${handshake.fromUser.username}`" alt="User's avatar"/>
        <p>{{ handshake.fromUser.username }}</p>
      </div>
      <div class="actions">
        <button @click="responseHandshake(handshake.sessionId, true, handshake.sdpOffer)">Accept</button>
        <button @click="responseHandshake(handshake.sessionId, false, handshake.sdpOffer)">Refuse</button>
      </div>
    </div>
  </div>

</template>

<style scoped>
.handshake {
  background: #f3f3f3;
  padding: 1rem;
  border-radius: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.header {
  display: flex;
  align-items: start;
  justify-content: space-between;
}

.pending-handshake {
  display: flex;
  gap: 0.5rem;
  align-items: center;
  justify-content: space-between;
  background: #ffffff;
  padding: 0.5rem;
  border-radius: 0.8rem;
}

.pending-handshake img {
  height: 30px;
  width: 30px;
  border-radius: 50%;
}

.pending-handshake .name {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.pending-handshake .actions {
  display: flex;
  align-items: center;
  gap: 0.2rem;
}
</style>