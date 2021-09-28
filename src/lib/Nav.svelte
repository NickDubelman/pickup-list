<script>
  import { page } from '$app/stores'

  const links = [
    { href: '/lists', label: 'Home' },
    { href: '/profile', label: 'Profile' },
    { href: '/lists/history', label: 'History' }
  ]
</script>

{#if $page.path !== '/'}
  <nav>
    {#each links as { href, label }}
      <a {href} class:active={$page.path === href}>{label}</a>
    {/each}

    <a class="logout" href="/">Logout</a>
  </nav>
{/if}

<style>
  nav {
    list-style-type: none;
    margin: 0;
    padding: 0;
    overflow: hidden;
    background: #333;
  }

  a {
    display: inline-block;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
  }

  a:hover:not(.active, .logout) {
    background: #111;
  }

  .active {
    background: #04aa6d;
  }

  .logout {
    float: right;
    background: #df5b1b;
  }
</style>
