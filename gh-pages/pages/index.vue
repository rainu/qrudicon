<template>
  <div>
    <generator-simple v-model="imgSettings" />

    <hr v-if="imgSettings.text" />
    <qrudicon v-model="imgSettings" />
  </div>
</template>

<script>
  import GeneratorSimple from "../components/generator/Simple";
  import Qrudicon from "../components/Qrudicon";

  export default {
    components: {
      Qrudicon,
      GeneratorSimple
    },
    data() {
      return {
        imgSettings: {},
      }
    },
    watch: {
      imgSettings: {
        deep: true,
        handler() {
          setTimeout(() => {
            window.scrollTo(0, document.body.scrollHeight)
          }, 500)
        }
      }
    }
  }
</script>

<style>
</style>
