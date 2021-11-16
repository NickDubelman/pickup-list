export function get() {
  return {
    status: 302, // redirect
    headers: {
      Location: '/',
      'set-cookie': [
        'jwt=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT',
        'refresh=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT'
      ]
    }
  }
}
