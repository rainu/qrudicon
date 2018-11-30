<template>
  <div>
    <b-form @submit.prevent="onSubmit" novalidate>
      <b-form-group label="Size:">
        <b-input-group :prepend="sizeLabel">
          <b-form-input type="range" v-model="form.size" min="128" max="4096"></b-form-input>
        </b-input-group>
      </b-form-group>

      <b-form-group label="Text:">
        <b-form-textarea v-model="form.text"
                         placeholder="Enter something"
                         :rows="5"
                         :max-rows="6">
        </b-form-textarea>
      </b-form-group>

      <b-form-group label="Type:">
        <b-form-radio-group v-model="form.type">
          <b-form-radio value="image/png">.png</b-form-radio>
          <b-form-radio value="image/jpeg">.jpg</b-form-radio>
        </b-form-radio-group>
      </b-form-group>

      <b-button type="submit" variant="primary">Generate</b-button>
    </b-form>
  </div>
</template>

<script>
  import { minValue, maxValue, minLength, maxLength } from 'vuelidate/lib/validators'

  export default {
    name: "generator-simple",
    props: {
      value: {
        required: true
      }
    },

    data() {
      return {
        form: {
          size: this.value.size || 1024,
          text: this.value.text || '',
          type: this.value.type || 'image/png'
        },
      }
    },

    validations: {
      form: {
        size: {
          minValue: minValue(128),
          maxValue: maxValue(4096),
        },
        text: {
          minLength: minLength(1),
          maxLength: maxLength(1536)
        },
        type: {

        }
      }
    },

    methods: {
      onSubmit() {
        this.$v.$touch();

        //emit only valid data
        if(!this.$v.form.$error) {
          this.$emit('input', {
            ...this.form
          })
        }
      }
    },
    computed: {
      sizeLabel() {
        return `${this.form.size} x ${this.form.size} px`
      }
    }
  }
</script>
