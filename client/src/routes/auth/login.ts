export function get({ query }) {
  const next = query.get('next')

  return {
    status: 302, // redirect
    headers: {
      Location: `http://localhost:8080/auth/login?next=${next}`
    }
  }
}
