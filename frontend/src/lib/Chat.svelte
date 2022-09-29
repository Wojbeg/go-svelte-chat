<script>
    import { Message, MessageRequest } from "../utils";
    import { messages, username, userID } from "../store";
    import { beforeUpdate, afterUpdate } from 'svelte';
    import ChatMessage from "./ChatMessage.svelte";

    const serverAddress = import.meta.env.VITE_WEBSOCKET_URL;

    console.log(serverAddress)

    let socket = new WebSocket(serverAddress);
    let currentMessage = "";

    let div;
	let autoscroll;

    beforeUpdate(() => {
		autoscroll = div && (div.offsetHeight + div.scrollTop) > (div.scrollHeight - 20);
	});

    afterUpdate(() => {
		if (autoscroll) div.scrollTo(0, div.scrollHeight);
	});

    socket.onopen = () => {
        console.log("Successfully connected to websocket");
    };

    socket.onmessage = (msg) => {
        console.log("Message received: ", msg.data);

        let data = JSON.parse(msg.data);
        
        if (Array.isArray(data)) {
            console.log("It's an array", data)

            messages.update( array => {
                const resultArray = data.map( obj => new Message(obj.messageId, obj.type, obj.userId, obj.username, obj.created_at, obj.body)) 
                return [...resultArray, ...array];
            }); 

        } else {
            console.log("It's normal data", data)

            var message = new Message(data.messageId, data.type, data.userId, data.username, data.created_at, data.body)
            console.log("message", message)

            //Registered user
            if (message.Type == -1 && message.Body === "You have joined a chat") {
                userID.set(message.UserID)
            }

            messages.update(array => {
                array.push(message)
                return array;
            });
        }

    };

    socket.onclose = (event) => {
        console.log("Websocket closed: ", event);
    };

    socket.onerror = (err) => {
        console.log("Error occured: ", err);
    };

    function sendMessage() {
        if (currentMessage != "") {
            console.log("działa!")
            var toSend = new MessageRequest($username, currentMessage)
            socket.send(JSON.stringify(toSend));
            currentMessage = "";
        }
    }

    function onKeyDown(e) {
		if(e.keyCode == 13) {
            sendMessage()
        }
	}

    function checkToday(date) {
		const today = new Date()
		return date.getDate() == today.getDate() &&
			date.getMonth() == today.getMonth() &&
			date.getFullYear() == today.getFullYear()
	}

</script>

<div class="card">
    <div>

        <div class="inner">
            <div class="card-body">
                <div class="direct-chat-messages" id="direct-chat-messages" bind:this={div}>
                    {#each $messages as message, i (i)}
                        <ChatMessage
                            message={message}
                            sentByMe={message.UserID === $userID}
                            isToday={checkToday(new Date(message.Created_At))}
                        />
                    {/each}
                </div>
            </div>

            <div class="input-group">
                <input
                    type="text"
                    placeholder="Type Message ..."
                    class="message-input"
                    bind:value={currentMessage}
                    on:keydown={onKeyDown}
                />
                <span class="input-group-append">
                    <button type="button" class="btn btn-primary" on:click={sendMessage}>Send</button>
                </span>
            </div>

        </div>
    </div>
</div>


<style>
    .card {
        display: flex;
        flex-direction: column;
        padding: 20px;
        margin: 0 auto;
        width: 70vh;
        height: 70vh;
        background: #1a1a1a;
        box-shadow: 0 8px 64px 0 rgba(22, 36, 233, 0.37);
        backdrop-filter: blur(4px);
        border-radius: 10px;
        border: 1px solid rgba(255, 255, 255, 0.2);
    }

    .inner {
        padding: 10px;
    }

    input {
        font-family: inherit;
        font-size: inherit;
        box-sizing: border-box;
        -webkit-padding: 0.4em 0;
        padding: 0.7em;
        margin: 0 1em 1em 0;
        border: 1px solid rgb(24, 28, 37);
        border-radius: 20px;
        background: rgb(68, 81, 95);
        padding: 10px 20px;
        color: white;
    }

    .inner .card-body {
        overflow-x: hidden;
        padding: 0;
        position: relative;
        margin-bottom: 20px;
    }

    .direct-chat-messages {
        -webkit-transform: translate(0, 0);
        transform: translate(0, 0);
        height: 400px;
        overflow-x: hidden;
        padding: 10px;
        transition: -webkit-transform 0.5s ease-in-out;
        transition: transform 0.5s ease-in-out;
        transition: transform 0.5s ease-in-out,
            -webkit-transform 0.5s ease-in-out;
        
    }

    .direct-chat-messages::-webkit-scrollbar {
        width: 3px;
	    background-color: none;
    }

    .direct-chat-messages::-webkit-scrollbar-thumb {
        background-color: rgb(68, 81, 95);
        border-radius: 15px;
        height: 10px;
    }

</style>
