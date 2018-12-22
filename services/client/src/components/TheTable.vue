<template>
    <table>
        <tr v-for="address in addresses" :key="address.id">
            <td>{{address.name}}</td>
            <td>{{address.phoneNumber}}
            <td><edit-button address-id="address.id"/><delete-button address-id="address.id"/></td>
        </tr>
    </table>
</template>

<script>
import EditButton from "EditButton";
import DeleteButton from "DeleteButton";

export default {
    name: "TheTable",
    components: {
        EditButton,
        DeleteButton
    },
    props: {
        addresses: Array
    }
}
</script>

<style>

</style>
