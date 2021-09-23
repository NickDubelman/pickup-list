<script>
	import Nav from '$lib/Nav.svelte'
</script>

<Nav />

<div>
	<slot />
</div>

<style>
	div {
		padding: 12px 48px;
	}
</style>
