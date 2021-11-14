<template>
  <div>
    <Hero title="Autonomous AI Asset Marketplace" desc="Combining the blockchain (Smart contract, NFT, IPFS), privacy-preserving computation and semantic web technologies, together with AI can take data utilization to a new level, building a decentralized, collaborative, trustworthy and autonomous AI asset marketplace."/>

    <ModelList/>

  </div>
</template>

<script>


  import Hero from "../components/Hero";
  import ModelList from "../components/ModelList";
  export default {
    name: 'Home',

    components: {
      ModelList,
      Hero

    },
  }
</script>

<style scoped>

</style>
