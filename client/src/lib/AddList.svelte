<script>
  import { lists } from '$lib/stores/lists'
  import { graphqlQuery } from '$lib/graphql'

  let adding = false
  let listName = ''

  const createListMutation = `
    mutation CreateList($input: CreateListInput!){
      createList(input: $input){
        id
        name
      }
    }
  `

  async function handleSubmit() {
    try {
      const { createList } = await graphqlQuery(fetch, {
        query: createListMutation,
        variables: { input: { name: listName } }
      })
      const { id, name } = createList
      lists.addList(id, name)
      onCancelAdd()
    } catch (e) {
      alert(e)
    }
  }

  function onCancelAdd() {
    adding = false
    listName = ''
  }
</script>

<div>
  <button on:click={() => (adding = true)}>Add list</button>

  {#if adding}
    <form on:submit|preventDefault={handleSubmit}>
      <input
        type="text"
        placeholder="name of list (ex: Friday Lunch)"
        bind:value={listName}
      />

      <button type="submit" disabled={listName.trim() === ''}>Add</button>
      <button on:click={onCancelAdd}>Cancel</button>
    </form>
  {/if}
</div>

<style>
  div {
    margin-bottom: 16px;
  }

  input {
    width: 224px;
  }

  form {
    display: inline;
  }
</style>
