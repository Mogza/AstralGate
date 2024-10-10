<template>
  <div>
    <div class="centered-text">

      <p>
        Êtes-vous prêts à accepter les paiements en<br />
        <span class="gradient-word">crypto-monnaies</span> à n'importe quel endroit du <span class="gradient-word">globe</span> ?
      </p>
    </div>
    <div class="description-text">
      <p>
        Développez votre commerce en toute transparence avec des transactions<br />
        rapides, sécurisées et mondiales de crypto-monnaies grâce à AstralGate.
      </p>
    </div>
    <div>
      <button class="gradient-button">
        <span>COMMENCER GRATUITEMENT</span>
        <img src="~assets/icons/Right.svg" alt="Right-Icon" class="right-icon" />
      </button>
    </div>
    <img src="~assets/img/EtherCoin.png" alt="EtherCoin" class="ether_coin1" />
    <img src="~assets/img/EtherCoin.png" alt="EtherCoin" class="ether_coin2" />
    <img src="~assets/img/EtherCoin.png" alt="EtherCoin" class="ether_coin3" />
  </div>
</template>


<script setup lang="ts">
</script>

<style scoped>
.centered-text {
  display: inline-block;
  border-radius: 10px;
  position: relative;
  left: 50%;
  transform: translateX(-50%);
  top: 250px;
  text-align: center;
  font-size: 3.3rem;
  font-weight: 900;
  color: white;
  z-index: 1;
}

.description-text {
  display: inline-block;
  border-radius: 10px;
  position: relative;
  left: 50%;
  transform: translateX(-50%);
  top: 190px;
  text-align: center;
  font-size: 1.8rem;
  font-weight: 700;
  color: white;
  opacity: 70%;
  z-index: 1;
}

.gradient-word {
  background: linear-gradient(160deg, #C252E1, #6ECBF5);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}

.gradient-button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 15px 25px;
  font-size: 1.5rem;
  font-weight: 500;
  position: relative;
  left: 50%;
  transform: translateX(-50%);
  top: 250px;
  color: white;
  border: none;
  border-radius: 100px;
  cursor: pointer;
  background: transparent;
}

.right-icon {
  margin-left: 10px;
  width: 33px;
  height: 33px;
}

.gradient-button:before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border: none;
  border-radius: 100px;
  pointer-events: none;
  box-shadow: inset 0 0 0 2px white;
  transition: box-shadow 0.1s ease-in-out;
}

.gradient-button:hover:before {
  box-shadow: inset 0 0 0 0; /* Remove the inner border on hover */
}

.gradient-button:hover {
  background: linear-gradient(160deg, #C252E1, #6ECBF5);
  border: none;
}

.ether_coin1 {
  display: flex;
  width: 150px;
  height: 150px;
  margin-right: 10px;
  transform: translateY(-70px) translateX(225px);
  z-index: 0;
}

.ether_coin2 {
  display: flex;
  width: 300px;
  height: 300px;
  margin-right: 10px;
  transform: translateY(-715px) translateX(990px);
  z-index: 0;
}

.ether_coin3 {
  display: flex;
  width: 350px;
  height: 350px;
  margin-right: 10px;
  transform: translateY(-425px) translateX(1525px) scaleX(-1);
  z-index: 0;
}
</style>