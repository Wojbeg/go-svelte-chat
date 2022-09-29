<script>
    import Chat from './lib/Chat.svelte';
    import Home from './lib/Home.svelte';
    import Join from './lib/Join.svelte';
    import { started, canConnect } from './store';
</script>

<main>
  {#if $started && $canConnect}
    <Chat />
  {:else if $started}
    <Join />
  {:else}
    <Home />
  {/if}
</main>

<style>
  
</style>
