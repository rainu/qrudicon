<template>
  <div>
    <div class="card" v-if="qrudiconLink" :class="{ 'bg-success': show }">
      <div class="card-body">
        <img :src="qrudiconLink" :alt="qrudiconLink" class="card-img" />
        <button class="btn btn-primary btn-block" @click="copyLink">Copy Link</button>
        <input type="hidden" v-model="qrudiconLink" ref="linkinput" />
      </div>
    </div>
  </div>
</template>

<script>
  import cfg from '../static/config'

  export default {
    name: "qrudicon",
    props: {
      value: {
        required: true
      }
    },

    data(){
      return {
        show : false,
      }
    },

    methods: {
      copyLink(){
        this.$refs.linkinput.setAttribute('type', 'text')
        this.$refs.linkinput.select();
        document.execCommand("copy");
        this.$refs.linkinput.setAttribute('type', 'hidden')

        this.show = true
        setTimeout(() => {
          this.show = false
        }, 1000)
      }
    },

    computed: {
      qrudiconLink() {
        if (this.value && this.value.text) {
          return `${cfg.API_ENDPOINT}?t=${this.value.text}&s=${this.value.size}&e=${this.value.type}`
        }else{
          return null
        }
      }
    }
  }
</script>