<template>
    <div>
        <va-loading v-if="isLoading" size="lg" color="blue" fixed center></va-loading>
        <h1>Addressbook</h1>
        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aut corporis ipsum minus aspernatur maiores totam.</p>
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
            const path = `${this.ROOT_API}/address/list`
            console.log(path)
            axios.get(path)
                .then((res) => {
                    addresses = res
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

