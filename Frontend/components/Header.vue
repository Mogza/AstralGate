<template>
  <header class="fixed top-0 left-0 w-full bg-transparent flex items-center justify-between p-6 z-50">
    <!-- Left side: Title and Buttons -->
    <div class="flex items-center space-x-4">
      <!-- Title (AstralGate) -->
      <span class="text-white font-extrabold text-xl">
        AstralGate
      </span>

      <!-- Accueil Button -->
      <button class="flex items-center space-x-2 px-6 py-3 text-xl font-normal text-white bg-transparent hover:underline transition-all">
        <span>Accueil</span>
      </button>

      <!-- Pricing Button -->
      <button class="flex items-center space-x-2 px-6 py-3 text-xl font-normal text-white bg-transparent hover:underline transition-all">
        <span>Pricing</span>
      </button>
    </div>

    <!-- Right side: Sign In and Sign Up Buttons -->
    <div class="flex items-center space-x-4 ml-auto">
      <!-- Sign In Button -->
      <button class="flex items-center space-x-2 px-6 py-3 text-xl font-bold rounded-full border-2 border-white hover:border-transparent hover:border-gradient-to-r hover:from-purple-400 hover:to-blue-400 bg-transparent hover:bg-gradient-to-r hover:from-purple-400 hover:to-blue-400 transition-all shadow-lg text-white">
        <span>Sign In</span>
      </button>

      <!-- Sign Up Button -->
      <button class="flex items-center space-x-2 px-6 py-3 text-xl font-bold rounded-full bg-gradient-to-r from-purple-400 to-blue-400 hover:from-purple-500 hover:to-blue-500 transition-shadow shadow-lg">
        <span>Sign Up</span>
      </button>
    </div>
  </header>
</template>

<script setup lang="ts">
</script>
