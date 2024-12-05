<template>
  <div class="bg-[#f4f3f2] flex h-screen">
    <!-- Sidebar -->
    <AdminSideBar :current-tab=currentTab @tab-changed="handleTabChange" @disconnect="handleDisconnection"/>

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
                <th class="border border-gray-300 px-4 py-2">User Name</th>
                <th class="border border-gray-300 px-4 py-2">Title</th>
                <th class="border border-gray-300 px-4 py-2">Description</th>
                <th class="border border-gray-300 px-4 py-2">Price</th>
              </tr>
              </thead>
              <tbody>
              <tr
                  v-for="product in products"
                  :key="product.id"
                  class="text-center odd:bg-gray-100 even:bg-gray-200"
              >
                <td class="border border-gray-300 px-4 py-2">{{ product.user_name }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ product.title }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ product.description }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ product.usd_price }}$</td>
              </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRouter } from 'vue-router';
import axios from "axios";
import Cookies from "js-cookie";

interface Product {
  id: number;
  user_id: number;
  user_name: string;
  title: string;
  description: string;
  usd_price: number;
}

const router = useRouter();
const products = ref<Product[]>([]);
const BEARER_TOKEN = useRuntimeConfig().public.adminToken;

async function fetchUsers() {
  try {
    const response = await axios.get("http://185.157.245.42:8080/admin/products", {
      headers: {
        Authorization: `Bearer ${BEARER_TOKEN}`,
      },
    });
    products.value = response.data;

    for (const product of products.value as Product[]) {
      try {
        const userResponse = await axios.get(`http://185.157.245.42:8080/admin/users/${product.user_id}`, {
          headers: {
            Authorization: `Bearer ${BEARER_TOKEN}`,
          },
        });
        product.user_name = userResponse.data.username;
      } catch (userError) {
        console.error(`Error fetching user data for user_id ${product.user_id}:`, userError);
      }
    }
  } catch (error) {
    console.error("Error fetching users:", error);
  }
}

onMounted(() => {
  fetchUsers();
});

const currentTab = ref("Products");

function handleTabChange(newTab: string) {
  currentTab.value = newTab.toLowerCase();
  console.log(`Tab changed to: ${newTab}`);
  router.push(`/admin/${newTab.toLowerCase()}`);
}

function handleDisconnection() {
  Cookies.remove("auth_token");
  router.push("/");
}
</script>
