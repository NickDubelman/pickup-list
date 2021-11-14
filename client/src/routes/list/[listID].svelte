<script context="module">
  import { graphqlQuery } from '$lib/graphql'

  const listQuery = `
    query ListQuery($id: ID!) {
      node(id: $id) {
        ... on List {
          name
          users {
            realName
            nbaPlayer { name }
          }
        }
      }
    }
  `

  export async function load({ page, fetch }) {
    try {
      const data = await graphqlQuery(fetch, {
        query: listQuery,
        variables: { id: page.params.listID }
      })

      return { props: { list: data.node } }
    } catch (error) {
      return { error, status: 500 }
    }
  }
</script>

<script>
  import { page } from '$app/stores'
  import { profile } from '$lib/stores/profile'

  export let list

  $: alreadyJoined = list.users.find(
    user =>
      user.realName === $profile.realName && user.nbaPlayer.name === $profile.nbaName
  )

  const joinListMutation = `
    mutation JoinList($input: JoinListInput!){
      joinList(input: $input){
        users {
          realName
          nbaPlayer { name }
        }
      }
    }
  `

  const unjoinListMutation = `
    mutation UnjoinList($input: JoinListInput!){
      unjoinList(input: $input){
        users {
          realName
          nbaPlayer { name }
        }
      }
    }
  `

  const onJoin = async () => {
    try {
      const { joinList } = await graphqlQuery(fetch, {
        query: joinListMutation,
        variables: { input: { listID: $page.params.listID } }
      })

      list.users = joinList.users
    } catch (e) {
      alert(e)
    }
  }

  const onUnjoin = async () => {
    try {
      const { unjoinList } = await graphqlQuery(fetch, {
        query: unjoinListMutation,
        variables: { input: { listID: $page.params.listID } }
      })

      list.users = unjoinList.users
    } catch (e) {
      alert(e)
    }
  }
</script>

<h1>{list.name}</h1>

{#if $profile && !alreadyJoined}
  <button on:click={onJoin}>Join this list</button>
{:else if $profile && alreadyJoined}
  <button on:click={onUnjoin}>Remove yourself from this list</button>
{:else}
  To join lists, you must first set your <a href="/profile">profile</a>
{/if}

{#if list.users.length > 0}
  <ol>
    {#each list.users as { realName, nbaPlayer }}
      <li>{realName} ({nbaPlayer.name})</li>
    {/each}
  </ol>
{:else}
  <div>No one has joined this list</div>
{/if}

<style>
  h1 {
    margin-bottom: 8px;
  }

  button {
    margin-bottom: 8px;
  }
</style>
