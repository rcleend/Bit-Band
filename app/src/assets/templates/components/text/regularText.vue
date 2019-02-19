<template>
    <p class="text" v-html="content"></p>
</template>

<script>
    export default {
        props:[
            'content'
        ]
    
    }
</script>

<style lang="scss" scoped>
    /** ANIMATIONS **/
    @include keyframes(blink) {
        0% {
            opacity: 0;
        }
        49% {
            opacity: 0;
        }
        50% {
            opacity: 1;
        }
        100% {
            opacity: 1;
        }
    }

    @include keyframes(loading) {
        0% {
            content: '';
        }
        25% {
            content: '.';
        }
        50% {
            content: '..';
        }
        75% {
            content: '...';
        }
    }

    /** STYLING **/
    .text {
        font-size: 25px;

        @include set-style(tablet){
            font-size: initial;
        }
    }

    .text_alert {
        color: $clr--primary-red;
        font-size: 30px;
        opacity: 0;

        animation: blink 2.5s infinite linear;
        animation-delay: 2s;

        @include set-style(tablet){
            font-size: 25px;
        }
    }

    .text_loading {
        &::after {
            content:'';
            animation: loading 2.5s infinite linear;
        }
        &::before {
            content: '';
            color: black;
            animation: loading 2.5s infinite linear;
        }
    }

    .text_fade-in {
        opacity: 0;
        animation: show 1s linear;
        animation-fill-mode: forwards;
        animation-delay: 1.5s;
    }

</style>


