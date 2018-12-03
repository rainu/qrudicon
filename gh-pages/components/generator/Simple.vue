<template>
  <div>
    <form @submit.prevent="onSubmit">
      <div class="form-group">
        <label for="range-slider">Size: {{sizeLabel}}</label>
        <input id="range-slider" type="range" v-model="form.size" class="form-control-range" min="128" max="4096" step="128" />
      </div>

      <div class="form-group">
        <label for="text-field">Text:</label>
        <textarea
            id="text-field"
            v-model="form.text"
            placeholder="Enter something"
            rows="5"
            class="form-control"
            required></textarea>
      </div>

      <div class="form-group">
        <label class="my-1 mr-2" for="qrudicon-type">Type:</label>
        <select class="custom-select my-1 mr-sm-2" id="qrudicon-type" v-model="form.type">
          <option value="image/png">.png</option>
          <option value="image/jpeg">.jpg</option>
        </select>
      </div>

      <div class="form-group container">
        <div class="row">
          <div class="col col-left">
            <button type="reset" class="btn btn-block btn-warning" @click.prevent="onReset">Reset</button>
          </div>
          <div class="col col-right">
            <button type="submit" class="btn btn-block btn-primary">Generate</button>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
  import { minValue, maxValue, minLength, maxLength } from 'vuelidate/lib/validators'

  const defaultValues = {
    size: 1024,
    text: '',
    type: 'image/png'
  }

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
          size: this.value.size || defaultValues.size,
          text: this.value.text || defaultValues.text,
          type: this.value.type || defaultValues.type
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
      },
      onReset() {
        this.form = {
          ...defaultValues
        }
        this.$emit('input', {})
      }
    },
    computed: {
      sizeLabel() {
        return `${this.form.size} x ${this.form.size} px`
      }
    }
  }
</script>

<style scoped>
  .col-left {
    padding-left: 0;
  }
  .col-right {
    padding-right: 0;
  }
</style>