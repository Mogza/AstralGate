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
        Create a product
      </h1>

      <!-- ContentBox -->
      <div class="flex-1 flex mt-14 ml-24">
        <div
            class="w-11/12 h-5/6 rounded-lg shadow-lg bg-gradient-to-r from-purple-400/20 to-blue-400/20 p-2 opacity-100"
        >
          <div class="w-full h-full bg-[#f4f3f2] text-white rounded-lg p-4 overflow-y-auto">
            <!-- S'inscrire -->
            <div class="text-black/70 font-extrabold text-4xl md:text-4xl sm:text-4xl mt-9 mx-auto max-w-md md:max-w-7xl">
              <p>
                Provide product details
              </p>
            </div>
            <!-- Form -->
            <form @submit.prevent="createProductSubmit" class="mt-12 w-11/12 mx-auto text-white text-xl flex flex-col">
              <p class="text-black/70 font-bold">
                Product name
              </p>
              <!-- Product Name Field -->
              <div class="w-full h-14 bg-[#f4f3f2] p-6 rounded-[10px] shadow-lg border-2 border-[#404040] mb-4 flex items-center">
                <input
                    type="text"
                    placeholder="Product Name"
                    v-model="createProductForm.title"
                    class="bg-transparent text-black/70 text-xl outline-none w-[1465px]"
                />
              </div>
              <p class="text-black/70 font-bold">
                Description
              </p>
              <!-- Description Field -->
              <div class="w-full h-28 bg-[#f4f3f2] p-4 rounded-[10px] shadow-lg border-2 border-[#404040] mb-4">
                <textarea
                    placeholder="Description"
                    v-model="createProductForm.description"
                    class="bg-transparent text-black/70 text-xl outline-none w-[1465px] resize-none"
                    rows="3"
                />
              </div>
              <p class="text-black/70 font-bold">
                Price in USD
              </p>
              <!-- Price Field -->
              <div class="w-full h-14 bg-[#f4f3f2] p-6 rounded-[10px] shadow-lg border-2 border-[#404040] mb-4 flex items-center">
                <input
                    type="text"
                    placeholder="Price"
                    v-model="createProductForm.usd_price"
                    class="bg-transparent text-black/70 text-xl outline-none w-[1465px]"
                />
              </div>
              <!-- Drag and Drop image -->
              <p class="text-black/70 font-bold">Product Image</p>
              <div
                  class="w-full h-32 bg-[#f4f3f2] rounded-[10px] shadow-lg border-2 border-dashed border-[#404040] mb-4 flex items-center justify-center"
                  @drop.prevent="handleFileDrop"
                  @dragover.prevent
              >
                <input type="file" @change="handleFileSelect" ref="fileInput" class="hidden" />
                <div @click="triggerFileSelect" class="cursor-pointer text-black/70">
                  Drag and drop an image here or click to select a file
                </div>
              </div>
              <!-- Image Preview -->
              <div v-if="imagePreview" class="mt-4">
                <p class="text-black/70 font-bold">Image Preview:</p>
                <img :src="imagePreview" alt="Image Preview" class="w-1/12 h-auto rounded-lg shadow-lg mt-2" />
              </div>
              <!-- Submit Field -->
              <button type=submit class="flex items-center w-[165px] px-6 py-3 text-xl mt-12 font-bold rounded-full bg-gradient-to-br from-purple-400 to-blue-400 hover:from-purple-500 hover:to-blue-500 transition-shadow shadow-lg">
                <span>Créer l'objet</span>
              </button>
            </form>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import axios from "axios";
import Cookies from "js-cookie";
import { ref } from 'vue';

let createProductForm = {
  title: '',
  description: '',
  usd_price: 0,
};

const router = useRouter();
const token = Cookies.get("auth_token")
const fileInput = ref<File | null>(null);
const imagePreview = ref<string | null>(null);

const handleFileSelect = (event: Event) => {
  const files = (event.target as HTMLInputElement).files;
  if (files && files[0]) {
    fileInput.value = files[0];
    generateImagePreview(files[0]);
  }
};

const handleFileDrop = (event: DragEvent) => {
  const files = event.dataTransfer?.files;
  if (files && files[0]) {
    fileInput.value = files[0];
    generateImagePreview(files[0]);
  }
};

const triggerFileSelect = () => {
  document.querySelector<HTMLInputElement>('input[type="file"]')?.click();
};

const generateImagePreview = (file: File) => {
  const reader = new FileReader();
  reader.onload = (e) => {
    imagePreview.value = e.target?.result as string;
  };
  reader.readAsDataURL(file);
};

const createProductSubmit = async () => {
  createProductForm.usd_price = Number(createProductForm.usd_price);

  const formData = new FormData();
  formData.append('title', createProductForm.title);
  formData.append('description', createProductForm.description);
  formData.append('usd_price', createProductForm.usd_price.toString());
  if (fileInput.value) {
    formData.append('image', fileInput.value);
  }

  try {
    await axios.post('http://185.157.245.42:8080/api/products/', formData, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'multipart/form-data',
      },
    });
    await router.push('/merchants/products');
  } catch (error) {
    console.error(error);
  }
}

function back() {
  router.push("/merchants/products");
}

</script>
