<template>
<div>
    <main>
        <component v-bind:is="currentView"></component>
    </main>
    <footer>
        <p class="text">&#169; 2019 <a href='https://github.com/rcleend'>Rcleend</a> & <button @click="setCurrentView('sequencer')">Jelmuis</button>. All Rights Reserved.</p>
    </footer>
</div>
</template>

<script>
import titleScreen from './sections/titleScreen.vue';
import sequencer from './sections/sequencer.vue';

export default {
    data() {
        return {
            currentView: "titleScreen"
        }
    },
    components: {
        titleScreen,
        sequencer
    },
    methods: {
        setCurrentView: function(component) {
            this.$data.currentView = component; 
        }
    }
};
</script>

