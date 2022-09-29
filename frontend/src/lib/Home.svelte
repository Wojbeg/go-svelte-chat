<script>
  import { started } from "../store";

  const handleStart = () => {
    started.set(true)
  }
</script>

  <div>
    <img src="/ic_go.svg" class="logo go" alt="Svelte Logo" />

    <img src="/svelte.svg" class="logo svelte" alt="Vite Logo" />
  </div>

  <h1>Go Svelte Chat</h1>

  <div class="card">
    <button on:click={handleStart}>
        Start
    </button>
  </div>

<style>

  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
  }

  .logo.go:hover {
    filter: drop-shadow(0 0 2em #00acd7);
  }

  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }

  h1 {
    color: rgba(212, 26, 26, 0.87);
  }
</style>