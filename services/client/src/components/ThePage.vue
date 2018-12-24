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

export default {
    components : {
        AddButton,
        TheTable
    },
    data: function () {
        return {
            addresses: 
            [{id: '1', name:'a', phone:'0342'},
            {id: '2', name:'aha', phone:'1324'},
            {id: '3', name:'ahaha', phone:'22432'},
            {id: '4', name:'ahahaha', phone:'034'},
            {id: '5', name:'ahahahaha', phone:'0243'}],
            isLoading: true
        }
    },
    methods: {
        update () {
            //get from backend
        }
    },
    mounted () {
        this.update()
    }
}
</script>

