<template>
  <div>
    <form @submit.prevent="onSubmit">
      <div class="form-group">
        <label for="range-slider">Size: {{sizeLabel}}</label>
        <input id="range-slider" type="range" v-model="form.size" class="form-control-range" min="128" max="4096" step="128" />
      </div>

      <div class="form-group">
        <label for="nickname">Nickname:</label>
        <div class="input-group">
          <div class="input-group-prepend">
            <span class="input-group-text" ><input type="checkbox" value="form.vcard.NICKNAME" v-model="form.identicon" class="form-check" tabindex="-1" /></span>
          </div>
          <input id="nickname" type="text" class="form-control" v-model="form.vcard.NICKNAME"/>
        </div>
      </div>
      <div class="form-group">
        <label for="firstname">Firstname:</label>
        <div class="input-group">
          <div class="input-group-prepend">
            <span class="input-group-text" ><input type="checkbox" value="form.firstname" v-model="form.identicon" class="form-check" tabindex="-1" /></span>
          </div>
          <input id="firstname" type="text" class="form-control" v-model="form.firstname" @change="changeName"/>
        </div>
      </div>
      <div class="form-group">
        <label for="lastname">Lastname:</label>
        <div class="input-group">
          <div class="input-group-prepend">
            <span class="input-group-text" ><input type="checkbox" value="form.lastname" v-model="form.identicon" class="form-check" tabindex="-1" /></span>
          </div>
          <input id="lastname" type="text" class="form-control" v-model="form.lastname" @change="changeName"/>
        </div>
      </div>
      <div class="form-group">
        <label for="telephone">Telephone:</label>
        <div class="input-group">
          <div class="input-group-prepend">
            <span class="input-group-text" ><input type="checkbox" value="form.vcard.TEL" v-model="form.identicon" class="form-check" tabindex="-1" /></span>
          </div>
          <input id="telephone" type="tel" class="form-control" v-model="form.vcard.TEL"/>
        </div>
      </div>
      <div class="form-group">
        <label for="email">EMail:</label>
        <div class="input-group">
          <div class="input-group-prepend">
            <span class="input-group-text" ><input type="checkbox" value="form.vcard.EMAIL" v-model="form.identicon" class="form-check" tabindex="-1" /></span>
          </div>
          <input id="email" type="email" class="form-control" v-model="form.vcard.EMAIL"/>
        </div>
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
  import cfg from '../../static/config'
  import { minValue, maxValue, minLength, maxLength } from 'vuelidate/lib/validators'

  const defaultValues = {
    size: 1024,
    type: 'image/png',
    link: cfg.API_ENDPOINT.vcard
  };

  const identiconOrder = {
    'form.firstname': 0,
    'form.lastname': 1,
    'form.vcard.NICKNAME': 2,
    'form.vcard.TEL': 3,
    'form.vcard.EMAIL': 4,
  };

  const getValueForKey = (object, key) => {
    if(!object) return null;

    if(!key.includes('.')) {
      return object[key];
    }

    let sKey = key.split('.');
    return getValueForKey(object[sKey[0]], sKey.slice(1).join('.'))
  };

  export default {
    name: "generator-vcard",
    props: {
      value: {
        required: true
      }
    },

    data() {
      return {
        form: {
          size: this.value.size || defaultValues.size,
          type: this.value.type || defaultValues.type,
          vcard : {},
          identicon: []
        },
      }
    },

    validations: {
      form: {
        size: {
          minValue: minValue(128),
          maxValue: maxValue(4096),
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
          this.form.link = this.generateLink()
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
      },
      changeName(){
        this.form.vcard.FN = `${this.form.firstname} ${this.form.lastname}`
      },
      generateLink() {
        let url = `${cfg.API_ENDPOINT.vcard}?s=${this.form.size}&e=${this.form.type}&t=${this.identiconContent}`

        for(let key in this.form.vcard) {
          url += `&VC_${key}=${encodeURIComponent(this.form.vcard[key])}`
        }

        return url
      },
    },
    computed: {
      sizeLabel() {
        return `${this.form.size} x ${this.form.size} px`
      },
      identiconContent() {
        let sortedKeys = this.form.identicon.sort((a, b) => {
          if(identiconOrder[a] > identiconOrder[b]) return 1;
          else if(identiconOrder[a] < identiconOrder[b]) return -1;
          else return 0
        })

        let content = []
        for(let key of sortedKeys) {
          content.push(getValueForKey(this, key))
        }

        console.log(content)
        return content.join(' ')
      }
    },
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