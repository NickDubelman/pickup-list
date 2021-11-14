import * as cookie from 'cookie'

export async function handle({ request, resolve }) {
  const cookies = cookie.parse(request.headers.cookie || '')
  request.locals.jwt = cookies.jwt
  return await resolve(request)
}

export function getSession({ locals }) {
  return {
    user: locals.user && {
      email: locals.user.email
    }
  }
}
