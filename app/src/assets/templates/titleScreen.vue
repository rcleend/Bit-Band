<template>
    <div class="header--container">
        <div class="header--container">
            <main-title class="main-title_large main-title_fade-in spacing-y--xl"></main-title>
        </div>
        <sub-title class="spacing-y--sm width--full" content="Make Music Together"></sub-title>
        <regular-text class="text_alert spacing-y--sm width--full" content="Please turn up your volume"></regular-text>
        <regular-text class="text_loading text_fade-in spacing-y--sm width--full" content="Looking for bands to join"></regular-text>
    </div>
</template>

<script>
    import mainTitle from './components/text/mainTitle.vue';
    import subTitle from './components/text/subTitle.vue';
    import regularText from './components/text/regularText.vue';

    export default {
        components: {
            mainTitle,
            subTitle,
            regularText
        }
    };
</script>

<style lang="scss" scoped>
    .header--container {
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        justify-content: center;
        text-align: center;
    }
</style>

