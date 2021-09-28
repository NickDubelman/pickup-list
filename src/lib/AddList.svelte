<script>
	import { lists } from '$lib/stores/lists'

	let adding = false
	let listName = ''

	function handleSubmit() {
		const id = Math.max(...$lists.map((list) => list.id)) + 1
		lists.update((prev) => [...prev, { id, name: listName.trim() }])
		onCancelAdd()
	}

	function onCancelAdd() {
		adding = false
		listName = ''
	}
</script>

<div>
	<button on:click={() => (adding = true)}>Add list</button>

	{#if adding}
		<form on:submit|preventDefault={handleSubmit}>
			<input
				type="text"
				placeholder="name of list (ex: Friday Lunch)"
				bind:value={listName}
			/>

			<button type="submit" disabled={listName.trim() === ''}>Add</button>
			<button on:click={onCancelAdd}>Cancel</button>
		</form>
	{/if}
</div>

<style>
	div {
		margin-bottom: 16px;
	}

	input {
		width: 224px;
	}

	form {
		display: inline;
	}
</style>
