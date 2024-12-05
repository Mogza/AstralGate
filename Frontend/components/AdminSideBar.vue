<template>
  <div class="bg-[#131419] text-white w-64 p-4">
    <h1 class="text-xl font-extrabold mb-4">AstralGate <span class="font-light">Admin</span></h1>
    <ul class="mt-24">
      <li
          v-for="tab in tabs"
          :key="tab"
          :class="['py-5 px-4 rounded cursor-pointer text-center text-xl', currentTab === tab ? 'bg-gray-700 bg-opacity-70' : 'hover:bg-gray-600 hover:bg-opacity-20']"
          @click="changeTab(tab)"
      >
        <img :src="tabIcons[tab]" alt="tab icon" class="w-8 h-8 inline-block mr-4" />
        {{ tab }}
      </li>
    </ul>
    <div class="flex justify-center mt-10 relative z-10">
      <button @click="disconnect" class="flex items-center mt-[30rem] space-x-2 px-6 py-3 text-lg font-medium rounded-full bg-red-700 hover:bg-red-800 transition-shadow shadow-lg">
        <span>Se déconnecter</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from "vue";

import UsersIcon from "../assets/icons/Users.svg";
import ProductsIcon from "../assets/icons/Products.svg";
import TransactionsIcon from "../assets/icons/Transactions.svg";

const props = defineProps({
  currentTab: {
    type: String,
    required: true,
  },
});

const emits = defineEmits(["tab-changed", "disconnect"]);

const tabs = ["Users", "Products", "Transactions"];
const tabIcons: Record<string, string> = {
  Users: UsersIcon,
  Products: ProductsIcon,
  Transactions: TransactionsIcon,
};

function changeTab(tab: string) {
  emits("tab-changed", tab);
}

function disconnect() {
  emits("disconnect");
}
</script>
