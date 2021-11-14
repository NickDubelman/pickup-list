export async function get({ query }) {
  const accessToken = query.get('accessToken')

  return {
    status: 302, // redirect
    headers: {
      Location: '/',
      'set-cookie': `jwt=${accessToken}; Path=/; HttpOnly`
    }
  }
}
