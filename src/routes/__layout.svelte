<script>
	import Nav from '$lib/Nav.svelte'
</script>

<svelte:head>
	<title>Pickup List</title>
</svelte:head>

<Nav />

<div>
	<slot />
</div>

<style>
	div {
		padding: 12px 48px;
	}
</style>
