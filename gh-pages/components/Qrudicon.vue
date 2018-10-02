<template>
  <div>
    <b-card border-variant="primary"
            :img-src="qrudiconLink"
            img-alt="qrudicon"
            style="max-width: 20rem;"
            class="mb-2"
            v-if="qrudiconLink">

      <b-button id="qrudicon-copy-link" variant="primary" @click="copyLink">Copy Link</b-button>
      <input type="hidden" v-model="qrudiconLink" ref="linkinput" />

      <b-popover :show.sync="show" target="qrudicon-copy-link" title="Success" placement="top">
        Link copied to clipboard!
      </b-popover>
    </b-card>
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