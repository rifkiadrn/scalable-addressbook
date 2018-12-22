<template>
    <form @submit.prevent="close">
        <input type="text" name="address-name" id="address-name">
        <input type="text" name="phonenumber" id="phonenumber">
        <input type="submit" value="submit">
    </form>
</template>

<script>
export default {
    name: "FormModal",
    props: {
        endpoint: String
    },
    methods: {
        close: function () {
            return;
        }
    }
}
</script>

<style>

</style>
