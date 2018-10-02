<template>
  <div>
    <b-container class="bv-example-row">
      <b-row>
        <b-col >
          <generator-simple v-model="imgSettings" />
        </b-col>
        <b-col >
          <qrudicon v-model="imgSettings" />
        </b-col>
      </b-row>
    </b-container>
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
    }
  }
</script>

<style>
</style>
