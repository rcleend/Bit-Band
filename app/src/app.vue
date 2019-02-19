<template>
    <div>
        <main>
            <title-screen></title-screen>
        </main>
        <footer>
            <regular-text content="&#169; 2019 <a href='https://github.com/rcleend'>Rcleend</a> & <a href='#'>Jelmuis</a>. All Rights Reserved."></regular-text>
        </footer>
    </div>
</template>

<script>
    import titleScreen from './assets/templates/titleScreen.vue';
    import regularText from'./assets/templates/components/text/regularText.vue';

    export default {
        components: {
            titleScreen,
            regularText
        }
    };
</script>

<style lang="scss">
    * {
        font-family: 'PressStart2P', cursive;
        color: white;

        margin: 0;
        padding: 0;
    }

    body {
        background-color: #121212;
        background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='79' height='79' viewBox='0 0 100 100'%3E%3Crect x='0' y='0' width='95' height='95' fill='%23000000'/%3E%3C/svg%3E");
        background-position: center center;

        height: 100%;
    }

    main {
        min-height: 100vh;
    }

    footer {
        height: 50px;
        margin-top: -50px;

        text-align: center;
    }
</style>

