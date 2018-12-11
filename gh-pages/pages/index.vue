<template>
  <div>
    <ul class="nav nav-tabs" role="tablist">
      <li class="nav-item">
        <a class="nav-link" role="tab" :class="{ 'active': mode === 'simple' }" @click="setMode('simple')">Simple</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" role="tab" :class="{ 'active': mode === 'vcard' }" @click="setMode('vcard')">VCard</a>
      </li>
    </ul>
    <div class="tab-content">
      <div class="tab-pane fade show active border" role="tabpanel" aria-labelledby="home-tab">
        <generator-simple v-model="imgSettings" v-if="mode === 'simple'"/>
        <generator-vcard v-model="imgSettings" v-else-if="mode === 'vcard'"/>
      </div>
    </div>


    <hr v-if="imgSettings.text" />
    <qrudicon v-model="imgSettings.link" :mode="mode" />
  </div>
</template>

<script>
  import GeneratorSimple from "../components/generator/Simple";
  import GeneratorVcard from "../components/generator/Vcard";
  import Qrudicon from "../components/Qrudicon";

  export default {
    components: {
      Qrudicon,
      GeneratorSimple,
      GeneratorVcard,
    },
    data() {
      return {
        mode: 'simple',
        imgSettings: {
          link: '',
        },
      }
    },
    methods: {
      setMode(mode) {
        this.mode = mode;
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

<style scoped>
  .nav > li {
    cursor: pointer;
  }
  .tab-pane {
    padding: 1em;
  }
</style>
