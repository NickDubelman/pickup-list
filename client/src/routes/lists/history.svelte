<script context="module">
  import { getStartOfPrevSunday } from '$lib/dateutils'
  import { graphqlQuery } from '$lib/graphql'

  export async function load({ fetch }) {
    const listsQuery = `
        query ListsQuery($to: Time!){
          lists(to: $to){
            id
            name
            createdAt
            users {
              id
              realName
              nbaPlayer { name }
            }
          }
        }
      `

    try {
      const { lists } = await graphqlQuery(fetch, {
        query: listsQuery,
        variables: { to: getStartOfPrevSunday().toISOString() }
      })
      return { props: { listData: lists } }
    } catch (error) {
      return { error, status: 500 }
    }
  }
</script>

<script lang="ts">
  export let listData

  const getLists = (weeksBack: number) => {
    return listData.filter(list => {
      const end = getStartOfPrevSunday()
      end.setDate(end.getDate() - 7 * (weeksBack - 1))

      const start = new Date(end.getTime())
      start.setDate(start.getDate() - 7)
      return start < new Date(list.createdAt) && new Date(list.createdAt) < end
    })
  }
</script>

<h1>History</h1>

<div>
  <h2>Last week</h2>
  {#if getLists(1).length > 0}
    {#each getLists(1) as { id, name }}
      <a href={`/list/${id}`}>{name}</a>
    {/each}
  {:else}
    <div>No lists</div>
  {/if}
</div>

<div>
  <h2>Two weeks ago</h2>
  {#if getLists(2).length > 0}
    {#each getLists(2) as { id, name }}
      <a href={`/list/${id}`}>{name}</a>
    {/each}
  {:else}
    <div>No lists</div>
  {/if}
</div>

<div>
  <h2>Three weeks ago</h2>
  {#if getLists(3).length > 0}
    {#each getLists(3) as { id, name }}
      <a href={`/list/${id}`}>{name}</a>
    {/each}
  {:else}
    <div>No lists</div>
  {/if}
</div>

<!-- <button>Show older</button> -->
<style>
  a {
    display: block;
    font-size: 1.24em;
    margin-top: 8px;
  }

  div {
    padding-bottom: 16px;
  }
</style>
