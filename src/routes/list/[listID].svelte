<script>
	import { page } from '$app/stores'
	import { lists } from '$lib/stores/lists'

	$: listID = +$page.params.listID
	$: list = $lists.find(l => l.id === listID)

	const onJoinList = () => {
		lists.joinList(listID, 'Nick Dubelman', 'Robert Sacre')
	}
</script>

<h1>Specific List!</h1>
<h2>12:00PM</h2>

<button on:click={onJoinList}>Join this run</button>

<ol>
	{#each list.people as { realName, nbaName }}
		<li>{realName} ({nbaName})</li>
	{/each}
</ol>

<style>
	h1 {
		margin-bottom: 8px;
	}

	h2 {
		margin-top: 8px;
		margin-bottom: 8px;
	}
</style>
