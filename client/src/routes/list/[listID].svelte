<script>
  import { page } from '$app/stores'

  import { lists } from '$lib/stores/lists'
  import { profile } from '$lib/stores/profile'

  $: listID = +$page.params.listID
  $: list = $lists.find(l => l.id === listID)

  const onJoinList = () => {
    lists.joinList(listID, $profile)
  }

  $: alreadyJoined = list.people.find(
    person =>
      person.realName === $profile.realName && person.nbaName === $profile.nbaName
  )
</script>

<h1>Specific List!</h1>
<h2>12:00PM</h2>

{#if $profile && !alreadyJoined}
  <button on:click={onJoinList}>Join this run</button>
{/if}

{#if list.people.length > 0}
  <ol>
    {#each list.people as { realName, nbaName }}
      <li>{realName} ({nbaName})</li>
    {/each}
  </ol>
{:else}
  <div>No one has joined this list</div>
{/if}

{#if !$profile}
  To join the list, you must set your <a href="/profile">profile</a>
{/if}

<style>
  h1 {
    margin-bottom: 8px;
  }

  h2 {
    margin-top: 8px;
    margin-bottom: 8px;
  }
</style>
