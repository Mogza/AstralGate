<template>
  <div>
    <!-- Header placed outside the main content -->
    <Header />

    <!-- Main Content -->
    <div class="relative text-white min-h-screen pt-20">
      <!-- EtherCoin -->
      <div class="absolute top-[30%] left-[45%] transform -translate-x-1/2 -translate-y-1/2 opacity-100 z-0">
        <img src="~assets/img/EtherCoin.png" alt="EtherCoin" class="max-w-[10%] md:max-w-[10%]" />
      </div>
      <!-- EtherCoin 2 -->
      <div class="absolute top-[40%] left-[87%] transform -translate-x-1/2 -translate-y-1/2 opacity-100 z-0">
        <img src="~assets/img/EtherCoin.png" alt="EtherCoin" class="max-w-[110%] md:max-w-[110%] transform scale-x-[-1]" />
      </div>
      <!-- EtherCoin 3 -->
      <div class="absolute top-[0%] left-[65%] transform -translate-x-1/2 -translate-y-1/2 opacity-100 z-0">
        <img src="~assets/img/EtherCoin.png" alt="EtherCoin" class="max-w-[200%] md:max-w-[60%]" />
      </div>

      <!-- Centered Text -->
      <div class="text-center font-extrabold text-5xl md:text-5xl sm:text-4xl mt-32 sm:mt-40 md:mt-52 mx-auto max-w-md md:max-w-7xl relative z-10">
        <p>
          Êtes-vous prêts à accepter les paiements en<br />
          <span class="text-transparent bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text">crypto-monnaies</span>
          à n'importe quel endroit du <span class="text-transparent bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text">globe</span> ?
        </p>
      </div>

      <!-- Description Text -->
      <div class="text-center font-semibold text-xl md:text-2xl sm:text-base opacity-70 mt-8 mx-auto max-w-7xl relative z-10">
        <p>
          Développez votre commerce en toute transparence avec des transactions<br />
          rapides, sécurisées et mondiales de crypto-monnaies grâce à AstralGate.
        </p>
      </div>

      <!-- Gradient Button -->
      <div class="flex justify-center mt-10 relative z-10">
        <button class="flex items-center space-x-2 px-6 py-3 text-xl font-medium rounded-full bg-gradient-to-r from-purple-400 to-blue-400 hover:from-purple-500 hover:to-blue-500 transition-shadow shadow-lg">
          <span>COMMENCER GRATUITEMENT</span>
          <img src="~assets/icons/Right.svg" alt="Right-Icon" class="w-8 h-8" />
        </button>
      </div>

      <!-- Additional Text -->
      <div class="text-center font-semibold text-lg md:text-base sm:text-base opacity-70 mt-2 mx-auto max-w-7xl relative z-10">
        <p>
          Aucune carte de crédit n'est requise
        </p>
      </div>

      <!-- Subscription Cards -->
      <div class="flex justify-center mt-20 space-x-6 px-4 pt-60 relative z-10">
        <!-- Start Card -->
        <div class="w-72 border-2 border-purple-400 text-white p-6 rounded-lg shadow-lg">
          <h3 class="text-2xl font-extrabold mb-4 text-center">Free</h3>
          <h3 class="text-2xl font-extrabold mb-4 text-center text-transparent bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text">Starter</h3>
          <p class="mb-7 pt-5">
            - ...<br />
            - ...<br />
            - ...<br />
            - ...<br />
            - ...<br />
          </p>
          <button class="w-full py-3 text-xl font-medium rounded-full bg-gradient-to-r from-purple-400 to-blue-400 hover:from-purple-500 hover:to-blue-500 transition-shadow shadow-lg">
            Start Now
          </button>
        </div>

        <!-- 25$ Card (Filled Style) -->
        <div class="w-72 bg-gradient-to-r from-purple-400 to-blue-400 text-white p-6 rounded-lg shadow-lg">
          <h3 class="text-2xl font-extrabold mb-4 text-center">25$</h3>
          <h3 class="text-2xl font-extrabold mb-4 text-center text-black">Business</h3>
          <p class="mb-7 pt-5">
            - ...<br />
            - ...<br />
            - ...<br />
            - ...<br />
            - ...<br />
          </p>
          <button class="w-full py-3 text-xl font-medium rounded-full bg-white text-purple-400 hover:bg-gray-200 transition-all">
            Order Now
          </button>
        </div>

        <!-- 250$ Card (Outlined Style) -->
        <div class="w-72 border-2 border-purple-400 text-white p-6 rounded-lg shadow-lg">
          <h3 class="text-2xl font-extrabold mb-4 text-center">250$</h3>
          <h3 class="text-2xl font-extrabold mb-4 text-center text-transparent bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text">Lifetime Access</h3>
          <p class="mb-7 pt-5">
            - ...<br />
            - ...<br />
            - ...<br />
            - ...<br />
            - ...<br />
          </p>
          <button class="w-full py-3 text-xl font-medium rounded-full bg-gradient-to-r from-purple-400 to-blue-400 hover:from-purple-500 hover:to-blue-500 transition-shadow shadow-lg">
            Order Now
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Header from './components/Header.vue'
</script>
