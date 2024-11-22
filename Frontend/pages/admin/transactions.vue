<template>
  <div class="bg-[#f4f3f2] flex h-screen">
    <!-- Sidebar -->
    <AdminSideBar :current-tab=currentTab @tab-changed="handleTabChange" />

    <!-- Title -->
    <h1
        class="text-transparent bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text font-bold text-5xl mt-10 ml-20"
    >
      Transactions
    </h1>

    <!-- ContentBox -->
    <div class="flex-1 flex items-center mt-14 -ml-20">
      <div
          class="w-11/12 h-5/6 rounded-lg shadow-lg bg-gradient-to-r from-purple-400 to-blue-400 p-2 opacity-100"
      >
        <div class="w-full h-full bg-[#f4f3f2] rounded-lg p-4 overflow-y-auto">
          <table class="table-auto w-full border-collapse border border-gray-300">
            <thead>
            <tr class="bg-gradient-to-r from-purple-400 to-blue-400 text-white">
              <th class="border border-gray-300 px-4 py-2">Wallet Id</th>
              <th class="border border-gray-300 px-4 py-2">Product Title</th>
              <th class="border border-gray-300 px-4 py-2">Client Address</th>
              <th class="border border-gray-300 px-4 py-2">Amount</th>
              <th class="border border-gray-300 px-4 py-2">Currency</th>
              <th class="border border-gray-300 px-4 py-2">Status</th>
            </tr>
            </thead>
            <tbody>
            <tr
                v-for="transaction in transactions"
                :key="transaction.id"
                class="text-center odd:bg-gray-100 even:bg-gray-200"
            >
              <td class="border border-gray-300 px-4 py-2">{{ transaction.wallet_id }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ transaction.product_title }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ transaction.client_address }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ transaction.amount }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ transaction.currency }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ transaction.status }}</td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRouter } from 'vue-router';
import axios from "axios";

interface Transaction {
  id: number;
  wallet_id: number;
  product_id: number;
  product_title: string;
  client_address: string;
  amount: number;
  currency: string;
  status: string;
}

const router = useRouter();
const transactions = ref<Transaction[]>([]);
const BEARER_TOKEN = useRuntimeConfig().public.adminToken;

async function fetchUsers() {
  try {
    const response = await axios.get("http://185.157.245.42:8080/admin/transaction", {
      headers: {
        Authorization: `Bearer ${BEARER_TOKEN}`,
      },
    });
    transactions.value = response.data;

    for (const transaction of transactions.value as Transaction[]) {
      try {
        const productResponse = await axios.get(`http://185.157.245.42:8080/api/products/${transaction.product_id}`, {
          headers: {
            Authorization: `Bearer ${BEARER_TOKEN}`,
          },
        });
        transaction.product_title = productResponse.data.title;
      } catch (userError) {
        console.error(`Error fetching user data for product_id ${transaction.product_id}:`, userError);
      }
    }
  } catch (error) {
    console.error("Error fetching users:", error);
  }
}

onMounted(() => {
  fetchUsers();
});

const currentTab = ref("Transactions");

function handleTabChange(newTab: string) {
  currentTab.value = newTab.toLowerCase();
  console.log(`Tab changed to: ${newTab}`);
  router.push(`/admin/${newTab.toLowerCase()}`);
}
</script>
