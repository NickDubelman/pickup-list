<script>
  import AddList from '$lib/AddList.svelte'
  import { lists } from '$lib/stores/lists'
</script>

<h1>Lists for this week</h1>
<h3>Monday September 20 ➜ Sunday September 26</h3>

<AddList />
{#if $lists.length > 0}
  {#each $lists as { id, name }}
    <a href={`/list/${id}`}>{name}</a>
  {/each}
{:else}
  <div>No lists have been added yet</div>
{/if}

<style>
  h1 {
    margin-bottom: 4px;
  }
  h3 {
    margin-top: 4px;
  }

  a {
    display: block;
    font-size: 1.24em;
    margin-top: 8px;
  }
</style>
