# pages/widget/index.vue
<template>
  <div class="min-h-screen bg-[#f4f3f2]">
    <div class="max-w-md mx-auto pt-6 px-4">
      <div class="rounded-lg shadow-lg bg-gradient-to-r from-purple-400/20 to-blue-400/20 p-2">
        <div class="bg-white rounded-lg p-6">
          <!-- Title -->
          <h2 class="text-2xl font-bold text-center mb-6">
            <span class="text-transparent bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text">
              Initialize Transaction
            </span>
          </h2>

          <!-- Alert Messages -->
          <div v-if="status.message"
               :class="[
                 'mb-4 p-4 rounded-lg text-sm',
                 status.type === 'error' ? 'bg-red-100 text-red-700' : 'bg-green-100 text-green-700'
               ]">
            {{ status.message }}
          </div>

          <!-- Form -->
          <form @submit.prevent="handleSubmit" class="space-y-4">
            <div>
              <label for="clientAddress" class="block text-sm font-medium text-gray-700 mb-1">
                Client Address
              </label>
              <input
                  id="clientAddress"
                  v-model="clientAddress"
                  type="text"
                  class="w-full p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-400 focus:border-transparent"
                  placeholder="Enter client address"
                  required
              />
            </div>

            <button
                type="submit"
                class="w-full py-2 px-4 bg-gradient-to-r from-purple-400 to-blue-400 text-white font-medium rounded-lg hover:from-purple-500 hover:to-blue-500 transition-all duration-200"
            >
              Submit Transaction
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const route = useRoute();
const productId = route.query.product_id;
const clientAddress = ref('');
const status = ref({ type: '', message: '' });

const handleSubmit = async () => {
  try {
    const response = await $fetch('http://185.157.245.42:8080/transaction/POL/', {
      method: 'POST',
      body: {
        product_id: productId,
        client_address: clientAddress.value
      }
    });

    status.value = {
      type: 'success',
      message: 'Transaction initiated successfully!'
    };
    clientAddress.value = '';

  } catch (error) {
    status.value = {
      type: 'error',
      message: 'Failed to initiate transaction. Please try again.'
    };
    console.error('Transaction error:', error);
  }
};
</script>

<style>
html, body {
  margin: 0;
  padding: 0;
  background: #f4f3f2;
  height: 100%;
}

#__nuxt {
  height: 100%;
}
</style>