
<template>
    <va-button type="danger" size="md" v-on:click="deleteAddress(addressId)">Delete</va-button>
</template>

<script>
import axios from 'axios'

export default {
    name:"DeleteButton",
    props: {
        addressId: String
    },
    methods: {
        deleteAddress: function (id) {
            //add api call
            axios.post(process.env.ROOT_API + '/address/delete', id)
                .then((res) => {
                    this.$emit('update')
                })
                .catch(error => {
                    console.log(error)
                })
        }
    }
}
</script>


<style>

</style>

