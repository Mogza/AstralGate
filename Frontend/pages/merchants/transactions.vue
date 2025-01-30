<template>
  <div class="bg-[#f4f3f2] flex h-screen">
    <!-- Sidebar -->
    <SideBar :current-tab=currentTab @tab-changed="handleTabChange" @disconnect="handleDisconnection"/>

    <!-- Page Content -->
    <div class="flex-1 flex flex-col">
      <!-- Title -->
      <h1
          class="text-transparent bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text font-bold text-5xl mt-10 ml-20"
      >
        Products
      </h1>

      <!-- ContentBox -->
      <div class="flex-1 flex mt-14 ml-24">
        <div
            class="w-11/12 h-5/6 rounded-lg shadow-lg bg-gradient-to-r from-purple-400/20 to-blue-400/20 p-2 opacity-100"
        >
          <div class="w-full h-full bg-[#f4f3f2] rounded-lg p-4 overflow-y-auto">
            <table class="table-auto w-full border-collapse border border-gray-300">
              <thead>
              <tr class="bg-gradient-to-r from-purple-400 to-blue-400 text-white">
                <th class="border border-gray-300 px-4 py-2">Client Address</th>
                <th class="border border-gray-300 px-4 py-2">Amount</th>
                <th class="border border-gray-300 px-4 py-2">Currency</th>
                <th class="border border-gray-300 px-4 py-2">Status</th>
                <th class="border border-gray-300 px-4 py-2">Tx Hash</th>
              </tr>
              </thead>
              <tbody>
              <tr
                  v-for="transaction in transactions"
                  :key="transaction.id"
                  class="text-center odd:bg-gray-100 even:bg-gray-200"
              >
                <td class="border border-gray-300 px-4 py-2">{{ transaction.client_address }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ transaction.amount }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ transaction.currency }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ transaction.status }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ transaction.tx_hash }}</td>
              </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Export Button -->
      <button @click="exportTransactions" class="items-center ml-[1300px] w-60 px-6 py-3 text-xl text-white font-bold rounded-full bg-gradient-to-br from-purple-400 to-blue-400 hover:from-purple-500 hover:to-blue-500 transition-shadow shadow-lg">
        <span>Export your transactions in .csv</span>
      </button>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRouter } from 'vue-router';
import axios from "axios";
import Cookies from "js-cookie";

interface Transaction {
  id: number;
  client_address: string;
  amount: string;
  currency: number;
  status: string;
  tx_hash: string;
}

const router = useRouter();
const token = Cookies.get("auth_token")

const transactions = ref<Transaction[]>([]);
async function fetchTransactions() {
  try {
    const response = await axios.get("http://185.157.245.42:8080/api/users/transactions/me", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    transactions.value = response.data;
  } catch (error) {
    console.error("Error fetching transactions:", error);
  }
}


onMounted(() => {
  fetchTransactions();
});

const currentTab = ref("Transactions");

function handleTabChange(newTab: string) {
  currentTab.value = newTab.toLowerCase();
  console.log(`Tab changed to: ${newTab}`);
  router.push(`/merchants/${newTab.toLowerCase()}`);
}

function handleDisconnection() {
  Cookies.remove("auth_token");
  router.push("/");
}

async function exportTransactions() {
  try {
    const response = await axios.get("http://185.157.245.42:8080/api/transactions/export", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      responseType: 'blob',
    });

    if (response.data.size === 0) {
      throw new Error('No data received');
    }

    const blob = new Blob([response.data], { type: 'text/csv' });
    const url = window.URL.createObjectURL(blob);

    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', 'transactions.csv');

    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error("Error downloading transactions:", error);
  }
}

</script>
