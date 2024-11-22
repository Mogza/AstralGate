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

const emits = defineEmits(["tab-changed"]);

const tabs = ["Users", "Products", "Transactions"];
const tabIcons: Record<string, string> = {
  Users: UsersIcon,
  Products: ProductsIcon,
  Transactions: TransactionsIcon,
};

function changeTab(tab: string) {
  emits("tab-changed", tab);
}
</script>
