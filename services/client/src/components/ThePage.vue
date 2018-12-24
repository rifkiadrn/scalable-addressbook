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
    data: function () {
        return {
            addresses: [],
            isLoading: true
        }
    },
    methods: {
        update () {
            isLoading = false
            //get from backend
            axios.get(process.env.BASE_API_URL + '/api/address/list')
                .then((res) => {
                    addresses = res
                    isLoading = false
                })
                .catch(error => {
                    console.log(error)
                })
        }
    },
    mounted () {
        this.update()
    }
}
</script>

