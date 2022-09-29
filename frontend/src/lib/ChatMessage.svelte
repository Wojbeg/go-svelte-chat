<script>

    export let sentByMe;
	export var message;		
	export let isToday;

</script>


<div class="direct-chat-msg" class:right="{sentByMe}" class:left="{!sentByMe}">
    
    <div class="direct-chat-text">
		{message.Body}
	</div>
	
	<div class="direct-chat-infos">

        <span class="direct-chat-timestamp" class:float-left="{sentByMe}" class:float-right="{!sentByMe}">
			{message.Username}
			{#if isToday}
				{new Date(message.Created_At).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit', second:'2-digit'})}
			{:else}
				{new Date(message.Created_At).toLocaleTimeString([], {year: 'numeric', month: '2-digit', day: '2-digit', hour: '2-digit', minute:'2-digit', second:'2-digit', hour12: false})}
			{/if}
		</span>
    </div>
</div>


<style>

.direct-chat-msg,
.direct-chat-text {
	display: block;
	white-space: normal;
	word-break: break-word;
}

.direct-chat-msg {
	margin-bottom: 10px;
}

.direct-chat-msg:before,
.direct-chat-msg:after {
	content: "";
	display: table;
}

.direct-chat-msg:after {
	clear: both;
}

.direct-chat-text {
	border-radius: 5px;
	position: relative;
	padding: 1px 6px 1px 6px;
	background: #4278e3;
	border: 1px solid #d2d6de;
	margin: 0;
	color: rgb(53, 42, 42);
	margin-right: 50px;
	margin-left: 50px;
	text-align: left;
}

.direct-chat-text:after,
.direct-chat-text:before {
	position: absolute;
	right: 100%;
	top: 2px;
	border: solid transparent;
	border-right-color: #4278e3;
	content: '';
	height: 0;
	width: 0;
	pointer-events: none;
}

.direct-chat-text:after {
	border-width: 6px;
	margin-top: 0px;
}

.direct-chat-text:before {
	border-width: 6px;
	margin-top: 0px;
}

.right .direct-chat-text {
	margin-right: 0px;
	margin-left: 50px;
}

.left .direct-chat-text {
	margin-right: 50px;
	margin-left: 0px;
}

.right .direct-chat-text:after,
.right .direct-chat-text:before {
	right: auto;
	left: 100%;
	bottom: 100%;
	border-right-color: transparent;
	border-left-color: #4278e3;
}

.direct-chat-infos {
	font-size: .5rem;
	text-align: end;
}

.left .direct-chat-infos {
	font-size: .5rem;
	text-align: start;
}

.direct-chat-timestamp {
	color: #999;
	margin: 0;
	padding: 0;
}

</style>
