export async function get({ query }) {
  const accessToken = query.get('accessToken')
  const refreshToken = query.get('refreshToken')
  const urlState = query.get('state') || '/'

  return {
    status: 302, // redirect
    headers: {
      Location: urlState,
      'set-cookie': [
        `jwt=${accessToken}; Path=/; HttpOnly`,
        `refresh=${refreshToken}; Path=/; HttpOnly`
      ]
    }
  }
}
