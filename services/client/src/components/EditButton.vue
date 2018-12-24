
<template>
    <span>
        <va-button size="md" type="primary" @click="showEditModal" class="edit-button">Edit</va-button>
        <va-modal width="600px" title="Edit Address" ref="editModal" :backdrop-clickable="false">
            <div slot="body">
                <va-form ref="form" type="inline" @submit.prevent="close">
                    <va-input v-model="address.name" placeholder="Address" name="address-name" id="address-name" />
                    <va-input v-model="address.phone" placeholder="Phone" name="address-phone" id="address-phone" />
                    <va-button type="primary" @click.native="submit">Save</va-button>
                </va-form>
            </div>
            <div slot="footer" />
        </va-modal>
    </span>
</template>

<script>
export default {
    name:"EditButton",
    props: {
        address: Object
    },
    methods: {
        showEditModal () {
            this.$refs.editModal.open()
        },
        submit() {
            //send to backend

            this.$refs.editModal.close()
            this.$emit('update')
        }
    }
}
</script>


<style>
.va-form-inline {
    display: flex;
    justify-content: space-between;
}
</style>