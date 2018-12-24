<template>
    <div>
        <va-loading v-if="isLoading" size="lg" color="blue" fixed center></va-loading>
        <h1>Addressbook</h1>
        <add-button @update="update"/>
        <the-table :addresses="addresses" @update="update"/>
    </div>
</template>

<script>
import AddButton from "./AddButton"
import TheTable from "./TheTable"
import axios from 'axios'

export default {
    components : {
        AddButton,
        TheTable
    },
    data () {
        return {
            addresses: [],
            isLoading: false,
            ROOT_API: process.env.ROOT_API
        }
    },
    methods: {
        update () {
            console.log("test here")
            console.log(this.isLoading)
            this.isLoading = true
            //get from backend
            console.log("ROOT_API", process.env)
            axios.get(process.env.BASE_URL_API+"/address/list")
                .then((res) => {
                    console.log("res",res.data)
                    this.addresses = res.data.result.data
                })
                .catch(error => {
                    console.log(error)
                    alert(error)
                })
                this.isLoading = false
        }
    },
    mounted: function() {
        this.update()
    }
}
</script>

