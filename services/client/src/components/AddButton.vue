
<template>
    <div>
        <va-button size="lg" type="primary" @click="showAddModal">Add New Address</va-button>
        <va-modal width="600px" title="New Address" ref="addModal" :backdrop-clickable="false">
            <div slot="body">
                <va-form ref="form" type="inline" @submit.prevent="close">
                    <va-input v-model="address.name" placeholder="Address" name="address-name" id="address-name" />
                    <va-input v-model="address.phone" placeholder="Phone" name="address-phone" id="address-phone" />
                    <va-button type="primary" @click.native="submit">Submit</va-button>
                </va-form>
            </div>
            <div slot="footer" />
        </va-modal>
    </div>
</template>

<script>
export default {
    name:"AddButton",
    methods: {
        showAddModal () {
            this.$refs.addModal.open()
        },
        submit() {
            //send to backend

            this.$emit('update')
            this.address.name = ''
            this.address.phone = ''
            this.$refs.addModal.close()
        }
    },
    data: function () {
        return {
            address : {
                name: '',
                phone: ''
            }
        }
    }
}
</script>


<style>

</style>

