<template>
    <h2 class="sub-title sub-title_large sub-title_fade-in">
        <span class="sub-title_span-before">></span>
        {{content}}
        <span class="sub-title_span-after"><</span>
    </h2>    
</template>

<script>
    export default {
        props:[
            'content'
        ] 
    }
</script>