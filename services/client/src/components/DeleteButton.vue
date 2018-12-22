
<template>
    <button v-on:click="deleteAddress(addressId)">Delete</button>
</template>

<script>
export default {
    name:"DeleteButton",
    props: {
        addressId: String
    },
    methods: {
        deleteAddress: function (id) {
            //add api call
            return;
        }
    }
}
</script>


<style>

</style>

