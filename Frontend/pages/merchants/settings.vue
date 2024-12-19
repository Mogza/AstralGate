<template>
  <div class="bg-[#f4f3f2] flex h-screen">

    <!-- Page Content -->
    <div class="flex-1 flex flex-col">
      <!-- Title -->
      <h1
          class="flex text-transparent bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text font-bold text-5xl mt-10 ml-20"
      >
        <button @click="back">
          <img src="~assets/icons/Back.svg" alt="Back-Icon" class="w-12 h-12 mr-3" />
        </button>
        Settings
      </h1>

      <!-- ContentBox -->
      <div class="flex-1 flex mt-14 ml-24">
        <div
            class="w-11/12 h-[700px] rounded-lg shadow-lg bg-gradient-to-r from-purple-400/20 to-blue-400/20 p-2 opacity-100"
        >
          <div class="w-full h-full bg-[#f4f3f2] text-white rounded-lg p-4 overflow-y-auto">
            <!-- Form -->
            <form @submit.prevent="updateUser" class="mt-12 w-11/12 mx-auto text-white text-xl flex flex-col">
              <p class="text-black/70 font-bold">
                Username
              </p>
              <!-- Username Field -->
              <div class="w-full h-14 bg-[#f4f3f2] p-6 rounded-[10px] shadow-lg border-2 border-[#404040] mb-4 flex items-center">
                <input
                    type="text"
                    :placeholder="userInfo?.username"
                    v-model="settingsForm.username"
                    class="bg-transparent text-black/70 text-xl outline-none w-[1465px]"
                />
              </div>
              <p class="text-black/70 font-bold">
                FirstName
              </p>
              <!-- First name Field -->
              <div class="w-full h-14 bg-[#f4f3f2] p-6 rounded-[10px] shadow-lg border-2 border-[#404040] mb-4 flex items-center">
                <input
                    type="text"
                    :placeholder="userInfo?.first_name"
                    v-model="settingsForm.first_name"
                    class="bg-transparent text-black/70 text-xl outline-none w-[1465px]"
                />
              </div>
              <p class="text-black/70 font-bold">
                LastName
              </p>
              <!-- Last name Field -->
              <div class="w-full h-14 bg-[#f4f3f2] p-6 rounded-[10px] shadow-lg border-2 border-[#404040] mb-4 flex items-center">
                <input
                    type="text"
                    :placeholder="userInfo?.last_name"
                    v-model="settingsForm.last_name"
                    class="bg-transparent text-black/70 text-xl outline-none w-[1465px]"
                />
              </div>
              <p class="text-black/70 font-bold">
                Email
              </p>
              <!-- Email Field -->
              <div class="w-full h-14 bg-[#f4f3f2] p-6 rounded-[10px] shadow-lg border-2 border-[#404040] mb-4 flex items-center">
                <input
                    type="text"
                    :placeholder="userInfo?.email"
                    v-model="settingsForm.email"
                    class="bg-transparent text-black/70 text-xl outline-none w-[1465px]"
                />
              </div>
              <p class="text-black/70 font-bold">
                Phone Number
              </p>
              <!-- Phone number Field -->
              <div class="w-full h-14 bg-[#f4f3f2] p-6 rounded-[10px] shadow-lg border-2 border-[#404040] mb-4 flex items-center">
                <input
                    type="text"
                    :placeholder="userInfo?.phone_number"
                    v-model="settingsForm.phone_number"
                    class="bg-transparent text-black/70 text-xl outline-none w-[1465px]"
                />
              </div>
              <!-- Submit Field -->
              <button type=submit class="flex items-center w-[191px] px-6 py-3 text-xl mt-12 font-bold rounded-full bg-gradient-to-br from-purple-400 to-blue-400 hover:from-purple-500 hover:to-blue-500 transition-shadow shadow-lg">
                <span>Update changes</span>
              </button>
            </form>
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

const router = useRouter();
const token = Cookies.get("auth_token")

interface User {
  id : number;
  username : string;
  first_name: string;
  last_name: string;
  email: string;
  phone_number: string
}

const userInfo = ref<User | null>(null);

async function fetchUser() {
  try {
    const response = await axios.get("http://185.157.245.42:8080/api/users/me", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    userInfo.value = response.data;
  } catch (error) {
    console.error("Error fetching users:", error);
  }
}
onMounted(() => {

  fetchUser();
});

let settingsForm = {
  username : '',
  first_name: '',
  last_name: '',
  email: '',
  phone_number: '',
};

const updateUser = async () => {
  try {

    if (settingsForm.username === '' && userInfo.value?.username) {
      settingsForm.username = userInfo.value.username;
    }
    if (settingsForm.first_name === '' && userInfo.value?.first_name) {
      settingsForm.first_name = userInfo.value.first_name;
    }
    if (settingsForm.last_name === '' && userInfo.value?.last_name) {
      settingsForm.last_name = userInfo.value.last_name;
    }
    if (settingsForm.email === '' && userInfo.value?.email) {
      settingsForm.email = userInfo.value.email;
    }
    if (settingsForm.phone_number === '' && userInfo.value?.phone_number) {
      settingsForm.phone_number = userInfo.value.phone_number;
    }

    await axios.put(`http://185.157.245.42:8080/api/users/${userInfo.value?.id}`, settingsForm,  {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  } catch (error) {
    console.error(error);
  }
  await router.push("/merchants/products")
}

function back() {
  router.push("/merchants/products");
}

</script>
