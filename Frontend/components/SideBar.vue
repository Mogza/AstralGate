<template>
  <div class="bg-[#131419] text-white w-64 p-4">
    <h1 class="text-xl font-extrabold mb-4">AstralGate</h1>
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
      <button @click="settings" class="flex items-center mt-[30rem] space-x-2 px-6 py-3 text-lg font-medium rounded-full bg-white/5 hover:bg-white/10 transition-shadow shadow-lg">
        <span>Settings</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from "vue";

import ProductsIcon from "../assets/icons/Products.svg";
import DashboardIcon from "../assets/icons/Receipt.svg";
import TransactionIcon from "../assets/icons/Transactions.svg";
import {useRouter} from "#vue-router";

const props = defineProps({
  currentTab: {
    type: String,
    required: true,
  },
});

const emits = defineEmits(["tab-changed"]);

const tabs = ["Dashboard", "Products", "Transactions"];
const tabIcons: Record<string, string> = {
  Dashboard: DashboardIcon,
  Products: ProductsIcon,
  Transactions: TransactionIcon
};

function changeTab(tab: string) {
  emits("tab-changed", tab);
}

const router = useRouter();
function settings() {
  router.push("/merchants/settings");
}
</script>
