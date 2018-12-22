
<template>
    <button v-on:click="editAddress(addressId)">Edit</button>
</template>

<script>
export default {
    name:"EditButton",
    props: {
        addressId: String
    },
    methods: {
        editAddress: function (id) {
            //add api call
            return;
        }
    }
}
</script>


<style>

</style>

