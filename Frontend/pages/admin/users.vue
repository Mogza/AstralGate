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
        Users
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
                <th class="border border-gray-300 px-4 py-2">Username</th>
                <th class="border border-gray-300 px-4 py-2">First Name</th>
                <th class="border border-gray-300 px-4 py-2">Last Name</th>
                <th class="border border-gray-300 px-4 py-2">Email</th>
                <th class="border border-gray-300 px-4 py-2">Phone Number</th>
              </tr>
              </thead>
              <tbody>
              <tr
                  v-for="user in users"
                  :key="user.id"
                  class="text-center odd:bg-gray-100 even:bg-gray-200"
              >
                <td class="border border-gray-300 px-4 py-2">{{ user.username }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ user.first_name }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ user.last_name }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ user.email }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ user.phone_number }}</td>
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

interface User {
  id: number;
  username: string;
  first_name: string;
  last_name: string;
  email: string;
  role: string;
  phone_number: string;
  created_at: string;
  updated_at: string;
}

const router = useRouter();
const users = ref<User[]>([]);
const BEARER_TOKEN = useRuntimeConfig().public.adminToken;

async function fetchUsers() {
  try {
    const response = await axios.get("http://185.157.245.42:8080/admin/users", {
      headers: {
        Authorization: `Bearer ${BEARER_TOKEN}`,
      },
    });
    users.value = response.data;
  } catch (error) {
    console.error("Error fetching users:", error);
  }
}

onMounted(() => {
  fetchUsers();
});

const currentTab = ref("Users");

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
