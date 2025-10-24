<script lang="ts" setup>
import {ref} from "vue";
import {invoke} from "@tauri-apps/api/core";

let pc: RTCPeerConnection;
let dataChannel: RTCDataChannel;

const userId = ref(null)

const createConnection = () => {
  pc = new RTCPeerConnection();

  dataChannel = pc.createDataChannel("chat");
  dataChannel.onopen = () => console.log("DataChannel ouvert !");
  dataChannel.onmessage = (event) => console.log("Received message :", event.data);

  pc.onicecandidate = (event) => {
    if (event.candidate) {
      console.log("New ICE candidate :", event.candidate);
    } else {
      console.log("All ICE candidates gathered");
    }
  };
};

const generateOffer = async (): Promise<string> => {
  if (!pc) createConnection();

  const offer = await pc.createOffer();
  await pc.setLocalDescription(offer);

  return offer.sdp!;
};

const handshakeInit = async () => {
  if (!userId.value) return;
  const sdpOffer = await generateOffer();
  console.log("sdpOffer", sdpOffer);
  await invoke("handshake_init", {toUserId: userId.value, sdpOffer: sdpOffer});
}
</script>

<template>
  <div class="new-handshake">
    <p>New conversation</p>
    <div class="content">
      <input v-model="userId" type="number"/>
      <button @click="handshakeInit">Create</button>
    </div>
  </div>
</template>

<style scoped>
.new-handshake {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.content {
  display: flex;
  gap: 0.5rem;
}
</style>