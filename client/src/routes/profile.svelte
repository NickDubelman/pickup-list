<script>
  import { goto } from '$app/navigation'
  import { graphqlQuery } from '$lib/graphql'
  import { profile } from '$lib/stores/profile'

  const players = ['Lebron James', 'Anthony Davis', 'Robert Sacre']

  let { realName, nbaName } = $profile

  const canSubmit = realName !== '' && nbaName !== ''

  const setUserMutation = `
    mutation SetUser($input: SetUserInput!){
      setUser(input: $input){
        realName
        nbaName
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
  <label for="real-name">Real name:</label>
  <input
    type="text"
    name="real-name"
    placeholder="First name only is fine if you prefer"
    bind:value={realName}
  />

  <label for="nba-name">NBA name:</label>
  <select name="nba-name" required bind:value={nbaName}>
    <option value="" disabled selected hidden>Select a player</option>
    {#each players as player}
      <option value={player}>{player}</option>
    {/each}
  </select>

  <button type="submit" disabled={!canSubmit}>Done</button>
</form>

<style>
  form {
    width: 420px;
  }

  input[type='text'],
  select {
    width: 100%;
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
