<template>
    <div>
        <main>
            <title-screen></title-screen>
        </main>
        <footer>
            <p class="text">&#169; 2019 <a href='https://github.com/rcleend'>Rcleend</a> & <a href='#'>Jelmuis</a>. All Rights Reserved.</p>
        </footer>
    </div>
</template>

<script>
    import titleScreen from './sections/titleScreen.vue';

    export default {
        components: {
            titleScreen,
        }
    };
</script>

