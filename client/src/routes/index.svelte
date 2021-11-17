<script context="module">
  export async function load({ fetch }) {
    const listsQuery = `
      query ListsQuery($from: Time!){
        lists(from: $from){
          id
          name
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
        variables: { from: getStartOfPrevSunday().toISOString() }
      })
      return { props: { listData: lists } }
    } catch (error) {
      return { error, status: 500 }
    }
  }
</script>

<script lang="ts">
  import { lists } from '$lib/stores/lists'

  import AddList from '$lib/AddList.svelte'
  import { graphqlQuery } from '$lib/graphql'
  import { getStartOfPrevSunday, monthNames } from '$lib/dateutils'

  export let listData
  lists.set(listData)

  const prevSunday = getStartOfPrevSunday()

  const formatDate = (date: Date) => {
    return `${monthNames[date.getMonth()]} ${date.getDate()}`
  }
</script>

<h1>Lists for week of {formatDate(prevSunday)}</h1>

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

  a {
    display: block;
    font-size: 1.24em;
    margin-top: 8px;
  }
</style>
