<script context="module">
  import { graphqlQuery } from '$lib/graphql'

  export async function load({ fetch }) {
    const playersQuery = `{ nbaPlayers { name } }`

    try {
      const { nbaPlayers } = await graphqlQuery(fetch, { query: playersQuery })
      return { props: { nbaPlayers } }
    } catch (error) {
      return { error, status: 500 }
    }
  }
</script>

<script>
  import { goto } from '$app/navigation'
  import { profile } from '$lib/stores/profile'

  export let nbaPlayers

  let { realName, nbaName } = $profile

  $: canSubmit = realName !== ''

  const setUserMutation = `
    mutation SetUser($input: SetUserInput!){
      setUser(input: $input){
        realName
        nbaPlayer { name }
      }
    }
  `

  async function handleSubmit() {
    // Request to set user
    try {
      await graphqlQuery(fetch, {
        query: setUserMutation,
        variables: { input: { realName, nbaName } }
      })
      profile.set({ ...$profile, realName, nbaName })
      alert('Updated profile!')
      goto('/lists')
    } catch (e) {
      alert(e)
    }
  }
</script>

<h1>Create your profile</h1>

<form on:submit|preventDefault={handleSubmit}>
  <div>
    <label for="real-name">Real name:</label>
    <input
      type="text"
      name="real-name"
      placeholder="First name only is fine if you prefer"
      bind:value={realName}
    />
  </div>

  <div>
    <label for="nba-name">NBA name:</label>
    <select name="nba-name" bind:value={nbaName}>
      <option value="" disabled selected hidden>Select a player</option>
      {#each nbaPlayers as player}
        <option value={player.name}>{player.name}</option>
      {/each}
    </select>

    {#if nbaName && nbaName !== ''}
      <button title="Unset" on:click|preventDefault={() => (nbaName = '')}>X</button>
    {/if}
  </div>

  <button type="submit" disabled={!canSubmit}>Done</button>
</form>

<style>
  form {
    width: 420px;
  }

  input[type='text'],
  select {
    width: 90%;
    padding: 6px 10px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-sizing: border-box;
  }

  option {
    color: 'red';
  }

  button[type='submit'] {
    width: 100%;
    background-color: #4caf50;
    color: white;
    padding: 14px 20px;
    margin: 8px 0;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  button[type='submit']:disabled {
    background-color: #cccccc;
  }
</style>
