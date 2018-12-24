<template>
    <va-table :hover="hover" size="sm">
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Phone Number</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="address in addresses" :key="address.id">
                    <td>{{address.name}}</td>
                    <td>{{address.phone_number}}</td>
                    <td><edit-button @update="emitUpdate" :address="address"/></td>
                    <td><delete-button :address-id="address.id"/></td>
                </tr>
            </tbody>
        </table>
    </va-table>
</template>

<script>
import EditButton from "./EditButton";
import DeleteButton from "./DeleteButton";

export default {
    name: "TheTable",
    components: {
        EditButton,
        DeleteButton
    },
    props: {
        addresses: Array
    },
    method: {
        emitUpdate() {
            this.$emit('update')
        }
    }
}
</script>

<style>

</style>
