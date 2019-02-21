<template>
    <div>
        <div class="header--container">
            <main-title class="main-title_large main-title_fade-in spacing-y--xl"></main-title>
        </div>
        <div class="header--container">
            <sub-title class="spacing-y--sm width--full" content="Make Music Together"></sub-title>
            <p class="text text_alert spacing-y--sm width--full">Please turn up your volume</p>
            <p class="text text_loading text_fade-in spacing-y--sm width--full">Looking for bands to join</p>
        </div>
    </div>
</template>

<script>
    import mainTitle from '../components/text/mainTitle.vue';
    import subTitle from '../components/text/subTitle.vue';

    export default {
        components: {
            mainTitle,
            subTitle
        }
    };
</script>
