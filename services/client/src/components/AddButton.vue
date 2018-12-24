
<template>
    <div>
        <va-button size="lg" type="primary" @click="showAddModal">Add New Address</va-button>
        <va-modal width="600px" title="New Address" ref="addModal" :backdrop-clickable="false">
            <div slot="body">
                <va-form ref="form" type="inline" @submit.prevent="close">
                    <va-input v-model="address.name" placeholder="Name" name="address-name" id="address-name" />
                    <va-input v-model="address.phone_number" placeholder="Phone" name="address-phone" id="address-phone" />
                    <va-button type="primary" @click.native="submit">Submit</va-button>
                </va-form>
            </div>
            <div slot="footer" />
        </va-modal>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name:"AddButton",
    methods: {
        showAddModal () {
            this.$refs.addModal.open()
        },
        submit() {
            //send to backend
            console.log("ROOT_API", process.env)
            axios.post(process.env.BASE_URL_API+"/address/create", this.address, {
                    headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                    }
                })
                .then((response) => {
                    this.$emit('update')
                    this.address.name = ''
                    this.address.phone_number = ''
                    this.$refs.addModal.close()
                })
                .catch((error) => {
                    console.log(error)
                })
        }
    },
    data: function () {
        return {
            address : {
                name: '',
                phone_number: ''
            }
        }
    }
}
</script>


<style>

</style>

