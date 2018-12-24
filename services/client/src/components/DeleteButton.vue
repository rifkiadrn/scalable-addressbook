
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
            let data = {
                id: parseInt(id)
            }
            console.log("data", data)
            console.log("ROOT_API", process.env)
            axios.post(process.env.BASE_URL_API+"/address/delete", data,
                {
                    _method: 'delete',
                    headers: {
                    'Content-Type': 'application/x-x-www-form-urlencoded'
                    }
                })
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

