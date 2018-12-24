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
            addresses: [{id: '1', name:'a', phone_number:'0342'},
            {id: '2', name:'aha', phone_number:'1324'},
            {id: '3', name:'ahaha', phone_number:'22432'},
            {id: '4', name:'ahahaha', phone_number:'034'},
            {id: '5', name:'ahahahaha', phone_number:'0243'}],
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

