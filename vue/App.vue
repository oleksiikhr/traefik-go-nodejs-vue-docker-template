<template>
  <button @click="onClick">Send</button>
  <p>Loading: {{ state.loading ? 'Yes' : 'No' }}</p>
  <p>Result: {{ state.result }}</p>
  <code>$ docker-compose logs -f node</code>
</template>

<script>
import { reactive } from 'vue'
import axios from 'axios'

export default {
  name: 'App',
  data() {
    return {
      result: null
    }
  },
  setup() {
    const state = reactive({ loading: false, result: null })

    const onClick = () => {
      state.loading = true

      axios.post('/api/')
        .then(({ data }) => state.result = data)
        .catch(({ response: { data } }) => state.result = data)
        .finally(() => state.loading = false)
    }

    return { onClick, state }
  }
}
</script>
