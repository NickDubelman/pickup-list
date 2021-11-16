export async function get({ query }) {
  const accessToken = query.get('accessToken')
  const refreshToken = query.get('refreshToken')

  return {
    status: 302, // redirect
    headers: {
      Location: '/',
      'set-cookie': [
        `jwt=${accessToken}; Path=/; HttpOnly`,
        `refresh=${refreshToken}; Path=/; HttpOnly`
      ]
    }
  }
}
