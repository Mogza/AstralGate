<template>
  <div class="bg-[#f4f3f2] flex h-screen">
    <!-- Sidebar -->
    <AdminSideBar @tab-changed="handleTabChange"/>

    <!-- Title -->
    <h1 class="text-transparent bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text font-bold text-5xl mt-10 ml-20">Users</h1>

    <!-- ContentBox -->
    <div class="flex-1 flex items-center mt-14 -ml-20">
      <div class="w-11/12 h-5/6 rounded-lg shadow-lg bg-gradient-to-r from-purple-400 to-blue-400 p-2 opacity-20">
        <div class="w-full h-full bg-[#f4f3f2] rounded-lg"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const currentTab = ref("Users");

function handleTabChange(newTab: string) {
  currentTab.value = newTab;
  console.log(`Tab changed to: ${newTab}`);
}
</script>
