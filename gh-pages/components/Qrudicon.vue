<template>
  <div>
    <div class="card" v-if="value" :class="{ 'bg-success': show }">
      <div class="card-body">
        <img :src="value" :alt="value" class="card-img" />
        <button class="btn btn-primary btn-block" @click="copyLink">Copy Link</button>
        <input type="hidden" v-model="value" ref="linkinput" />
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "qrudicon",
    props: {
      value: {
        required: true
      },
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
  }
</script>