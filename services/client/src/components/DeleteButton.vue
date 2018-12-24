
<template>
    <va-button type="danger" size="md" v-on:click="deleteAddress(addressId)">Delete</va-button>
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
            this.$emit('update')
            return;
        }
    }
}
</script>


<style>

</style>

