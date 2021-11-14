<script context="module">
  import { graphqlQuery } from '$lib/graphql'

  export async function load({ fetch }) {
    const userQuery = `{
      user {
        realName
        nbaPlayer { name }
      }
    }`

    try {
      const { user } = await graphqlQuery(fetch, { query: userQuery })
      return {
        props: {
          user: {
            realName: user.realName,
            nbaName: user.nbaPlayer.name
          }
        }
      }
    } catch (error) {
      return { error, status: 500 }
    }
  }
</script>

<script>
  import Nav from '$lib/Nav.svelte'
  import { profile } from '$lib/stores/profile'

  export let user = { realName: '', nbaName: '' }
  profile.set(user)
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
