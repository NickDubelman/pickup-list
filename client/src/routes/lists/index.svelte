<script context="module">
  export async function load({ fetch }) {
    const listsQuery = `{
      lists {
        id
        name
        users {
          id
          realName
          nbaName
        }
      }
    }`

    try {
      const { lists } = await graphqlQuery(fetch, { query: listsQuery })
      return { props: { listData: lists } }
    } catch (error) {
      return { error, status: 500 }
    }
  }
</script>

<script>
  import { lists } from '$lib/stores/lists'

  import AddList from '$lib/AddList.svelte'
  import { graphqlQuery } from '$lib/graphql'

  export let listData
  lists.set(listData)
</script>

<h1>Lists for this week</h1>
<h3>Monday October 25 ➜ Sunday October 31</h3>

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
