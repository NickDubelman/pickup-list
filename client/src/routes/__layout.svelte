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
          user: user && {
            realName: user.realName,
            nbaName: user.nbaPlayer?.name || ''
          }
        }
      }
    } catch (error) {
      return { error, status: 500 }
    }
  }
</script>

<script lang="ts">
  import Nav from '$lib/Nav.svelte'
  import { profile } from '$lib/stores/profile'
  import Landing from './_landing.svelte'

  export let user = null
  profile.set(user)
</script>

<svelte:head>
  <title>Pickup List</title>
</svelte:head>

{#if user}
  <Nav />

  <div>
    <slot />
  </div>
{:else}
  <Landing />
{/if}

<style>
  div {
    padding: 12px 48px;
  }
</style>
