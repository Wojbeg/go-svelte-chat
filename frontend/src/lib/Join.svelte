<script>
    import { canConnect, username } from "../store";
    let hasError = false;

    function handleInput(e) {
        username.set(e.target.value.trim());
        hasError = false;
    }

    const handleJoin = () => {
        if ($username === "") {
            hasError = true;
        } else {
            canConnect.set(true)
        }
    }

</script>

<h1>Go Svelte Chat</h1>

{#if hasError == true}
    <p class="error-alert">Username can not be empty</p>
{/if}

<form id="joinForm">
    <input type="text" placeholder="Username" on:input={handleInput} required>

    <button on:click={() => {handleJoin()}}>
        Go!
    </button>
    
</form>


<style>
    input {
        font-family: inherit;
        font-size: inherit;
        box-sizing: border-box;
        -webkit-padding: 0.4em 0;
        padding: 0.6em;
        margin: 0 1em 1em 0;
        border: 1px solid rgb(24, 28, 37);
        border-radius: 20px;
        background:rgb(68, 81, 95);
        padding: 10px 20px;
    }

    input {
        color: white;
    }

    input::placeholder{
        color:rgb(149, 168, 189)
    }

    input:disabled {
        color: #ccc;
    }

    .error-alert {
		padding: 6px;
		text-align: center;
		color: #c00;
		border-radius: 3px;
	}
</style>