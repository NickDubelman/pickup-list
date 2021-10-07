<script>
  import { goto } from '$app/navigation'

  let password = ''

  function handleSubmit() {
    if (password === 'yekshimesh') {
      goto('/profile')
    } else {
      alert('wrong password')
    }
  }
</script>

<div>
  <h1>🏀 🏀 🏀 Pickup List 🏀 🏀 🏀</h1>
  <p>Add yourself to a list of people who want to play pickup...</p>

  <form on:submit|preventDefault={handleSubmit}>
    <p>Enter the secret password</p>
    <input type="password" bind:value={password} />

    <button type="submit" disabled={password === ''}>Go</button>
  </form>
</div>

<style>
  div {
    border: 1px solid purple;
    padding: 0 24px 24px 24px;
    margin: auto;
    width: 420px;
  }
</style>
